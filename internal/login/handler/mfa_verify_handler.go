package handler

import (
	"net/http"

	"github.com/caos/zitadel/internal/auth_request/model"
)

const (
	tmplMfaVerify = "mfaverify"
)

type mfaVerifyFormData struct {
	MfaType model.MfaType `schema:"mfaType"`
	Code    string        `schema:"code"`
}

func (l *Login) handleMfaVerify(w http.ResponseWriter, r *http.Request) {
	data := new(mfaVerifyFormData)
	authReq, err := l.getAuthRequestAndParseData(r, data)
	if err != nil {
		l.renderError(w, r, authReq, err)
		return
	}
	if data.MfaType == model.MfaTypeOTP {
		err = l.authRepo.VerifyMfaOTP(setContext(r.Context(), authReq.UserOrgID), authReq.ID, authReq.UserID, data.Code, model.BrowserInfoFromRequest(r))
	}
	if err != nil {
		l.renderError(w, r, authReq, err)
		return
	}
	l.renderNextStep(w, r, authReq)
}

func (l *Login) renderMfaVerify(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest, verificationStep *model.MfaVerificationStep, err error) {
	var errType, errMessage string
	if err != nil {
		errMessage = l.getErrorMessage(r, err)
	}
	data := userData{
		baseData:  l.getBaseData(r, authReq, "Mfa Verify", errType, errMessage),
		LoginName: authReq.LoginName,
	}
	if verificationStep != nil {
		data.MfaProviders = verificationStep.MfaProviders
		data.SelectedMfaProvider = verificationStep.MfaProviders[0]
	}
	l.renderer.RenderTemplate(w, r, l.renderer.Templates[tmplMfaVerify], data, nil)
}
