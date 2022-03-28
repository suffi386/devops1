package login

import (
	"encoding/base64"
	"net/http"

	"github.com/caos/logging"
	"github.com/caos/zitadel/internal/api/authz"

	http_mw "github.com/caos/zitadel/internal/api/http/middleware"
	"github.com/caos/zitadel/internal/domain"
	"github.com/caos/zitadel/internal/eventstore/v1/models"
	"github.com/caos/zitadel/internal/query"
)

const (
	tmplPasswordlessRegistration     = "passwordlessregistration"
	tmplPasswordlessRegistrationDone = "passwordlessregistrationdone"
)

type passwordlessRegistrationData struct {
	webAuthNData
	Code                string
	CodeID              string
	UserID              string
	OrgID               string
	RequestPlatformType authPlatform
	Disabled            bool
}

type passwordlessRegistrationDoneDate struct {
	userData
	HideNextButton bool
}

type passwordlessRegistrationFormData struct {
	webAuthNFormData
	passwordlessRegistrationQueries
	TokenName string `schema:"name"`
}

type passwordlessRegistrationQueries struct {
	Code                string       `schema:"code"`
	CodeID              string       `schema:"codeID"`
	UserID              string       `schema:"userID"`
	OrgID               string       `schema:"orgID"`
	RequestPlatformType authPlatform `schema:"requestPlatformType"`
}

type authPlatform domain.AuthenticatorAttachment

func (a authPlatform) MarshalText() (text []byte, err error) {
	switch domain.AuthenticatorAttachment(a) {
	case domain.AuthenticatorAttachmentPlattform:
		return []byte("platform"), nil
	case domain.AuthenticatorAttachmentCrossPlattform:
		return []byte("crossPlatform"), nil
	default:
		return []byte("unspecified"), nil
	}
}

func (a *authPlatform) UnmarshalText(text []byte) (err error) {
	switch string(text) {
	case "platform",
		"1":
		*a = authPlatform(domain.AuthenticatorAttachmentPlattform)
	case "crossPlatform",
		"2":
		*a = authPlatform(domain.AuthenticatorAttachmentCrossPlattform)
	}
	return nil
}

func (l *Login) handlePasswordlessRegistration(w http.ResponseWriter, r *http.Request) {
	queries := new(passwordlessRegistrationQueries)
	err := l.parser.Parse(r, queries)
	l.renderPasswordlessRegistration(w, r, nil, queries.UserID, queries.OrgID, queries.CodeID, queries.Code, queries.RequestPlatformType, err)
}

func (l *Login) renderPasswordlessRegistration(w http.ResponseWriter, r *http.Request, authReq *domain.AuthRequest, userID, orgID, codeID, code string, requestedPlatformType authPlatform, err error) {
	var errID, errMessage, credentialData string
	var disabled bool
	if authReq != nil {
		userID = authReq.UserID
		orgID = authReq.UserOrgID
	}
	var webAuthNToken *domain.WebAuthNToken
	if err == nil {
		if authReq != nil {
			webAuthNToken, err = l.authRepo.BeginPasswordlessSetup(setContext(r.Context(), authReq.UserOrgID), authz.GetInstance(r.Context()).ID, userID, authReq.UserOrgID, domain.AuthenticatorAttachment(requestedPlatformType))
		} else {
			webAuthNToken, err = l.authRepo.BeginPasswordlessInitCodeSetup(setContext(r.Context(), orgID), authz.GetInstance(r.Context()).ID, userID, orgID, codeID, code, domain.AuthenticatorAttachment(requestedPlatformType))
		}
	}
	if err != nil {
		errID, errMessage = l.getErrorMessage(r, err)
		disabled = true
	}
	if webAuthNToken != nil {
		credentialData = base64.RawURLEncoding.EncodeToString(webAuthNToken.CredentialCreationData)
	}
	data := &passwordlessRegistrationData{
		webAuthNData{
			userData:               l.getUserData(r, authReq, "Login Passwordless", errID, errMessage),
			CredentialCreationData: credentialData,
		},
		code,
		codeID,
		userID,
		orgID,
		requestedPlatformType,
		disabled,
	}
	translator := l.getTranslator(authReq)
	if authReq == nil {
		policy, err := l.query.ActiveLabelPolicyByOrg(r.Context(), orgID)
		logging.Log("HANDL-XjWKE").OnError(err).Error("unable to get active label policy")
		data.LabelPolicy = labelPolicyToDomain(policy)

		translator, err = l.renderer.NewTranslator()
		if err == nil {
			texts, err := l.authRepo.GetLoginText(r.Context(), orgID)
			logging.Log("LOGIN-HJK4t").OnError(err).Warn("could not get custom texts")
			l.addLoginTranslations(translator, texts)
		}
	}
	l.renderer.RenderTemplate(w, r, translator, l.renderer.Templates[tmplPasswordlessRegistration], data, nil)
}

