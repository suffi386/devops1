package oidc

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/caos/oidc/pkg/op"
	"github.com/rakyll/statik/fs"
	"golang.org/x/text/language"

	"github.com/caos/zitadel/internal/api/assets"
	http_utils "github.com/caos/zitadel/internal/api/http"
	"github.com/caos/zitadel/internal/api/http/middleware"
	"github.com/caos/zitadel/internal/api/ui/login"
	"github.com/caos/zitadel/internal/auth/repository"
	"github.com/caos/zitadel/internal/command"
	"github.com/caos/zitadel/internal/config/systemdefaults"
	"github.com/caos/zitadel/internal/crypto"
	"github.com/caos/zitadel/internal/eventstore"
	"github.com/caos/zitadel/internal/eventstore/handler/crdb"
	"github.com/caos/zitadel/internal/i18n"
	"github.com/caos/zitadel/internal/query"
	"github.com/caos/zitadel/internal/telemetry/metrics"
)

const (
	HandlerPrefix = "/oauth/v2"
	AuthCallback  = HandlerPrefix + "/authorize/callback?id="
)

type Config struct {
	CodeMethodS256                    bool
	AuthMethodPost                    bool
	AuthMethodPrivateKeyJWT           bool
	GrantTypeRefreshToken             bool
	RequestObjectSupported            bool
	SigningKeyAlgorithm               string
	DefaultAccessTokenLifetime        time.Duration
	DefaultIdTokenLifetime            time.Duration
	DefaultRefreshTokenIdleExpiration time.Duration
	DefaultRefreshTokenExpiration     time.Duration
	UserAgentCookieConfig             *middleware.UserAgentCookieConfig
	Cache                             *middleware.CacheConfig
	CustomEndpoints                   *EndpointConfig
}

type EndpointConfig struct {
	Auth          *Endpoint
	Token         *Endpoint
	Introspection *Endpoint
	Userinfo      *Endpoint
	Revocation    *Endpoint
	EndSession    *Endpoint
	Keys          *Endpoint
}

type Endpoint struct {
	Path string
	URL  string
}

type OPStorage struct {
	repo                              repository.Repository
	command                           *command.Commands
	query                             *query.Queries
	eventstore                        *eventstore.Eventstore
	defaultLoginURL                   string
	defaultAccessTokenLifetime        time.Duration
	defaultIdTokenLifetime            time.Duration
	signingKeyAlgorithm               string
	defaultRefreshTokenIdleExpiration time.Duration
	defaultRefreshTokenExpiration     time.Duration
	encAlg                            crypto.EncryptionAlgorithm
	keyChan                           <-chan interface{}
	currentKey                        query.PrivateKey
	signingKeyRotationCheck           time.Duration
	signingKeyGracefulPeriod          time.Duration
	locker                            crdb.Locker
	assetAPIPrefix                    string
}

func NewProvider(ctx context.Context, config Config, issuer, defaultLogoutRedirectURI string, command *command.Commands, query *query.Queries, repo repository.Repository, keyConfig systemdefaults.KeyConfig, es *eventstore.Eventstore, projections *sql.DB, keyChan <-chan interface{}, keyStorage crypto.KeyStorage, oidcKeyConfig *crypto.KeyConfig, userAgentCookie func(http.Handler) http.Handler) (op.OpenIDProvider, error) {
	opConfig, err := createOPConfig(config, issuer, defaultLogoutRedirectURI, keyStorage, oidcKeyConfig)
	if err != nil {
		return nil, fmt.Errorf("cannot create op config: %w", err)
	}
	storage, err := newStorage(config, command, query, repo, keyConfig, oidcKeyConfig, es, projections, keyStorage, keyChan)
	if err != nil {
		return nil, fmt.Errorf("cannot create storage: %w", err)
	}
	options, err := createOptions(config, userAgentCookie)
	if err != nil {
		return nil, fmt.Errorf("cannot create options: %w", err)
	}
	provider, err := op.NewOpenIDProvider(
		ctx,
		opConfig,
		storage,
		options...,
	)
	if err != nil {
		return nil, fmt.Errorf("cannot create provider: %w", err)
	}
	return provider, nil
}

func Issuer(domain string, port uint16, externalSecure bool) string {
	return http_utils.BuildHTTP(domain, port, externalSecure) + HandlerPrefix
}

func createOPConfig(config Config, issuer, defaultLogoutRedirectURI string, keyStorage crypto.KeyStorage, oidcKeyConfig *crypto.KeyConfig) (*op.Config, error) {
	supportedLanguages, err := getSupportedLanguages()
	if err != nil {
		return nil, err
	}
	opConfig := &op.Config{
		Issuer:                   issuer,
		DefaultLogoutRedirectURI: defaultLogoutRedirectURI,
		CodeMethodS256:           config.CodeMethodS256,
		AuthMethodPost:           config.AuthMethodPost,
		AuthMethodPrivateKeyJWT:  config.AuthMethodPrivateKeyJWT,
		GrantTypeRefreshToken:    config.GrantTypeRefreshToken,
		RequestObjectSupported:   config.RequestObjectSupported,
		SupportedUILocales:       supportedLanguages,
	}
	if err := cryptoKey(opConfig, oidcKeyConfig, keyStorage); err != nil {
		return nil, err
	}
	return opConfig, nil
}

