package command

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/zitadel/logging"

	"github.com/zitadel/zitadel/internal/api/authz"
	"github.com/zitadel/zitadel/internal/crypto"
	"github.com/zitadel/zitadel/internal/domain"
	"github.com/zitadel/zitadel/internal/eventstore"
	"github.com/zitadel/zitadel/internal/id"
	"github.com/zitadel/zitadel/internal/repository/authrequest"
	"github.com/zitadel/zitadel/internal/repository/oidcsession"
	"github.com/zitadel/zitadel/internal/telemetry/tracing"
	"github.com/zitadel/zitadel/internal/zerrors"
)

const (
	TokenDelimiter            = "-"
	AccessTokenPrefix         = "at_"
	RefreshTokenPrefix        = "rt_"
	oidcTokenSubjectDelimiter = ":"
	oidcTokenFormat           = "%s" + oidcTokenSubjectDelimiter + "%s"
)

type OIDCSession struct {
	TokenID           string
	ClientID          string
	UserID            string
	Audience          []string
	Expiration        time.Time
	Scope             []string
	AuthMethods       []domain.UserAuthMethodType
	AuthTime          time.Time
	PreferredLanguage string
	UserAgent         *domain.UserAgent
	Reason            domain.TokenReason
	Actor             *domain.TokenActor
	RefreshToken      string
}

type AuthRequestComplianceChecker func(context.Context, *AuthRequestWriteModel) error

// CreateOIDCSessionFromAuthRequest creates a new OIDC Session, creates an access token and refresh token.
// It returns the access token id, expiration and the refresh token.
// If the underlying [AuthRequest] is a OIDC Auth Code Flow, it will set the code as exchanged.
func (c *Commands) CreateOIDCSessionFromAuthRequest(ctx context.Context, authReqId string, complianceCheck AuthRequestComplianceChecker) (session *OIDCSession, state string, err error) {
	ctx, span := tracing.NewSpan(ctx)
	defer func() { span.EndWithError(err) }()

	if authReqId == "" {
		return nil, "", zerrors.ThrowPreconditionFailed(nil, "COMMAND-Sf3g2", "Errors.AuthRequest.InvalidCode")
	}

	authReqModel, err := c.getAuthRequestWriteModel(ctx, authReqId)
	if err != nil {
		return nil, "", err
	}

	if authReqModel.ResponseType == domain.OIDCResponseTypeCode && authReqModel.AuthRequestState != domain.AuthRequestStateCodeAdded {
		return nil, "", zerrors.ThrowPreconditionFailed(nil, "COMMAND-Iung5", "Errors.AuthRequest.NoCode")
	}

	sessionModel := NewSessionWriteModel(authReqModel.SessionID, authz.GetInstance(ctx).InstanceID())
	err = c.eventstore.FilterToQueryReducer(ctx, sessionModel)
	if err != nil {
		return nil, "", err
	}
	if err = sessionModel.CheckIsActive(); err != nil {
		return nil, "", err
	}

	cmd, err := c.newOIDCSessionAddEvents(ctx, sessionModel.UserResourceOwner)
	if err != nil {
		return nil, "", err
	}
	if authReqModel.ResponseType == domain.OIDCResponseTypeCode {
		if err = cmd.SetAuthRequestCodeExchanged(ctx, authReqModel); err != nil {
			return nil, "", err
		}
	}
	if err = complianceCheck(ctx, authReqModel); err != nil {
		return nil, "", err
	}

	cmd.AddSession(ctx, sessionModel.UserID, sessionModel.AggregateID, authReqModel.ClientID, authReqModel.Audience, authReqModel.Scope, authReqModel.AuthMethods, authReqModel.AuthTime, sessionModel.UserAgent)

	if authReqModel.ResponseType != domain.OIDCResponseTypeIDToken {
		if err = cmd.AddAccessToken(ctx, authReqModel.Scope, domain.TokenReasonAuthRequest, nil); err != nil {
			return nil, "", err
		}
	}
	if authReqModel.NeedRefreshToken {
		if err = cmd.AddRefreshToken(ctx, sessionModel.UserID); err != nil {
			return nil, "", err
		}
	}
	cmd.SetAuthRequestSuccessful(ctx, authReqModel.aggregate)
	session, err = cmd.PushEvents(ctx)
	return session, authReqModel.State, err
}

