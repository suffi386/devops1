package handler

import (
	"errors"
	"fmt"
	"github.com/caos/logging"
	iam_model "github.com/caos/zitadel/internal/iam/model"
	"github.com/gorilla/csrf"
	"golang.org/x/text/language"
	"html/template"
	"net/http"
	"path"

	http_mw "github.com/caos/zitadel/internal/api/http/middleware"
	"github.com/caos/zitadel/internal/auth_request/model"
	caos_errs "github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/i18n"
	"github.com/caos/zitadel/internal/renderer"
)

const (
	tmplError = "error"
)

type Renderer struct {
	*renderer.Renderer
	pathPrefix string
}

func CreateRenderer(pathPrefix string, staticDir http.FileSystem, cookieName string, defaultLanguage language.Tag) *Renderer {
	r := &Renderer{
		pathPrefix: pathPrefix,
	}
	tmplMapping := map[string]string{
		tmplError:                  "error.html",
		tmplLogin:                  "login.html",
		tmplUserSelection:          "select_user.html",
		tmplPassword:               "password.html",
		tmplMfaVerify:              "mfa_verify.html",
		tmplMfaPrompt:              "mfa_prompt.html",
		tmplMfaInitVerify:          "mfa_init_verify.html",
		tmplMfaInitDone:            "mfa_init_done.html",
		tmplMailVerification:       "mail_verification.html",
		tmplMailVerified:           "mail_verified.html",
		tmplInitPassword:           "init_password.html",
		tmplInitPasswordDone:       "init_password_done.html",
		tmplInitUser:               "init_user.html",
		tmplInitUserDone:           "init_user_done.html",
		tmplPasswordResetDone:      "password_reset_done.html",
		tmplChangePassword:         "change_password.html",
		tmplChangePasswordDone:     "change_password_done.html",
		tmplRegisterOption:         "register_option.html",
		tmplRegister:               "register.html",
		tmplLogoutDone:             "logout_done.html",
		tmplRegisterOrg:            "register_org.html",
		tmplChangeUsername:         "change_username.html",
		tmplChangeUsernameDone:     "change_username_done.html",
		tmplLinkUsersDone:          "link_users_done.html",
		tmplExternalNotFoundOption: "external_not_found_option.html",
	}
	funcs := map[string]interface{}{
		"resourceUrl": func(file string) string {
			return path.Join(r.pathPrefix, EndpointResources, file)
		},
		"resourceThemeUrl": func(file, theme string) string {
			return path.Join(r.pathPrefix, EndpointResources, "themes", theme, file)
		},
		"loginUrl": func() string {
			return path.Join(r.pathPrefix, EndpointLogin)
		},
		"externalIDPAuthURL": func(authReqID, idpConfigID string) string {
			return path.Join(r.pathPrefix, fmt.Sprintf("%s?%s=%s&%s=%s", EndpointExternalLogin, queryAuthRequestID, authReqID, queryIDPConfigID, idpConfigID))
		},
		"externalIDPRegisterURL": func(authReqID, idpConfigID string) string {
			return path.Join(r.pathPrefix, fmt.Sprintf("%s?%s=%s&%s=%s", EndpointExternalRegister, queryAuthRequestID, authReqID, queryIDPConfigID, idpConfigID))
		},
		"registerUrl": func(id string) string {
			return path.Join(r.pathPrefix, fmt.Sprintf("%s?%s=%s", EndpointRegister, queryAuthRequestID, id))
		},
		"loginNameUrl": func() string {
			return path.Join(r.pathPrefix, EndpointLoginName)
		},
		"loginNameChangeUrl": func(id string) string {
			return path.Join(r.pathPrefix, fmt.Sprintf("%s?%s=%s", EndpointLoginName, queryAuthRequestID, id))
		},
		"userSelectionUrl": func() string {
			return path.Join(r.pathPrefix, EndpointUserSelection)
		},
		"passwordResetUrl": func(id string) string {
			return path.Join(r.pathPrefix, fmt.Sprintf("%s?%s=%s", EndpointPasswordReset, queryAuthRequestID, id))
		},
		"passwordUrl": func() string {
			return path.Join(r.pathPrefix, EndpointPassword)
		},
		"mfaVerifyUrl": func() string {
			return path.Join(r.pathPrefix, EndpointMfaVerify)
		},
		"mfaPromptUrl": func() string {
			return path.Join(r.pathPrefix, EndpointMfaPrompt)
		},
		"mfaPromptChangeUrl": func(id string, provider model.MfaType) string {
			return path.Join(r.pathPrefix, fmt.Sprintf("%s?%s=%s;%s=%v", EndpointMfaPrompt, queryAuthRequestID, id, "provider", provider))
		},
		"mfaInitVerifyUrl": func() string {
			return path.Join(r.pathPrefix, EndpointMfaInitVerify)
		},
		"mailVerificationUrl": func() string {
			return path.Join(r.pathPrefix, EndpointMailVerification)
		},
		"initPasswordUrl": func() string {
			return path.Join(r.pathPrefix, EndpointInitPassword)
		},
		"initUserUrl": func() string {
			return path.Join(r.pathPrefix, EndpointInitUser)
		},
		"changePasswordUrl": func() string {
			return path.Join(r.pathPrefix, EndpointChangePassword)
		},
		"registerOptionUrl": func() string {
			return path.Join(r.pathPrefix, EndpointRegisterOption)
		},
		"registrationUrl": func() string {
			return path.Join(r.pathPrefix, EndpointRegister)
		},
		"orgRegistrationUrl": func() string {
			return path.Join(r.pathPrefix, EndpointRegisterOrg)
		},
		"changeUsernameUrl": func() string {
			return path.Join(r.pathPrefix, EndpointChangeUsername)
		},
		"externalNotFoundOptionUrl": func() string {
			return path.Join(r.pathPrefix, EndpointExternalNotFoundOption)
		},
		"selectedLanguage": func(l string) bool {
			return false
		},
		"selectedGender": func(g int32) bool {
			return false
		},
		"hasExternalLogin": func() bool {
			return false
		},
	}
	var err error
	r.Renderer, err = renderer.NewRenderer(
		staticDir,
		tmplMapping, funcs,
		i18n.TranslatorConfig{DefaultLanguage: defaultLanguage, CookieName: cookieName},
	)
	logging.Log("APP-40tSoJ").OnError(err).WithError(err).Panic("error creating renderer")
	return r
}