func labelPolicyToDomain(p *query.LabelPolicy) *domain.LabelPolicy {
	return &domain.LabelPolicy{
		ObjectRoot: models.ObjectRoot{
			AggregateID:   p.ID,
			Sequence:      p.Sequence,
			ResourceOwner: p.ResourceOwner,
			CreationDate:  p.CreationDate,
			ChangeDate:    p.ChangeDate,
		},
		State:               p.State,
		Default:             p.IsDefault,
		PrimaryColor:        p.Light.PrimaryColor,
		BackgroundColor:     p.Light.BackgroundColor,
		WarnColor:           p.Light.WarnColor,
		FontColor:           p.Light.FontColor,
		LogoURL:             p.Light.LogoURL,
		IconURL:             p.Light.IconURL,
		PrimaryColorDark:    p.Dark.PrimaryColor,
		BackgroundColorDark: p.Dark.BackgroundColor,
		WarnColorDark:       p.Dark.WarnColor,
		FontColorDark:       p.Dark.FontColor,
		LogoDarkURL:         p.Dark.LogoURL,
		IconDarkURL:         p.Dark.IconURL,
		Font:                p.FontURL,
		HideLoginNameSuffix: p.HideLoginNameSuffix,
		ErrorMsgPopup:       p.ShouldErrorPopup,
		DisableWatermark:    p.WatermarkDisabled,
	}
}

func (l *Login) handlePasswordlessRegistrationCheck(w http.ResponseWriter, r *http.Request) {
	formData := new(passwordlessRegistrationFormData)
	authReq, err := l.getAuthRequestAndParseData(r, formData)
	if err != nil {
		l.renderError(w, r, authReq, err)
		return
	}
	l.checkPasswordlessRegistration(w, r, authReq, formData)
}

func (l *Login) checkPasswordlessRegistration(w http.ResponseWriter, r *http.Request, authReq *domain.AuthRequest, formData *passwordlessRegistrationFormData) {
	credData, err := base64.URLEncoding.DecodeString(formData.CredentialData)
	if err != nil {
		l.renderPasswordlessRegistration(w, r, authReq, formData.UserID, formData.OrgID, formData.CodeID, formData.Code, formData.RequestPlatformType, err)
		return
	}
	userAgentID, _ := http_mw.UserAgentIDFromCtx(r.Context())
	if authReq != nil {
		err = l.authRepo.VerifyPasswordlessSetup(setContext(r.Context(), authReq.UserOrgID), formData.UserID, authReq.UserOrgID, userAgentID, formData.TokenName, credData)
	} else {
		err = l.authRepo.VerifyPasswordlessInitCodeSetup(setContext(r.Context(), formData.OrgID), formData.UserID, formData.OrgID, userAgentID, formData.TokenName, formData.CodeID, formData.Code, credData)
	}
	if err != nil {
		l.renderPasswordlessRegistration(w, r, authReq, formData.UserID, formData.OrgID, formData.CodeID, formData.Code, formData.RequestPlatformType, err)
		return
	}
	l.renderPasswordlessRegistrationDone(w, r, authReq, nil)
}

func (l *Login) renderPasswordlessRegistrationDone(w http.ResponseWriter, r *http.Request, authReq *domain.AuthRequest, err error) {
	var errID, errMessage string
	if err != nil {
		errID, errMessage = l.getErrorMessage(r, err)
	}
	data := passwordlessRegistrationDoneDate{
		userData:       l.getUserData(r, authReq, "Passwordless Registration Done", errID, errMessage),
		HideNextButton: authReq == nil,
	}
	l.renderer.RenderTemplate(w, r, l.getTranslator(authReq), l.renderer.Templates[tmplPasswordlessRegistrationDone], data, nil)
}
