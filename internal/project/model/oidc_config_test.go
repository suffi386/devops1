package model

import (
	"reflect"
	"testing"
)

func TestGetOIDCC1Compliance(t *testing.T) {
	type args struct {
		appType      OIDCApplicationType
		grantTypes   []OIDCGrantType
		authMethod   OIDCAuthMethodType
		redirectUris []string
	}
	type result struct {
		noneCompliant bool
	}
	tests := []struct {
		name   string
		args   args
		result result
	}{
		{
			name: "compliant native config",
			args: args{
				appType:    OIDCApplicationTypeNative,
				grantTypes: []OIDCGrantType{OIDCGrantTypeAuthorizationCode},
				authMethod: OIDCAuthMethodTypeNone,
				redirectUris: []string{
					"zitadel://auth/callback",
					"http://localhost/auth/callback",
				},
			},
			result: result{
				noneCompliant: false,
			},
		},
		{
			name: "none compliant native config",
			args: args{
				appType:    OIDCApplicationTypeNative,
				grantTypes: []OIDCGrantType{OIDCGrantTypeAuthorizationCode},
				authMethod: OIDCAuthMethodTypePost,
				redirectUris: []string{
					"https://zitadel.ch/auth/callback",
				},
			},
			result: result{
				noneCompliant: true,
			},
		},
		{
			name: "compliant web config",
			args: args{
				grantTypes: []OIDCGrantType{OIDCGrantTypeAuthorizationCode},
				redirectUris: []string{
					"https://zitadel.ch/auth/callback",
				},
			},
			result: result{
				noneCompliant: false,
			},
		},
		{
			name: "none compliant web config",
			args: args{
				grantTypes: []OIDCGrantType{OIDCGrantTypeAuthorizationCode},
				redirectUris: []string{
					"http://zitadel.ch/auth/callback",
				},
			},
			result: result{
				noneCompliant: true,
			},
		},
		{
			name: "compliant user agent config",
			args: args{
				appType:    OIDCApplicationTypeUserAgent,
				grantTypes: []OIDCGrantType{OIDCGrantTypeImplicit},
				authMethod: OIDCAuthMethodTypeNone,
				redirectUris: []string{
					"https://zitadel.ch/auth/callback",
				},
			},
			result: result{
				noneCompliant: false,
			},
		},
		{
			name: "none compliant  user agent config",
			args: args{
				appType:    OIDCApplicationTypeUserAgent,
				grantTypes: []OIDCGrantType{OIDCGrantTypeImplicit},
				authMethod: OIDCAuthMethodTypeBasic,
				redirectUris: []string{
					"https://zitadel.ch/auth/callback",
				},
			},
			result: result{
				noneCompliant: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetOIDCV1Compliance(tt.args.appType, tt.args.grantTypes, tt.args.authMethod, tt.args.redirectUris)
			if tt.result.noneCompliant != result.NoneCompliant {
				t.Errorf("got wrong result nonecompliant: expected: %v, actual: %v ", tt.result.noneCompliant, result.NoneCompliant)
			}
		})
	}
}