func (l *Login) renderNextStep(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest) {
	userAgentID, _ := http_mw.UserAgentIDFromCtx(r.Context())
	authReq, err := l.authRepo.AuthRequestByID(r.Context(), authReq.ID, userAgentID)
	if err != nil {
		l.renderInternalError(w, r, authReq, caos_errs.ThrowInternal(nil, "APP-sio0W", "could not get authreq"))
	}
	if len(authReq.PossibleSteps) == 0 {
		l.renderInternalError(w, r, authReq, caos_errs.ThrowInternal(nil, "APP-9sdp4", "no possible steps"))
		return
	}
	l.chooseNextStep(w, r, authReq, 0, nil)
}

func (l *Login) renderError(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest, err error) {
	if authReq == nil || len(authReq.PossibleSteps) == 0 {
		l.renderInternalError(w, r, authReq, caos_errs.ThrowInternal(err, "APP-OVOiT", "no possible steps"))
		return
	}
	l.chooseNextStep(w, r, authReq, 0, err)
}

func (l *Login) chooseNextStep(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest, stepNumber int, err error) {
	switch step := authReq.PossibleSteps[stepNumber].(type) {
	case *model.LoginStep:
		if len(authReq.PossibleSteps) > 1 {
			l.chooseNextStep(w, r, authReq, 1, err)
			return
		}
		l.renderLogin(w, r, authReq, err)
	case *model.SelectUserStep:
		l.renderUserSelection(w, r, authReq, step)
	case *model.InitPasswordStep:
		l.renderInitPassword(w, r, authReq, authReq.UserID, "", err)
	case *model.PasswordStep:
		l.renderPassword(w, r, authReq, nil)
	case *model.MfaVerificationStep:
		l.renderMfaVerify(w, r, authReq, step, err)
	case *model.RedirectToCallbackStep:
		if len(authReq.PossibleSteps) > 1 {
			l.chooseNextStep(w, r, authReq, 1, err)
			return
		}
		l.redirectToCallback(w, r, authReq)
	case *model.ChangePasswordStep:
		l.renderChangePassword(w, r, authReq, err)
	case *model.VerifyEMailStep:
		l.renderMailVerification(w, r, authReq, "", err)
	case *model.MfaPromptStep:
		l.renderMfaPrompt(w, r, authReq, step, err)
	case *model.InitUserStep:
		l.renderInitUser(w, r, authReq, "", "", step.PasswordSet, nil)
	case *model.ChangeUsernameStep:
		l.renderChangeUsername(w, r, authReq, nil)
	case *model.LinkUsersStep:
		l.linkUsers(w, r, authReq, err)
	case *model.ExternalNotFoundOptionStep:
		l.renderExternalNotFoundOption(w, r, authReq, err)
	default:
		l.renderInternalError(w, r, authReq, caos_errs.ThrowInternal(nil, "APP-ds3QF", "step no possible"))
	}
}