func (c *Commands) CreateOIDCSession(ctx context.Context, userID, resourceOwner, clientID string, scope, audience []string, authMethods []domain.UserAuthMethodType, authTime time.Time, userAgent *domain.UserAgent, reason domain.TokenReason, actor *domain.TokenActor, needRefreshToken bool) (session *OIDCSession, err error) {
	ctx, span := tracing.NewSpan(ctx)
	defer func() { span.EndWithError(err) }()

	cmd, err := c.newOIDCSessionAddEvents(ctx, resourceOwner)
	if err != nil {
		return nil, err
	}
	cmd.AddSession(ctx, userID, "", clientID, audience, scope, authMethods, authTime, userAgent)
	if err = cmd.AddAccessToken(ctx, scope, reason, actor); err != nil {
		return nil, err
	}
	if needRefreshToken {
		if err = cmd.AddRefreshToken(ctx, userID); err != nil {
			return nil, err
		}
	}
	return cmd.PushEvents(ctx)
}

type RefreshTokenComplianceChecker func(context.Context, *OIDCSessionWriteModel) error

// ExchangeOIDCSessionRefreshAndAccessToken updates an existing OIDC Session, creates a new access and refresh token.
// It returns the access token id and expiration and the new refresh token.
func (c *Commands) ExchangeOIDCSessionRefreshAndAccessToken(ctx context.Context, refreshToken string, scope []string, complianceCheck RefreshTokenComplianceChecker) (_ *OIDCSession, err error) {
	ctx, span := tracing.NewSpan(ctx)
	defer func() { span.EndWithError(err) }()

	cmd, err := c.newOIDCSessionUpdateEvents(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	if err = complianceCheck(ctx, cmd.oidcSessionWriteModel); err != nil {
		return nil, err
	}
	if err = cmd.AddAccessToken(ctx, scope, domain.TokenReasonRefresh, nil); err != nil {
		return nil, err
	}
	if err = cmd.RenewRefreshToken(ctx); err != nil {
		return nil, err
	}
	return cmd.PushEvents(ctx)
}

// OIDCSessionByRefreshToken computes the current state of an existing OIDCSession by a refresh_token (to start a Refresh Token Grant).
// If either the session is not active, the token is invalid or expired (incl. idle expiration) an invalid refresh token error will be returned.
func (c *Commands) OIDCSessionByRefreshToken(ctx context.Context, refreshToken string) (_ *OIDCSessionWriteModel, err error) {
	ctx, span := tracing.NewSpan(ctx)
	defer func() { span.EndWithError(err) }()

	oidcSessionID, refreshTokenID, err := parseRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}
	writeModel := NewOIDCSessionWriteModel(oidcSessionID, "")
	err = c.eventstore.FilterToQueryReducer(ctx, writeModel)
	if err != nil {
		return nil, zerrors.ThrowPreconditionFailed(err, "OIDCS-SAF31", "Errors.OIDCSession.RefreshTokenInvalid")
	}
	if err = writeModel.CheckRefreshToken(refreshTokenID); err != nil {
		return nil, err
	}
	return writeModel, nil
}

func oidcSessionTokenIDsFromToken(token string) (oidcSessionID, refreshTokenID, accessTokenID string, err error) {
	split := strings.Split(token, TokenDelimiter)
	if len(split) != 2 {
		return "", "", "", zerrors.ThrowPreconditionFailed(nil, "OIDCS-S87kl", "Errors.OIDCSession.Token.Invalid")
	}
	if strings.HasPrefix(split[1], RefreshTokenPrefix) {
		return split[0], split[1], "", nil
	}
	if strings.HasPrefix(split[1], AccessTokenPrefix) {
		return split[0], "", split[1], nil
	}
	return "", "", "", zerrors.ThrowPreconditionFailed(nil, "OIDCS-S87kl", "Errors.OIDCSession.Token.Invalid")
}

