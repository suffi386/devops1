package idp

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zitadel/logging"

	"github.com/zitadel/zitadel/internal/api/authz"
	http_utils "github.com/zitadel/zitadel/internal/api/http"
	"github.com/zitadel/zitadel/internal/command"
	"github.com/zitadel/zitadel/internal/crypto"
	z_errs "github.com/zitadel/zitadel/internal/errors"
	"github.com/zitadel/zitadel/internal/form"
	"github.com/zitadel/zitadel/internal/idp"
	"github.com/zitadel/zitadel/internal/idp/providers/azuread"
	"github.com/zitadel/zitadel/internal/idp/providers/github"
	"github.com/zitadel/zitadel/internal/idp/providers/gitlab"
	"github.com/zitadel/zitadel/internal/idp/providers/google"
	"github.com/zitadel/zitadel/internal/idp/providers/jwt"
	"github.com/zitadel/zitadel/internal/idp/providers/ldap"
	"github.com/zitadel/zitadel/internal/idp/providers/oauth"
	openid "github.com/zitadel/zitadel/internal/idp/providers/oidc"
	saml2 "github.com/zitadel/zitadel/internal/idp/providers/saml"
	"github.com/zitadel/zitadel/internal/query"
)

const (
	HandlerPrefix       = "/idps"
	callbackPath        = "/callback"
	defaultMetadataPath = "/saml/metadata"
	defaultAcsPath      = "/saml/acs"
	metadataPath        = "/{" + varIDPID + ":[0-9]+}" + defaultMetadataPath
	acsPath             = "/{" + varIDPID + ":[0-9]+}" + defaultAcsPath

	paramIntentID         = "id"
	paramToken            = "token"
	paramUserID           = "user"
	paramError            = "error"
	paramErrorDescription = "error_description"
	varIDPID              = "idpid"
)

type Handler struct {
	commands            *command.Commands
	queries             *query.Queries
	parser              *form.Parser
	encryptionAlgorithm crypto.EncryptionAlgorithm
	callbackURL         func(ctx context.Context) string
	samlRootURL         func(ctx context.Context, idpID string) string
}

type externalIDPCallbackData struct {
	State            string `schema:"state"`
	Code             string `schema:"code"`
	Error            string `schema:"error"`
	ErrorDescription string `schema:"error_description"`
}

type externalSAMLIDPCallbackData struct {
	IDPID      string
	Response   string
	RelayState string
}

// CallbackURL generates the instance specific URL to the IDP callback handler
func CallbackURL(externalSecure bool) func(ctx context.Context) string {
	return func(ctx context.Context) string {
		return http_utils.BuildOrigin(authz.GetInstance(ctx).RequestedHost(), externalSecure) + HandlerPrefix + callbackPath
	}
}

func SAMLRootURL(externalSecure bool) func(ctx context.Context, idpID string) string {
	return func(ctx context.Context, idpID string) string {
		return http_utils.BuildOrigin(authz.GetInstance(ctx).RequestedHost(), externalSecure) + HandlerPrefix + "/" + idpID + "/"
	}
}

func NewHandler(
	commands *command.Commands,
	queries *query.Queries,
	encryptionAlgorithm crypto.EncryptionAlgorithm,
	externalSecure bool,
	instanceInterceptor func(next http.Handler) http.Handler,
) http.Handler {
	h := &Handler{
		commands:            commands,
		queries:             queries,
		parser:              form.NewParser(),
		encryptionAlgorithm: encryptionAlgorithm,
		callbackURL:         CallbackURL(externalSecure),
		samlRootURL:         SAMLRootURL(externalSecure),
	}

	router := mux.NewRouter()
	router.Use(instanceInterceptor)
	router.HandleFunc(callbackPath, h.handleCallback)
	router.HandleFunc(metadataPath, h.handleMetadata)
	router.HandleFunc(acsPath, h.handleACS)
	return router
}

func parseSAMLRequest(r *http.Request) *externalSAMLIDPCallbackData {
	vars := mux.Vars(r)
	return &externalSAMLIDPCallbackData{
		IDPID:      vars[varIDPID],
		Response:   r.FormValue("SAMLResponse"),
		RelayState: r.FormValue("RelayState"),
	}
}

func (h *Handler) getProvider(ctx context.Context, idpID string) (idp.Provider, error) {
	return h.commands.GetProvider(ctx, idpID, h.callbackURL(ctx), h.samlRootURL(ctx, idpID))
}

