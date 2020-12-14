package view

import (
	"fmt"

	caos_errs "github.com/caos/zitadel/internal/errors"
	iam_model "github.com/caos/zitadel/internal/iam/model"
	"github.com/caos/zitadel/internal/iam/repository/view/model"
	global_model "github.com/caos/zitadel/internal/model"
	"github.com/caos/zitadel/internal/view/repository"
	"github.com/jinzhu/gorm"
)

func GetMailTemplateByAggregateID(db *gorm.DB, table, aggregateID string) (*model.MailTemplateView, error) {
	template := new(model.MailTemplateView)
	userIDQuery := &model.MailTemplateSearchQuery{Key: iam_model.MailTemplateSearchKeyAggregateID, Value: aggregateID, Method: global_model.SearchMethodEquals}
	query := repository.PrepareGetByQuery(table, userIDQuery)
	err := query(db, template)
	if caos_errs.IsNotFound(err) {
		return nil, caos_errs.ThrowNotFound(nil, "VIEW-iPnmU", "Errors.IAM.MailTemplate.NotExisting")
	}
	//	sDec, _ := b64.StdEncoding.DecodeString(string(template.Template))
	fmt.Println(string(template.Template))
	return template, err
}

func PutMailTemplate(db *gorm.DB, table string, template *model.MailTemplateView) error {
	save := repository.PrepareSave(table)
	return save(db, template)
}

func DeleteMailTemplate(db *gorm.DB, table, aggregateID string) error {
	delete := repository.PrepareDeleteByKey(table, model.MailTemplateSearchKey(iam_model.MailTemplateSearchKeyAggregateID), aggregateID)

	return delete(db)
}