// RevokeOIDCSessionToken revokes an access_token or refresh_token
// if the OIDCSession cannot be retrieved by the provided token, is not active or if the token is already revoked,
// then no error will be returned.
// The only possible error (except db connection or other internal errors) occurs if a client tries to revoke a token,
// which was not part of the audience.
func (c *Commands) RevokeOIDCSessionToken(ctx context.Context, token, clientID string) (err error) {
	oidcSessionID, refreshTokenID, accessTokenID, err := oidcSessionTokenIDsFromToken(token)
	if err != nil {
		logging.WithError(err).Info("token revocation with invalid token format")
		return nil
	}
	writeModel := NewOIDCSessionWriteModel(oidcSessionID, "")
	err = c.eventstore.FilterToQueryReducer(ctx, writeModel)
	if err != nil {
		return zerrors.ThrowInternal(err, "OIDCS-NB3t2", "Errors.Internal")
	}
	if err = writeModel.CheckClient(clientID); err != nil {
		return err
	}
	if refreshTokenID != "" {
		if err = writeModel.CheckRefreshToken(refreshTokenID); err != nil {
			logging.WithFields("oidcSessionID", oidcSessionID, "refreshTokenID", refreshTokenID).WithError(err).
				Info("refresh token revocation with invalid token")
			return nil
		}
		return c.pushAppendAndReduce(ctx, writeModel, oidcsession.NewRefreshTokenRevokedEvent(ctx, writeModel.aggregate))
	}
	if err = writeModel.CheckAccessToken(accessTokenID); err != nil {
		logging.WithFields("oidcSessionID", oidcSessionID, "accessTokenID", accessTokenID).WithError(err).
			Info("access token revocation with invalid token")
		return nil
	}
	return c.pushAppendAndReduce(ctx, writeModel, oidcsession.NewAccessTokenRevokedEvent(ctx, writeModel.aggregate))
}

func (c *Commands) newOIDCSessionAddEvents(ctx context.Context, resourceOwner string, pending ...eventstore.Command) (*OIDCSessionEvents, error) {
	accessTokenLifetime, refreshTokenLifeTime, refreshTokenIdleLifetime, err := c.tokenTokenLifetimes(ctx)
	if err != nil {
		return nil, err
	}
	sessionID, err := c.idGenerator.Next()
	if err != nil {
		return nil, err
	}
	sessionID = IDPrefixV2 + sessionID
	return &OIDCSessionEvents{
		eventstore:               c.eventstore,
		idGenerator:              c.idGenerator,
		encryptionAlg:            c.keyAlgorithm,
		events:                   pending,
		oidcSessionWriteModel:    NewOIDCSessionWriteModel(sessionID, resourceOwner),
		accessTokenLifetime:      accessTokenLifetime,
		refreshTokenLifeTime:     refreshTokenLifeTime,
		refreshTokenIdleLifetime: refreshTokenIdleLifetime,
	}, nil
}

func (c *Commands) decryptRefreshToken(refreshToken string) (sessionID, refreshTokenID string, err error) {
	decoded, err := base64.RawURLEncoding.DecodeString(refreshToken)
	if err != nil {
		return "", "", zerrors.ThrowInvalidArgument(err, "OIDCS-Cux9a", "Errors.User.RefreshToken.Invalid")
	}
	decrypted, err := c.keyAlgorithm.DecryptString(decoded, c.keyAlgorithm.EncryptionKeyID())
	if err != nil {
		return "", "", err
	}
	return parseRefreshToken(decrypted)
}

func parseRefreshToken(refreshToken string) (oidcSessionID, refreshTokenID string, err error) {
	split := strings.Split(refreshToken, TokenDelimiter)
	if len(split) < 2 || !strings.HasPrefix(split[1], RefreshTokenPrefix) {
		return "", "", zerrors.ThrowPreconditionFailed(nil, "OIDCS-JOI23", "Errors.OIDCSession.RefreshTokenInvalid")
	}
	// the oidc library requires that every token has the format of <tokenID>:<userID>
	// the V2 tokens don't use the userID anymore, so let's just remove it
	return split[0], strings.Split(split[1], oidcTokenSubjectDelimiter)[0], nil
}

