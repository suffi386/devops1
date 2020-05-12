package systemdefaults

import (
	"github.com/caos/zitadel/internal/crypto"
	pol "github.com/caos/zitadel/internal/policy"
)

type SystemDefaults struct {
	SecretGenerators    SecretGenerators
	UserVerificationKey *crypto.KeyConfig
	Multifactors        MultifactorConfig
	DefaultPolicies     DefaultPolicies
}

type SecretGenerators struct {
	PasswordSaltCost         int
	ClientSecretGenerator    crypto.GeneratorConfig
	InitializeUserCode       crypto.GeneratorConfig
	EmailVerificationCode    crypto.GeneratorConfig
	PhoneVerificationCode    crypto.GeneratorConfig
	PasswordVerificationCode crypto.GeneratorConfig
}

type MultifactorConfig struct {
	OTP OTPConfig
}

type OTPConfig struct {
	Issuer          string
	VerificationKey *crypto.KeyConfig
}

type DefaultPolicies struct {
	Age        pol.PasswordAgePolicyDefault
	Complexity pol.PasswordComplexityPolicyDefault
	Lockout    pol.PasswordLockoutPolicyDefault
}
