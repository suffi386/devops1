package grpc

import (
	"encoding/json"

	"github.com/caos/logging"
	"github.com/caos/zitadel/internal/eventstore/models"
	"github.com/caos/zitadel/internal/model"
	proj_model "github.com/caos/zitadel/internal/project/model"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

func appFromModel(app *proj_model.Application) *Application {
	creationDate, err := ptypes.TimestampProto(app.CreationDate)
	logging.Log("GRPC-iejs3").OnError(err).Debug("unable to parse timestamp")

	changeDate, err := ptypes.TimestampProto(app.ChangeDate)
	logging.Log("GRPC-di7rw").OnError(err).Debug("unable to parse timestamp")

	return &Application{
		Id:           app.AppID,
		State:        appStateFromModel(app.State),
		CreationDate: creationDate,
		ChangeDate:   changeDate,
		Name:         app.Name,
		Sequence:     app.Sequence,
		AppConfig:    appConfigFromModel(app),
	}
}

func appConfigFromModel(app *proj_model.Application) isApplication_AppConfig {
	if app.Type == proj_model.APPTYPE_OIDC {
		return &Application_OidcConfig{
			OidcConfig: oidcConfigFromModel(app.OIDCConfig),
		}
	}
	return nil
}

func oidcConfigFromModel(config *proj_model.OIDCConfig) *OIDCConfig {
	return &OIDCConfig{
		RedirectUris:           config.RedirectUris,
		ResponseTypes:          oidcResponseTypesFromModel(config.ResponseTypes),
		GrantTypes:             oidcGrantTypesFromModel(config.GrantTypes),
		ApplicationType:        oidcApplicationTypeFromModel(config.ApplicationType),
		ClientId:               config.ClientID,
		ClientSecret:           config.ClientSecretString,
		AuthMethodType:         oidcAuthMethodTypeFromModel(config.AuthMethodType),
		PostLogoutRedirectUris: config.PostLogoutRedirectUris,
	}
}

func oidcConfigFromApplicationViewModel(app *proj_model.ApplicationView) *OIDCConfig {
	return &OIDCConfig{
		RedirectUris:           app.OIDCRedirectUris,
		ResponseTypes:          oidcResponseTypesFromModel(app.OIDCResponseTypes),
		GrantTypes:             oidcGrantTypesFromModel(app.OIDCGrantTypes),
		ApplicationType:        oidcApplicationTypeFromModel(app.OIDCApplicationType),
		ClientId:               app.OIDCClientID,
		AuthMethodType:         oidcAuthMethodTypeFromModel(app.OIDCAuthMethodType),
		PostLogoutRedirectUris: app.OIDCPostLogoutRedirectUris,
	}
}

func oidcAppCreateToModel(app *OIDCApplicationCreate) *proj_model.Application {
	return &proj_model.Application{
		ObjectRoot: models.ObjectRoot{
			AggregateID: app.ProjectId,
		},
		Name: app.Name,
		Type: proj_model.APPTYPE_OIDC,
		OIDCConfig: &proj_model.OIDCConfig{
			RedirectUris:           app.RedirectUris,
			ResponseTypes:          oidcResponseTypesToModel(app.ResponseTypes),
			GrantTypes:             oidcGrantTypesToModel(app.GrantTypes),
			ApplicationType:        oidcApplicationTypeToModel(app.ApplicationType),
			AuthMethodType:         oidcAuthMethodTypeToModel(app.AuthMethodType),
			PostLogoutRedirectUris: app.PostLogoutRedirectUris,
		},
	}
}

func appUpdateToModel(app *ApplicationUpdate) *proj_model.Application {
	return &proj_model.Application{
		ObjectRoot: models.ObjectRoot{
			AggregateID: app.ProjectId,
		},
		AppID: app.Id,
		Name:  app.Name,
	}
}

func oidcConfigUpdateToModel(app *OIDCConfigUpdate) *proj_model.OIDCConfig {
	return &proj_model.OIDCConfig{
		ObjectRoot: models.ObjectRoot{
			AggregateID: app.ProjectId,
		},
		AppID:                  app.ApplicationId,
		RedirectUris:           app.RedirectUris,
		ResponseTypes:          oidcResponseTypesToModel(app.ResponseTypes),
		GrantTypes:             oidcGrantTypesToModel(app.GrantTypes),
		ApplicationType:        oidcApplicationTypeToModel(app.ApplicationType),
		AuthMethodType:         oidcAuthMethodTypeToModel(app.AuthMethodType),
		PostLogoutRedirectUris: app.PostLogoutRedirectUris,
	}
}

func applicationSearchRequestsToModel(request *ApplicationSearchRequest) *proj_model.ApplicationSearchRequest {
	return &proj_model.ApplicationSearchRequest{
		Offset:  request.Offset,
		Limit:   request.Limit,
		Queries: applicationSearchQueriesToModel(request.ProjectId, request.Queries),
	}
}

func applicationSearchQueriesToModel(projectID string, queries []*ApplicationSearchQuery) []*proj_model.ApplicationSearchQuery {
	converted := make([]*proj_model.ApplicationSearchQuery, len(queries)+1)
	for i, q := range queries {
		converted[i] = applicationSearchQueryToModel(q)
	}
	converted[len(queries)] = &proj_model.ApplicationSearchQuery{Key: proj_model.APPLICATIONSEARCHKEY_PROJECT_ID, Method: model.SEARCHMETHOD_EQUALS, Value: projectID}

	return converted
}

func applicationSearchQueryToModel(query *ApplicationSearchQuery) *proj_model.ApplicationSearchQuery {
	return &proj_model.ApplicationSearchQuery{
		Key:    applicationSearchKeyToModel(query.Key),
		Method: searchMethodToModel(query.Method),
		Value:  query.Value,
	}
}

func applicationSearchKeyToModel(key ApplicationSearchKey) proj_model.ApplicationSearchKey {
	switch key {
	case ApplicationSearchKey_APPLICATIONSEARCHKEY_APP_NAME:
		return proj_model.APPLICATIONSEARCHKEY_NAME
	default:
		return proj_model.APPLICATIONSEARCHKEY_UNSPECIFIED
	}
}

func applicationSearchResponseFromModel(response *proj_model.ApplicationSearchResponse) *ApplicationSearchResponse {
	return &ApplicationSearchResponse{
		Offset:      response.Offset,
		Limit:       response.Limit,
		TotalResult: response.TotalResult,
		Result:      applicationViewsFromModel(response.Result),
	}
}

func applicationViewsFromModel(apps []*proj_model.ApplicationView) []*ApplicationView {
	converted := make([]*ApplicationView, len(apps))
	for i, app := range apps {
		converted[i] = applicationViewFromModel(app)
	}
	return converted
}

func applicationViewFromModel(application *proj_model.ApplicationView) *ApplicationView {
	creationDate, err := ptypes.TimestampProto(application.CreationDate)
	logging.Log("GRPC-lo9sw").OnError(err).Debug("unable to parse timestamp")

	changeDate, err := ptypes.TimestampProto(application.ChangeDate)
	logging.Log("GRPC-8uwsd").OnError(err).Debug("unable to parse timestamp")

	converted := &ApplicationView{
		Id:           application.ID,
		State:        appStateFromModel(application.State),
		CreationDate: creationDate,
		ChangeDate:   changeDate,
		Name:         application.Name,
		Sequence:     application.Sequence,
	}
	if application.IsOIDC {
		converted.AppConfig = &ApplicationView_OidcConfig{
			OidcConfig: oidcConfigFromApplicationViewModel(application),
		}
	}
	return converted
}

func appStateFromModel(state proj_model.AppState) AppState {
	switch state {
	case proj_model.APPSTATE_ACTIVE:
		return AppState_APPSTATE_ACTIVE
	case proj_model.APPSTATE_INACTIVE:
		return AppState_APPSTATE_INACTIVE
	default:
		return AppState_APPSTATE_UNSPECIFIED
	}
}

func oidcResponseTypesToModel(responseTypes []OIDCResponseType) []proj_model.OIDCResponseType {
	if responseTypes == nil || len(responseTypes) == 0 {
		return []proj_model.OIDCResponseType{proj_model.OIDCRESPONSETYPE_CODE}
	}
	oidcResponseTypes := make([]proj_model.OIDCResponseType, len(responseTypes))

	for i, responseType := range responseTypes {
		switch responseType {
		case OIDCResponseType_OIDCRESPONSETYPE_CODE:
			oidcResponseTypes[i] = proj_model.OIDCRESPONSETYPE_CODE
		case OIDCResponseType_OIDCRESPONSETYPE_ID_TOKEN:
			oidcResponseTypes[i] = proj_model.OIDCRESPONSETYPE_ID_TOKEN
		case OIDCResponseType_OIDCRESPONSETYPE_TOKEN:
			oidcResponseTypes[i] = proj_model.OIDCRESPONSETYPE_TOKEN
		}
	}

	return oidcResponseTypes
}

func oidcResponseTypesFromModel(responseTypes []proj_model.OIDCResponseType) []OIDCResponseType {
	oidcResponseTypes := make([]OIDCResponseType, len(responseTypes))

	for i, responseType := range responseTypes {
		switch responseType {
		case proj_model.OIDCRESPONSETYPE_CODE:
			oidcResponseTypes[i] = OIDCResponseType_OIDCRESPONSETYPE_CODE
		case proj_model.OIDCRESPONSETYPE_ID_TOKEN:
			oidcResponseTypes[i] = OIDCResponseType_OIDCRESPONSETYPE_ID_TOKEN
		case proj_model.OIDCRESPONSETYPE_TOKEN:
			oidcResponseTypes[i] = OIDCResponseType_OIDCRESPONSETYPE_TOKEN
		}
	}

	return oidcResponseTypes
}

func oidcGrantTypesToModel(grantTypes []OIDCGrantType) []proj_model.OIDCGrantType {
	if grantTypes == nil || len(grantTypes) == 0 {
		return []proj_model.OIDCGrantType{proj_model.OIDCGRANTTYPE_AUTHORIZATION_CODE}
	}
	oidcGrantTypes := make([]proj_model.OIDCGrantType, len(grantTypes))

	for i, grantType := range grantTypes {
		switch grantType {
		case OIDCGrantType_OIDCGRANTTYPE_AUTHORIZATION_CODE:
			oidcGrantTypes[i] = proj_model.OIDCGRANTTYPE_AUTHORIZATION_CODE
		case OIDCGrantType_OIDCGRANTTYPE_IMPLICIT:
			oidcGrantTypes[i] = proj_model.OIDCGRANTTYPE_IMPLICIT
		case OIDCGrantType_OIDCGRANTTYPE_REFRESH_TOKEN:
			oidcGrantTypes[i] = proj_model.OIDCGRANTTYPE_REFRESH_TOKEN
		}
	}
	return oidcGrantTypes
}

func oidcGrantTypesFromModel(grantTypes []proj_model.OIDCGrantType) []OIDCGrantType {
	oidcGrantTypes := make([]OIDCGrantType, len(grantTypes))

	for i, grantType := range grantTypes {
		switch grantType {
		case proj_model.OIDCGRANTTYPE_AUTHORIZATION_CODE:
			oidcGrantTypes[i] = OIDCGrantType_OIDCGRANTTYPE_AUTHORIZATION_CODE
		case proj_model.OIDCGRANTTYPE_IMPLICIT:
			oidcGrantTypes[i] = OIDCGrantType_OIDCGRANTTYPE_IMPLICIT
		case proj_model.OIDCGRANTTYPE_REFRESH_TOKEN:
			oidcGrantTypes[i] = OIDCGrantType_OIDCGRANTTYPE_REFRESH_TOKEN
		}
	}
	return oidcGrantTypes
}

func oidcApplicationTypeToModel(appType OIDCApplicationType) proj_model.OIDCApplicationType {
	switch appType {
	case OIDCApplicationType_OIDCAPPLICATIONTYPE_WEB:
		return proj_model.OIDCAPPLICATIONTYPE_WEB
	case OIDCApplicationType_OIDCAPPLICATIONTYPE_USER_AGENT:
		return proj_model.OIDCAPPLICATIONTYPE_USER_AGENT
	case OIDCApplicationType_OIDCAPPLICATIONTYPE_NATIVE:
		return proj_model.OIDCAPPLICATIONTYPE_NATIVE
	}
	return proj_model.OIDCAPPLICATIONTYPE_WEB
}

func oidcApplicationTypeFromModel(appType proj_model.OIDCApplicationType) OIDCApplicationType {
	switch appType {
	case proj_model.OIDCAPPLICATIONTYPE_WEB:
		return OIDCApplicationType_OIDCAPPLICATIONTYPE_WEB
	case proj_model.OIDCAPPLICATIONTYPE_USER_AGENT:
		return OIDCApplicationType_OIDCAPPLICATIONTYPE_USER_AGENT
	case proj_model.OIDCAPPLICATIONTYPE_NATIVE:
		return OIDCApplicationType_OIDCAPPLICATIONTYPE_NATIVE
	default:
		return OIDCApplicationType_OIDCAPPLICATIONTYPE_WEB
	}
}

func oidcAuthMethodTypeToModel(authType OIDCAuthMethodType) proj_model.OIDCAuthMethodType {
	switch authType {
	case OIDCAuthMethodType_OIDCAUTHMETHODTYPE_BASIC:
		return proj_model.OIDCAUTHMETHODTYPE_BASIC
	case OIDCAuthMethodType_OIDCAUTHMETHODTYPE_POST:
		return proj_model.OIDCAUTHMETHODTYPE_POST
	case OIDCAuthMethodType_OIDCAUTHMETHODTYPE_NONE:
		return proj_model.OIDCAUTHMETHODTYPE_NONE
	default:
		return proj_model.OIDCAUTHMETHODTYPE_BASIC
	}
}

func oidcAuthMethodTypeFromModel(authType proj_model.OIDCAuthMethodType) OIDCAuthMethodType {
	switch authType {
	case proj_model.OIDCAUTHMETHODTYPE_BASIC:
		return OIDCAuthMethodType_OIDCAUTHMETHODTYPE_BASIC
	case proj_model.OIDCAUTHMETHODTYPE_POST:
		return OIDCAuthMethodType_OIDCAUTHMETHODTYPE_POST
	case proj_model.OIDCAUTHMETHODTYPE_NONE:
		return OIDCAuthMethodType_OIDCAUTHMETHODTYPE_NONE
	default:
		return OIDCAuthMethodType_OIDCAUTHMETHODTYPE_BASIC
	}
}

func appChangesToResponse(response *proj_model.ApplicationChanges, offset uint64, limit uint64) (_ *Changes) {
	return &Changes{
		Limit:   limit,
		Offset:  offset,
		Changes: appChangesToMgtAPI(response),
	}
}

func appChangesToMgtAPI(changes *proj_model.ApplicationChanges) (_ []*Change) {
	result := make([]*Change, len(changes.Changes))

	for i, change := range changes.Changes {
		b, err := json.Marshal(change.Data)
		data := &structpb.Struct{}
		err = protojson.Unmarshal(b, data)
		if err != nil {
		}
		result[i] = &Change{
			ChangeDate: change.ChangeDate,
			EventType:  change.EventType,
			Sequence:   change.Sequence,
			Data:       data,
		}
	}

	return result
}