func (c *Commands) newOIDCSessionUpdateEvents(ctx context.Context, refreshToken string) (*OIDCSessionEvents, error) {
	oidcSessionID, refreshTokenID, err := c.decryptRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}
	sessionWriteModel := NewOIDCSessionWriteModel(oidcSessionID, "")
	if err = c.eventstore.FilterToQueryReducer(ctx, sessionWriteModel); err != nil {
		return nil, err
	}
	if err = sessionWriteModel.CheckRefreshToken(refreshTokenID); err != nil {
		return nil, err
	}
	accessTokenLifetime, refreshTokenLifeTime, refreshTokenIdleLifetime, err := c.tokenTokenLifetimes(ctx)
	if err != nil {
		return nil, err
	}
	return &OIDCSessionEvents{
		eventstore:               c.eventstore,
		idGenerator:              c.idGenerator,
		encryptionAlg:            c.keyAlgorithm,
		oidcSessionWriteModel:    sessionWriteModel,
		accessTokenLifetime:      accessTokenLifetime,
		refreshTokenLifeTime:     refreshTokenLifeTime,
		refreshTokenIdleLifetime: refreshTokenIdleLifetime,
	}, nil
}

type OIDCSessionEvents struct {
	eventstore            *eventstore.Eventstore
	idGenerator           id.Generator
	encryptionAlg         crypto.EncryptionAlgorithm
	events                []eventstore.Command
	oidcSessionWriteModel *OIDCSessionWriteModel

	accessTokenLifetime      time.Duration
	refreshTokenLifeTime     time.Duration
	refreshTokenIdleLifetime time.Duration

	// accessTokenID is set by the command
	accessTokenID string

	// refreshToken is set by the command
	refreshTokenID string
	refreshToken   string
}

func (c *OIDCSessionEvents) AddSession(
	ctx context.Context,
	userID,
	sessionID,
	clientID string,
	audience,
	scope []string,
	authMethods []domain.UserAuthMethodType,
	authTime time.Time,
	userAgent *domain.UserAgent,
) {
	c.events = append(c.events, oidcsession.NewAddedEvent(
		ctx,
		c.oidcSessionWriteModel.aggregate,
		userID,
		sessionID,
		clientID,
		audience,
		scope,
		authMethods,
		authTime,
		userAgent,
	))
}

func (c *OIDCSessionEvents) SetAuthRequestCodeExchanged(ctx context.Context, model *AuthRequestWriteModel) error {
	event := authrequest.NewCodeExchangedEvent(ctx, model.aggregate)
	model.AppendEvents(event)
	c.events = append(c.events, event)
	return model.Reduce()
}

func (c *OIDCSessionEvents) SetAuthRequestSuccessful(ctx context.Context, authRequestAggregate *eventstore.Aggregate) {
	c.events = append(c.events, authrequest.NewSucceededEvent(ctx, authRequestAggregate))
}

func (c *OIDCSessionEvents) SetAuthRequestFailed(ctx context.Context, authRequestAggregate *eventstore.Aggregate, err error) {
	c.events = append(c.events, authrequest.NewFailedEvent(ctx, authRequestAggregate, domain.OIDCErrorReasonFromError(err)))
}

func (c *OIDCSessionEvents) AddAccessToken(ctx context.Context, scope []string, reason domain.TokenReason, actor *domain.TokenActor) error {
	accessTokenID, err := c.idGenerator.Next()
	if err != nil {
		return err
	}
	c.accessTokenID = AccessTokenPrefix + accessTokenID
	c.events = append(c.events, oidcsession.NewAccessTokenAddedEvent(ctx, c.oidcSessionWriteModel.aggregate, c.accessTokenID, scope, c.accessTokenLifetime, reason, actor))
	return nil
}

func (c *OIDCSessionEvents) AddRefreshToken(ctx context.Context, userID string) (err error) {
	c.refreshTokenID, c.refreshToken, err = c.generateRefreshToken(userID)
	if err != nil {
		return err
	}
	c.events = append(c.events, oidcsession.NewRefreshTokenAddedEvent(ctx, c.oidcSessionWriteModel.aggregate, c.refreshTokenID, c.refreshTokenLifeTime, c.refreshTokenIdleLifetime))
	return nil
}

