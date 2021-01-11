package domain

import (
	"github.com/caos/zitadel/internal/eventstore/models"
)

const (
	ResourceOwnerIAM = "IAM"
)

type IAM struct {
	models.ObjectRoot

	GlobalOrgID                     string
	IAMProjectID                    string
	SetUpDone                       Step
	SetUpStarted                    Step
	Members                         []*Member
	IDPs                            []*IDPConfig
	DefaultLoginPolicy              *LoginPolicy
	DefaultLabelPolicy              *LabelPolicy
	DefaultOrgIAMPolicy             *OrgIAMPolicy
	DefaultPasswordComplexityPolicy *PasswordComplexityPolicy
	DefaultPasswordAgePolicy        *PasswordAgePolicy
	DefaultPasswordLockoutPolicy    *PasswordLockoutPolicy
}