func (h *Handler) handleMetadata(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data := parseSAMLRequest(r)

	provider, err := h.getProvider(ctx, data.IDPID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	samlProvider, ok := provider.(*saml2.Provider)
	if !ok {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sp, err := samlProvider.GetSP()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sp.ServeMetadata(w, r)
}

func (h *Handler) handleACS(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data := parseSAMLRequest(r)

	provider, err := h.getProvider(ctx, data.IDPID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	samlProvider, ok := provider.(*saml2.Provider)
	if !ok {
		err := z_errs.ThrowInvalidArgument(nil, "SAML-ui9wyux0hp", "Errors.Intent.IDPInvalid")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	intent, err := h.commands.GetActiveIntent(ctx, data.RelayState)
	if err != nil {
		if z_errs.IsNotFound(err) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		redirectToFailureURLErr(w, r, intent, err)
		return
	}

	sp, err := samlProvider.GetSP()
	if err != nil {
		cmdErr := h.commands.FailIDPIntent(ctx, intent, err.Error())
		logging.WithFields("intent", intent.AggregateID).OnError(cmdErr).Error("failed to push failed event on idp intent")
		redirectToFailureURLErr(w, r, intent, err)
		return
	}

	responseData, err := base64.StdEncoding.DecodeString(data.Response)
	if err != nil {
		err = z_errs.ThrowInvalidArgument(err, "SAML-9klilyge7e", "Errors.Intent.ResponseInvalid")
		cmdErr := h.commands.FailIDPIntent(ctx, intent, err.Error())
		logging.WithFields("intent", intent.AggregateID).OnError(cmdErr).Error("failed to push failed event on idp intent")
		redirectToFailureURLErr(w, r, intent, err)
		return
	}

	assertion, err := sp.ServiceProvider.ParseXMLResponse(responseData, []string{intent.RequestID})
	if err != nil {
		err = z_errs.ThrowInvalidArgument(err, "SAML-hq6hwiyy03", "Errors.Intent.ResponseInvalid")
		cmdErr := h.commands.FailIDPIntent(ctx, intent, err.Error())
		logging.WithFields("intent", intent.AggregateID).OnError(cmdErr).Error("failed to push failed event on idp intent")
		redirectToFailureURLErr(w, r, intent, err)
		return
	}

	idpUser, err := saml2.ParseAssertionToUser(assertion)
	if err != nil {
		err = z_errs.ThrowInvalidArgument(err, "SAML-i07awpe1tm", "Errors.Intent.ResponseInvalid")
		cmdErr := h.commands.FailIDPIntent(ctx, intent, err.Error())
		logging.WithFields("intent", intent.AggregateID).OnError(cmdErr).Error("failed to push failed event on idp intent")
		redirectToFailureURLErr(w, r, intent, err)
		return
	}

	userID, err := h.checkExternalUser(ctx, intent.IDPID, assertion.Subject.NameID.Value)
	logging.WithFields("intent", intent.AggregateID).OnError(err).Error("could not check if idp user already exists")

	token, err := h.commands.SucceedSAMLIDPIntent(ctx, intent, idpUser, userID, data.Response)
	if err != nil {
		redirectToFailureURLErr(w, r, intent, z_errs.ThrowInternal(err, "IDP-JdD3g", "Errors.Intent.TokenCreationFailed"))
		return
	}
	redirectToSuccessURL(w, r, intent, token, userID)
}

func (h *Handler) handleCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data, err := h.parseCallbackRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	intent, err := h.commands.GetActiveIntent(ctx, data.State)
	if err != nil {
		if z_errs.IsNotFound(err) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		redirectToFailureURLErr(w, r, intent, err)
		return
	}

	// the provider might have returned an error
	if data.Error != "" {
		cmdErr := h.commands.FailIDPIntent(ctx, intent, reason(data.Error, data.ErrorDescription))
		logging.WithFields("intent", intent.AggregateID).OnError(cmdErr).Error("failed to push failed event on idp intent")
		redirectToFailureURL(w, r, intent, data.Error, data.ErrorDescription)
		return
	}

	provider, err := h.getProvider(ctx, intent.IDPID)
	if err != nil {
		cmdErr := h.commands.FailIDPIntent(ctx, intent, err.Error())
		logging.WithFields("intent", intent.AggregateID).OnError(cmdErr).Error("failed to push failed event on idp intent")
		redirectToFailureURLErr(w, r, intent, err)
		return
	}

	idpUser, idpSession, err := h.fetchIDPUserFromCode(ctx, provider, data.Code)
	if err != nil {
		cmdErr := h.commands.FailIDPIntent(ctx, intent, err.Error())
		logging.WithFields("intent", intent.AggregateID).OnError(cmdErr).Error("failed to push failed event on idp intent")
		redirectToFailureURLErr(w, r, intent, err)
		return
	}
	userID, err := h.checkExternalUser(ctx, intent.IDPID, idpUser.GetID())
	logging.WithFields("intent", intent.AggregateID).OnError(err).Error("could not check if idp user already exists")

	if userID == "" {
		userID, err = h.tryMigrateExternalUser(ctx, intent.IDPID, idpUser, idpSession)
		logging.WithFields("intent", intent.AggregateID).OnError(err).Error("migration check failed")
	}

	token, err := h.commands.SucceedIDPIntent(ctx, intent, idpUser, idpSession, userID)
	if err != nil {
		redirectToFailureURLErr(w, r, intent, z_errs.ThrowInternal(err, "IDP-JdD3g", "Errors.Intent.TokenCreationFailed"))
		return
	}
	redirectToSuccessURL(w, r, intent, token, userID)
}

func (h *Handler) tryMigrateExternalUser(ctx context.Context, idpID string, idpUser idp.User, idpSession idp.Session) (userID string, err error) {
	migration, ok := idpSession.(idp.SessionSupportsMigration)
	if !ok {
		return "", nil
	}
	previousID, err := migration.RetrievePreviousID()
	if err != nil || previousID == "" {
		return "", err
	}
	userID, err = h.checkExternalUser(ctx, idpID, previousID)
	if err != nil {
		return "", err
	}
	return userID, h.commands.MigrateUserIDP(ctx, userID, "", idpID, previousID, idpUser.GetID())
}

func (h *Handler) parseCallbackRequest(r *http.Request) (*externalIDPCallbackData, error) {
	data := new(externalIDPCallbackData)
	err := h.parser.Parse(r, data)
	if err != nil {
		return nil, err
	}
	if data.State == "" {
		return nil, z_errs.ThrowInvalidArgument(nil, "IDP-Hk38e", "Errors.Intent.StateMissing")
	}
	return data, nil
}

func redirectToSuccessURL(w http.ResponseWriter, r *http.Request, intent *command.IDPIntentWriteModel, token, userID string) {
	queries := intent.SuccessURL.Query()
	queries.Set(paramIntentID, intent.AggregateID)
	queries.Set(paramToken, token)
	if userID != "" {
		queries.Set(paramUserID, userID)
	}
	intent.SuccessURL.RawQuery = queries.Encode()
	http.Redirect(w, r, intent.SuccessURL.String(), http.StatusFound)
}

func redirectToFailureURLErr(w http.ResponseWriter, r *http.Request, i *command.IDPIntentWriteModel, err error) {
	msg := err.Error()
	var description string
	zErr := new(z_errs.CaosError)
	if errors.As(err, &zErr) {
		msg = zErr.GetID()
		description = zErr.GetMessage() // TODO: i18n?
	}
	redirectToFailureURL(w, r, i, msg, description)
}

func redirectToFailureURL(w http.ResponseWriter, r *http.Request, i *command.IDPIntentWriteModel, err, description string) {
	queries := i.FailureURL.Query()
	queries.Set(paramIntentID, i.AggregateID)
	queries.Set(paramError, err)
	queries.Set(paramErrorDescription, description)
	i.FailureURL.RawQuery = queries.Encode()
	http.Redirect(w, r, i.FailureURL.String(), http.StatusFound)
}

func (h *Handler) fetchIDPUserFromCode(ctx context.Context, identityProvider idp.Provider, code string) (user idp.User, idpTokens idp.Session, err error) {
	var session idp.Session
	switch provider := identityProvider.(type) {
	case *oauth.Provider:
		session = &oauth.Session{Provider: provider, Code: code}
	case *openid.Provider:
		session = &openid.Session{Provider: provider, Code: code}
	case *azuread.Provider:
		session = &azuread.Session{Session: &oauth.Session{Provider: provider.Provider, Code: code}}
	case *github.Provider:
		session = &oauth.Session{Provider: provider.Provider, Code: code}
	case *gitlab.Provider:
		session = &openid.Session{Provider: provider.Provider, Code: code}
	case *google.Provider:
		session = &openid.Session{Provider: provider.Provider, Code: code}
	case *jwt.Provider, *ldap.Provider, *saml2.Provider:
		return nil, nil, z_errs.ThrowInvalidArgument(nil, "IDP-52jmn", "Errors.ExternalIDP.IDPTypeNotImplemented")
	default:
		return nil, nil, z_errs.ThrowUnimplemented(nil, "IDP-SSDg", "Errors.ExternalIDP.IDPTypeNotImplemented")
	}

	user, err = session.FetchUser(ctx)
	if err != nil {
		return nil, nil, err
	}
	return user, session, nil
}

func (h *Handler) checkExternalUser(ctx context.Context, idpID, externalUserID string) (userID string, err error) {
	idQuery, err := query.NewIDPUserLinkIDPIDSearchQuery(idpID)
	if err != nil {
		return "", err
	}
	externalIDQuery, err := query.NewIDPUserLinksExternalIDSearchQuery(externalUserID)
	if err != nil {
		return "", err
	}
	queries := []query.SearchQuery{
		idQuery, externalIDQuery,
	}
	links, err := h.queries.IDPUserLinks(ctx, &query.IDPUserLinksSearchQuery{Queries: queries}, false)
	if err != nil {
		return "", err
	}
	if len(links.Links) != 1 {
		return "", nil
	}
	return links.Links[0].UserID, nil
}

func reason(err, description string) string {
	if description == "" {
		return err
	}
	return err + ": " + description
}