func (c *OIDCSessionEvents) RenewRefreshToken(ctx context.Context) (err error) {
	var refreshTokenID string
	refreshTokenID, c.refreshToken, err = c.generateRefreshToken(c.oidcSessionWriteModel.UserID)
	if err != nil {
		return err
	}
	c.events = append(c.events, oidcsession.NewRefreshTokenRenewedEvent(ctx, c.oidcSessionWriteModel.aggregate, refreshTokenID, c.refreshTokenIdleLifetime))
	return nil
}

func (c *OIDCSessionEvents) generateRefreshToken(userID string) (refreshTokenID, refreshToken string, err error) {
	refreshTokenID, err = c.idGenerator.Next()
	if err != nil {
		return "", "", err
	}
	refreshTokenID = RefreshTokenPrefix + refreshTokenID
	token, err := c.encryptionAlg.Encrypt([]byte(fmt.Sprintf(oidcTokenFormat, c.oidcSessionWriteModel.OIDCRefreshTokenID(refreshTokenID), userID)))
	if err != nil {
		return "", "", err
	}
	return refreshTokenID, base64.RawURLEncoding.EncodeToString(token), nil
}

func (c *OIDCSessionEvents) PushEvents(ctx context.Context) (*OIDCSession, error) {
	pushedEvents, err := c.eventstore.Push(ctx, c.events...)
	if err != nil {
		return nil, err
	}
	err = AppendAndReduce(c.oidcSessionWriteModel, pushedEvents...)
	if err != nil {
		return nil, err
	}
	session := &OIDCSession{
		ClientID:          c.oidcSessionWriteModel.ClientID,
		UserID:            c.oidcSessionWriteModel.UserID,
		Audience:          c.oidcSessionWriteModel.Audience,
		Expiration:        c.oidcSessionWriteModel.AccessTokenExpiration,
		Scope:             c.oidcSessionWriteModel.Scope,
		AuthMethods:       c.oidcSessionWriteModel.AuthMethods,
		AuthTime:          c.oidcSessionWriteModel.AuthTime,
		PreferredLanguage: "", // ??
		UserAgent:         c.oidcSessionWriteModel.UserAgent,
		Reason:            c.oidcSessionWriteModel.AccessTokenReason,
		Actor:             c.oidcSessionWriteModel.AccessTokenActor,
		RefreshToken:      c.refreshToken,
	}
	if c.accessTokenID != "" {
		// prefix the returned id with the oidcSessionID so that we can retrieve it later on
		// we need to use `-` as a delimiter because the OIDC library uses `:` and will check for a length of 2 parts
		session.TokenID = c.oidcSessionWriteModel.AggregateID + TokenDelimiter + c.accessTokenID
	}

	return session, nil
}

func (c *Commands) tokenTokenLifetimes(ctx context.Context) (accessTokenLifetime time.Duration, refreshTokenLifetime time.Duration, refreshTokenIdleLifetime time.Duration, err error) {
	oidcSettings := NewInstanceOIDCSettingsWriteModel(ctx)
	err = c.eventstore.FilterToQueryReducer(ctx, oidcSettings)
	if err != nil {
		return 0, 0, 0, err
	}
	accessTokenLifetime = c.defaultAccessTokenLifetime
	refreshTokenLifetime = c.defaultRefreshTokenLifetime
	refreshTokenIdleLifetime = c.defaultRefreshTokenIdleLifetime
	if oidcSettings.AccessTokenLifetime > 0 {
		accessTokenLifetime = oidcSettings.AccessTokenLifetime
	}
	if oidcSettings.RefreshTokenExpiration > 0 {
		refreshTokenLifetime = oidcSettings.RefreshTokenExpiration
	}
	if oidcSettings.RefreshTokenIdleExpiration > 0 {
		refreshTokenIdleLifetime = oidcSettings.RefreshTokenIdleExpiration
	}
	return accessTokenLifetime, refreshTokenLifetime, refreshTokenIdleLifetime, nil
}