func cryptoKey(config *op.Config, keyConfig *crypto.KeyConfig, keyStorage crypto.KeyStorage) error {
	tokenKey, err := crypto.LoadKey(keyStorage, keyConfig.EncryptionKeyID)
	if err != nil {
		return fmt.Errorf("cannot load OP crypto key: %w", err)
	}
	cryptoKey := []byte(tokenKey)
	if len(cryptoKey) != 32 {
		return fmt.Errorf("OP crypto key must be exactly 32 bytes")
	}
	copy(config.CryptoKey[:], cryptoKey)
	return nil
}

func createOptions(config Config, userAgentCookie func(http.Handler) http.Handler) ([]op.Option, error) {
	metricTypes := []metrics.MetricType{metrics.MetricTypeRequestCount, metrics.MetricTypeStatusCode, metrics.MetricTypeTotalCount}
	interceptor := op.WithHttpInterceptors(
		middleware.MetricsHandler(metricTypes),
		middleware.TelemetryHandler(),
		middleware.NoCacheInterceptor,
		userAgentCookie,
		http_utils.CopyHeadersToContext,
	)
	endpoints := customEndpoints(config.CustomEndpoints)
	if len(endpoints) == 0 {
		return []op.Option{interceptor}, nil
	}
	return append(endpoints, interceptor), nil
}

func customEndpoints(endpointConfig *EndpointConfig) []op.Option {
	if endpointConfig == nil {
		return nil
	}
	options := []op.Option{}
	if endpointConfig.Auth != nil {
		options = append(options, op.WithCustomAuthEndpoint(op.NewEndpointWithURL(endpointConfig.Auth.Path, endpointConfig.Auth.URL)))
	}
	if endpointConfig.Token != nil {
		options = append(options, op.WithCustomTokenEndpoint(op.NewEndpointWithURL(endpointConfig.Token.Path, endpointConfig.Token.URL)))
	}
	if endpointConfig.Introspection != nil {
		options = append(options, op.WithCustomIntrospectionEndpoint(op.NewEndpointWithURL(endpointConfig.Introspection.Path, endpointConfig.Introspection.URL)))
	}
	if endpointConfig.Userinfo != nil {
		options = append(options, op.WithCustomUserinfoEndpoint(op.NewEndpointWithURL(endpointConfig.Userinfo.Path, endpointConfig.Userinfo.URL)))
	}
	if endpointConfig.Revocation != nil {
		options = append(options, op.WithCustomRevocationEndpoint(op.NewEndpointWithURL(endpointConfig.Revocation.Path, endpointConfig.Revocation.URL)))
	}
	if endpointConfig.EndSession != nil {
		options = append(options, op.WithCustomEndSessionEndpoint(op.NewEndpointWithURL(endpointConfig.EndSession.Path, endpointConfig.EndSession.URL)))
	}
	if endpointConfig.Keys != nil {
		options = append(options, op.WithCustomKeysEndpoint(op.NewEndpointWithURL(endpointConfig.Keys.Path, endpointConfig.Keys.URL)))
	}
	return options
}

func newStorage(config Config, command *command.Commands, query *query.Queries, repo repository.Repository, keyConfig systemdefaults.KeyConfig, c *crypto.KeyConfig, es *eventstore.Eventstore, projections *sql.DB, keyStorage crypto.KeyStorage, keyChan <-chan interface{}) (*OPStorage, error) {
	encAlg, err := crypto.NewAESCrypto(c, keyStorage)
	if err != nil {
		return nil, err
	}
	return &OPStorage{
		repo:                              repo,
		command:                           command,
		query:                             query,
		eventstore:                        es,
		defaultLoginURL:                   fmt.Sprintf("%s%s?%s=", login.HandlerPrefix, login.EndpointLogin, login.QueryAuthRequestID),
		signingKeyAlgorithm:               config.SigningKeyAlgorithm,
		defaultAccessTokenLifetime:        config.DefaultAccessTokenLifetime,
		defaultIdTokenLifetime:            config.DefaultIdTokenLifetime,
		defaultRefreshTokenIdleExpiration: config.DefaultRefreshTokenIdleExpiration,
		defaultRefreshTokenExpiration:     config.DefaultRefreshTokenExpiration,
		encAlg:                            encAlg,
		signingKeyGracefulPeriod:          keyConfig.SigningKeyGracefulPeriod,
		signingKeyRotationCheck:           keyConfig.SigningKeyRotationCheck,
		locker:                            crdb.NewLocker(projections, locksTable, signingKey),
		keyChan:                           keyChan,
		assetAPIPrefix:                    assets.HandlerPrefix,
	}, nil
}

func (o *OPStorage) Health(ctx context.Context) error {
	return o.repo.Health(ctx)
}

func getSupportedLanguages() ([]language.Tag, error) {
	statikLoginFS, err := fs.NewWithNamespace("login")
	if err != nil {
		return nil, err
	}
	return i18n.SupportedLanguages(statikLoginFS)
}