func (l *Login) renderInternalError(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest, err error) {
	var msg string
	if err != nil {
		msg = err.Error()
	}
	data := l.getBaseData(r, authReq, "Error", "Internal", msg)
	l.renderer.RenderTemplate(w, r, l.renderer.Templates[tmplError], data, nil)
}

func (l *Login) getUserData(r *http.Request, authReq *model.AuthRequest, title string, errType, errMessage string) userData {
	return userData{
		baseData:    l.getBaseData(r, authReq, title, errType, errMessage),
		profileData: l.getProfileData(authReq),
		Linking:     len(authReq.LinkingUsers) > 0,
	}
}

func (l *Login) getBaseData(r *http.Request, authReq *model.AuthRequest, title string, errType, errMessage string) baseData {
	baseData := baseData{
		errorData: errorData{
			ErrType:    errType,
			ErrMessage: errMessage,
		},
		Lang:      l.renderer.Lang(r).String(),
		Title:     title,
		Theme:     l.getTheme(r),
		ThemeMode: l.getThemeMode(r),
		OrgID:     l.getOrgID(authReq),
		AuthReqID: getRequestID(authReq, r),
		CSRF:      csrf.TemplateField(r),
		Nonce:     http_mw.GetNonce(r),
	}
	if authReq != nil {
		baseData.LoginPolicy = authReq.LoginPolicy
		baseData.IDPProviders = authReq.AllowedExternalIDPs
	}
	return baseData
}

func (l *Login) getProfileData(authReq *model.AuthRequest) profileData {
	var loginName, displayName string
	if authReq != nil {
		loginName = authReq.LoginName
		displayName = authReq.DisplayName
	}
	return profileData{
		LoginName:   loginName,
		DisplayName: displayName,
	}
}

func (l *Login) getErrorMessage(r *http.Request, err error) (errMsg string) {
	caosErr := new(caos_errs.CaosError)
	if errors.As(err, &caosErr) {
		localized := l.renderer.LocalizeFromRequest(r, caosErr.Message, nil)
		return localized + " (" + caosErr.ID + ")"

	}
	return err.Error()
}

func (l *Login) getTheme(r *http.Request) string {
	return "zitadel" //TODO: impl
}

func (l *Login) getThemeMode(r *http.Request) string {
	return "" //TODO: impl
}

func (l *Login) getOrgID(authReq *model.AuthRequest) string {
	if authReq == nil {
		return ""
	}
	if authReq.UserOrgID != "" {
		return authReq.UserOrgID
	}
	if authReq.Request == nil {
		return ""
	}
	primaryDomain := authReq.GetScopeOrgPrimaryDomain()
	org, _ := l.authRepo.GetOrgByPrimaryDomain(primaryDomain)
	if org != nil {
		return org.ID
	}
	return ""
}

func getRequestID(authReq *model.AuthRequest, r *http.Request) string {
	if authReq != nil {
		return authReq.ID
	}
	return r.FormValue(queryAuthRequestID)
}

func (l *Login) csrfErrorHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := csrf.FailureReason(r)
		l.renderInternalError(w, r, nil, err)
	})
}

func (l *Login) cspErrorHandler(err error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l.renderInternalError(w, r, nil, err)
	})
}

type baseData struct {
	errorData
	Lang         string
	Title        string
	Theme        string
	ThemeMode    string
	OrgID        string
	AuthReqID    string
	CSRF         template.HTML
	Nonce        string
	LoginPolicy  *iam_model.LoginPolicyView
	IDPProviders []*iam_model.IDPProviderView
}

type errorData struct {
	ErrType    string
	ErrMessage string
}

type userData struct {
	baseData
	profileData
	PasswordChecked     string
	MfaProviders        []model.MfaType
	SelectedMfaProvider model.MfaType
	Linking             bool
}

type profileData struct {
	LoginName   string
	DisplayName string
}

type passwordData struct {
	baseData
	profileData
	PasswordPolicyDescription string
	MinLength                 uint64
	HasUppercase              string
	HasLowercase              string
	HasNumber                 string
	HasSymbol                 string
}

type userSelectionData struct {
	baseData
	Users   []model.UserSelection
	Linking bool
}

type mfaData struct {
	baseData
	profileData
	MfaProviders []model.MfaType
	MfaRequired  bool
}

type mfaVerifyData struct {
	baseData
	profileData
	MfaType model.MfaType
	otpData
}

type mfaDoneData struct {
	baseData
	profileData
	MfaType model.MfaType
}

type otpData struct {
	Url    string
	Secret string
	QrCode string
}
