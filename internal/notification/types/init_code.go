package types

import (
	"context"

	"github.com/caos/zitadel/internal/api/ui/login"
	"github.com/caos/zitadel/internal/crypto"
	"github.com/caos/zitadel/internal/domain"
	"github.com/caos/zitadel/internal/i18n"
	"github.com/caos/zitadel/internal/notification/channels/fs"
	"github.com/caos/zitadel/internal/notification/channels/log"
	"github.com/caos/zitadel/internal/notification/channels/smtp"
	"github.com/caos/zitadel/internal/notification/templates"
	"github.com/caos/zitadel/internal/query"
	es_model "github.com/caos/zitadel/internal/user/repository/eventsourcing/model"
	view_model "github.com/caos/zitadel/internal/user/repository/view/model"
)

type InitCodeEmailData struct {
	templates.TemplateData
	URL string
}

type UrlData struct {
	UserID      string
	Code        string
	PasswordSet bool
}

func SendUserInitCode(ctx context.Context, mailhtml string, translator *i18n.Translator, user *view_model.NotifyUser, code *es_model.InitUserCode, smtpConfig func(ctx context.Context) (*smtp.EmailConfig, error), getFileSystemProvider func(ctx context.Context) (*fs.FSConfig, error), getLogProvider func(ctx context.Context) (*log.LogConfig, error), alg crypto.EncryptionAlgorithm, colors *query.LabelPolicy, assetsPrefix string) error {
	codeString, err := crypto.DecryptString(code.Code, alg)
	if err != nil {
		return err
	}
	var origin string //TODO: build / retrieve origin
	url := login.InitUserLink(origin, user.ID, codeString, user.PasswordSet)
	var args = mapNotifyUserToArgs(user)
	args["Code"] = codeString

	initCodeData := &InitCodeEmailData{
		TemplateData: GetTemplateData(translator, args, assetsPrefix, url, domain.InitCodeMessageType, user.PreferredLanguage, colors),
		URL:          url,
	}
	template, err := templates.GetParsedTemplate(mailhtml, initCodeData)
	if err != nil {
		return err
	}
	return generateEmail(ctx, user, initCodeData.Subject, template, smtpConfig, getFileSystemProvider, getLogProvider, true)
}