func TestGetOIDCNativeApplicationCompliance(t *testing.T) {
	type args struct {
		grantTypes   []OIDCGrantType
		authMethod   OIDCAuthMethodType
		redirectUris []string
	}
	type result struct {
		noneCompliant      bool
		complianceProblems []string
	}
	tests := []struct {
		name   string
		args   args
		result result
	}{
		{
			name: "compliant authorizationcode config",
			args: args{
				grantTypes: []OIDCGrantType{OIDCGrantTypeAuthorizationCode},
				authMethod: OIDCAuthMethodTypeNone,
				redirectUris: []string{
					"zitadel://auth/callback",
					"http://localhost/auth/callback",
				},
			},
			result: result{
				noneCompliant: false,
			},
		},
		{
			name: "none compliant authorizationcode config, auth method not none",
			args: args{
				grantTypes: []OIDCGrantType{OIDCGrantTypeAuthorizationCode},
				authMethod: OIDCAuthMethodTypePost,
				redirectUris: []string{
					"zitadel://auth/callback",
				},
			},
			result: result{
				noneCompliant: true,
				complianceProblems: []string{
					"Application.OIDC.V1.NotCompliant",
					"Application.OIDC.V1.Native.AuthMethodType.NotNone",
				},
			},
		},
		{
			name: "none compliant authorizationcode config, not localhost http",
			args: args{
				grantTypes: []OIDCGrantType{OIDCGrantTypeAuthorizationCode},
				authMethod: OIDCAuthMethodTypeNone,
				redirectUris: []string{
					"http://zitadel.ch/auth/callback",
				},
			},
			result: result{
				noneCompliant: true,
				complianceProblems: []string{
					"Application.OIDC.V1.NotCompliant",
					"Application.OIDC.V1.Native.RediredtUris.HttpOnlyForLocalhost",
				},
			},
		},
		{
			name: "none compliant authorizationcode config, https",
			args: args{
				grantTypes: []OIDCGrantType{OIDCGrantTypeAuthorizationCode},
				authMethod: OIDCAuthMethodTypeNone,
				redirectUris: []string{
					"https://zitadel.ch/auth/callback",
				},
			},
			result: result{
				noneCompliant: true,
				complianceProblems: []string{
					"Application.OIDC.V1.NotCompliant",
					"Application.OIDC.V1.Native.RediredtUris.HttpsNotAllowed",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetOIDCV1NativeApplicationCompliance(tt.args.grantTypes, tt.args.authMethod, tt.args.redirectUris)
			if tt.result.noneCompliant != result.NoneCompliant {
				t.Errorf("got wrong result nonecompliant: expected: %v, actual: %v ", tt.result.noneCompliant, result.NoneCompliant)
			}
			if tt.result.noneCompliant {
				if len(tt.result.complianceProblems) != len(result.Problems) {
					t.Errorf("got wrong result compliance problems len: expected: %v, actual: %v ", len(tt.result.complianceProblems), len(result.Problems))
				}
				if !reflect.DeepEqual(tt.result.complianceProblems, result.Problems) {
					t.Errorf("got wrong result compliance problems: expected: %v, actual: %v ", tt.result.complianceProblems, result.Problems)
				}
			}
		})
	}
}

func TestGetOIDCWebApplicationCompliance(t *testing.T) {
	type args struct {
		grantTypes   []OIDCGrantType
		redirectUris []string
	}
	type result struct {
		noneCompliant      bool
		complianceProblems []string
	}
	tests := []struct {
		name   string
		args   args
		result result
	}{
		{
			name: "compliant authorizationcode config",
			args: args{
				grantTypes: []OIDCGrantType{OIDCGrantTypeAuthorizationCode},
				redirectUris: []string{
					"https://zitadel.ch/auth/callback",
				},
			},
			result: result{
				noneCompliant: false,
			},
		},
		{
			name: "none compliant authorizationcode config, not https",
			args: args{
				grantTypes: []OIDCGrantType{OIDCGrantTypeAuthorizationCode},
				redirectUris: []string{
					"http://zitadel.ch/auth/callback",
				},
			},
			result: result{
				noneCompliant: true,
				complianceProblems: []string{
					"Application.OIDC.V1.NotCompliant",
					"Application.OIDC.V1.Web.RediredtUris.NotHttps",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetOIDCV1WebApplicationCompliance(tt.args.grantTypes, tt.args.redirectUris)
			if tt.result.noneCompliant != result.NoneCompliant {
				t.Errorf("got wrong result nonecompliant: expected: %v, actual: %v ", tt.result.noneCompliant, result.NoneCompliant)
			}
			if tt.result.noneCompliant {
				if len(tt.result.complianceProblems) != len(result.Problems) {
					t.Errorf("got wrong result compliance problems len: expected: %v, actual: %v ", len(tt.result.complianceProblems), len(result.Problems))
				}
				if !reflect.DeepEqual(tt.result.complianceProblems, result.Problems) {
					t.Errorf("got wrong result compliance problems: expected: %v, actual: %v ", tt.result.complianceProblems, result.Problems)
				}
			}
		})
	}
}

func TestGetOIDCUserAgentApplicationCompliance(t *testing.T) {
	type args struct {
		grantTypes   []OIDCGrantType
		authMethod   OIDCAuthMethodType
		redirectUris []string
	}
	type result struct {
		noneCompliant      bool
		complianceProblems []string
	}
	tests := []struct {
		name   string
		args   args
		result result
	}{
		{
			name: "compliant implicit config",
			args: args{
				grantTypes: []OIDCGrantType{OIDCGrantTypeImplicit},
				authMethod: OIDCAuthMethodTypeNone,
				redirectUris: []string{
					"https://zitadel.ch/auth/callback",
				},
			},
			result: result{
				noneCompliant: false,
			},
		},
		{
			name: "none compliant implicit config, not none",
			args: args{
				grantTypes: []OIDCGrantType{OIDCGrantTypeImplicit},
				authMethod: OIDCAuthMethodTypeBasic,
				redirectUris: []string{
					"https://zitadel.ch/auth/callback",
				},
			},
			result: result{
				noneCompliant: true,
				complianceProblems: []string{
					"Application.OIDC.V1.NotCompliant",
					"Application.OIDC.V1.UserAgent.Implicit.AuthMethodType.NotNone",
				},
			},
		},
		{
			name: "none compliant implicit config, not https",
			args: args{
				grantTypes: []OIDCGrantType{OIDCGrantTypeImplicit},
				authMethod: OIDCAuthMethodTypeNone,
				redirectUris: []string{
					"http://zitadel.ch/auth/callback",
				},
			},
			result: result{
				noneCompliant: true,
				complianceProblems: []string{
					"Application.OIDC.V1.NotCompliant",
					"Application.OIDC.V1.UserAgent.RediredtUris.NotHttps",
				},
			},
		},
		{
			name: "compliant authorizationcode config",
			args: args{
				grantTypes: []OIDCGrantType{OIDCGrantTypeAuthorizationCode},
				authMethod: OIDCAuthMethodTypeNone,
				redirectUris: []string{
					"https://zitadel.ch/auth/callback",
				},
			},
			result: result{
				noneCompliant: false,
			},
		},
		{
			name: "none compliant authorizationcode config, auth method not none",
			args: args{
				grantTypes: []OIDCGrantType{OIDCGrantTypeAuthorizationCode},
				authMethod: OIDCAuthMethodTypePost,
				redirectUris: []string{
					"https://zitadel.ch/auth/callback",
				},
			},
			result: result{
				noneCompliant: true,
				complianceProblems: []string{
					"Application.OIDC.V1.NotCompliant",
					"Application.OIDC.V1.UserAgent.AuthorizationCodeFlow.AuthMethodType.NotNone",
				},
			},
		},
		{
			name: "none compliant authorizationcode config, note https",
			args: args{
				grantTypes: []OIDCGrantType{OIDCGrantTypeAuthorizationCode},
				authMethod: OIDCAuthMethodTypeNone,
				redirectUris: []string{
					"http://zitadel.ch/auth/callback",
				},
			},
			result: result{
				noneCompliant: true,
				complianceProblems: []string{
					"Application.OIDC.V1.NotCompliant",
					"Application.OIDC.V1.UserAgent.RediredtUris.NotHttps",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetOIDCV1UserAgentApplicationCompliance(tt.args.grantTypes, tt.args.authMethod, tt.args.redirectUris)
			if tt.result.noneCompliant != result.NoneCompliant {
				t.Errorf("got wrong result nonecompliant: expected: %v, actual: %v ", tt.result.noneCompliant, result.NoneCompliant)
			}
			if tt.result.noneCompliant {
				if len(tt.result.complianceProblems) != len(result.Problems) {
					t.Errorf("got wrong result compliance problems len: expected: %v, actual: %v ", len(tt.result.complianceProblems), len(result.Problems))
				}
				if !reflect.DeepEqual(tt.result.complianceProblems, result.Problems) {
					t.Errorf("got wrong result compliance problems: expected: %v, actual: %v ", tt.result.complianceProblems, result.Problems)
				}
			}
		})
	}
}

func TestGetRequiredGrantTypes(t *testing.T) {
	type args struct {
		oidcConfig OIDCConfig
	}
	tests := []struct {
		name   string
		args   args
		result []OIDCGrantType
	}{
		{
			name: "oidc response type code",
			args: args{
				oidcConfig: OIDCConfig{
					ResponseTypes: []OIDCResponseType{OIDCResponseTypeCode},
				},
			},
			result: []OIDCGrantType{OIDCGrantTypeAuthorizationCode},
		},
		{
			name: "oidc response type id_token",
			args: args{
				oidcConfig: OIDCConfig{
					ResponseTypes: []OIDCResponseType{OIDCResponseTypeIDToken},
				},
			},
			result: []OIDCGrantType{OIDCGrantTypeImplicit},
		},
		{
			name: "oidc response type id_token and id_token token",
			args: args{
				oidcConfig: OIDCConfig{
					ResponseTypes: []OIDCResponseType{OIDCResponseTypeIDToken, OIDCResponseTypeIDTokenToken},
				},
			},
			result: []OIDCGrantType{OIDCGrantTypeImplicit},
		},
		{
			name: "oidc response type code, id_token and id_token token",
			args: args{
				oidcConfig: OIDCConfig{
					ResponseTypes: []OIDCResponseType{OIDCResponseTypeCode, OIDCResponseTypeIDToken, OIDCResponseTypeIDTokenToken},
				},
			},
			result: []OIDCGrantType{OIDCGrantTypeAuthorizationCode, OIDCGrantTypeImplicit},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.oidcConfig.getRequiredGrantTypes()
			if !reflect.DeepEqual(tt.result, result) {
				t.Errorf("got wrong result: expected: %v, actual: %v ", tt.result, result)
			}
		})
	}
}

func TestContainsOIDCGrantType(t *testing.T) {
	type args struct {
		grantTypes []OIDCGrantType
		grantType  OIDCGrantType
	}
	tests := []struct {
		name   string
		args   args
		result bool
	}{
		{
			name: "contains grant type",
			args: args{
				grantTypes: []OIDCGrantType{
					OIDCGrantTypeAuthorizationCode,
					OIDCGrantTypeImplicit,
				},
				grantType: OIDCGrantTypeImplicit,
			},
			result: true,
		},
		{
			name: "doesnt contain grant type",
			args: args{
				grantTypes: []OIDCGrantType{
					OIDCGrantTypeAuthorizationCode,
					OIDCGrantTypeRefreshToken,
				},
				grantType: OIDCGrantTypeImplicit,
			},
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := containsOIDCGrantType(tt.args.grantTypes, tt.args.grantType)
			if result != tt.result {
				t.Errorf("got wrong result: expected: %v, actual: %v ", tt.result, result)
			}
		})
	}
}

func TestUrlsAreHttps(t *testing.T) {
	type args struct {
		uris []string
	}
	tests := []struct {
		name   string
		args   args
		result bool
	}{
		{
			name: "only https uris",
			args: args{
				uris: []string{
					"https://zitadel.ch",
					"https://caos.ch",
				},
			},
			result: true,
		},
		{
			name: "http localhost uris",
			args: args{
				uris: []string{
					"https://zitadel.com",
					"http://localhost",
				},
			},
			result: false,
		},
		{
			name: "http not localhsot",
			args: args{
				uris: []string{
					"https://zitadel.com",
					"http://caos.ch",
				},
			},
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := urlsAreHttps(tt.args.uris)
			if result != tt.result {
				t.Errorf("got wrong result: expected: %v, actual: %v ", tt.result, result)
			}
		})
	}
}

func TestOnlyLocalhostIsHttp(t *testing.T) {
	type args struct {
		uris []string
	}
	tests := []struct {
		name   string
		args   args
		result bool
	}{

		{
			name: "http not localhost",
			args: args{
				uris: []string{
					"https://zitadel.com",
					"http://caos.ch",
				},
			},
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := onlyLocalhostIsHttp(tt.args.uris)
			if result != tt.result {
				t.Errorf("got wrong result: expected: %v, actual: %v ", tt.result, result)
			}
		})
	}
}
