package projection

import (
	"context"

	"github.com/caos/logging"
	"github.com/caos/zitadel/internal/domain"
	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore"
	"github.com/caos/zitadel/internal/eventstore/handler"
	"github.com/caos/zitadel/internal/eventstore/handler/crdb"
	"github.com/caos/zitadel/internal/repository/project"
	"github.com/lib/pq"
)

type AppProjection struct {
	crdb.StatementHandler
}

const (
	AppProjectionTable = "zitadel.projections.apps"
	AppAPISuffix       = "api_configs"
	AppOIDCSuffix      = "oidc_configs"
)

func NewAppProjection(ctx context.Context, config crdb.StatementHandlerConfig) *AppProjection {
	p := &AppProjection{}
	config.ProjectionName = AppProjectionTable
	config.Reducers = p.reducers()
	p.StatementHandler = crdb.NewStatementHandler(ctx, config)
	return p
}

func (p *AppProjection) reducers() []handler.AggregateReducer {
	return []handler.AggregateReducer{
		{
			Aggregate: project.AggregateType,
			EventRedusers: []handler.EventReducer{
				{
					Event:  project.ApplicationAddedType,
					Reduce: p.reduceAppAdded,
				},
				{
					Event:  project.ApplicationChangedType,
					Reduce: p.reduceAppChanged,
				},
				{
					Event:  project.ApplicationDeactivatedType,
					Reduce: p.reduceAppDeactivated,
				},
				{
					Event:  project.ApplicationReactivatedType,
					Reduce: p.reduceAppReactivated,
				},
				{
					Event:  project.ApplicationRemovedType,
					Reduce: p.reduceAppRemoved,
				},
				{
					Event:  project.APIConfigAddedType,
					Reduce: p.reduceAPIConfigAdded,
				},
				{
					Event:  project.APIConfigChangedType,
					Reduce: p.reduceAPIConfigChanged,
				},
				{
					Event:  project.APIConfigSecretChangedType,
					Reduce: p.reduceAPIConfigSecretChanged,
				},
				{
					Event:  project.OIDCConfigAddedType,
					Reduce: p.reduceOIDCConfigAdded,
				},
				{
					Event:  project.OIDCConfigChangedType,
					Reduce: p.reduceOIDCConfigChanged,
				},
				{
					Event:  project.OIDCConfigSecretChangedType,
					Reduce: p.reduceOIDCConfigSecretChanged,
				},
			},
		},
	}
}

const (
	AppColumnID            = "id"
	AppColumnName          = "name"
	AppColumnProjectID     = "project_id"
	AppColumnCreationDate  = "creation_date"
	AppColumnChangeDate    = "change_date"
	AppColumnResourceOwner = "resource_owner"
	AppColumnState         = "state"
	AppColumnSequence      = "sequence"

	APIConfigColumnAppID        = "app_id"
	APIConfigColumnClientID     = "client_id"
	APIConfigColumnClientSecret = "client_secert"
	APIConfigColumnAuthMethod   = "auth_method"

	OIDCConfigColumnAppID                    = "app_id"
	OIDCConfigColumnVersion                  = "version"
	OIDCConfigColumnClientID                 = "client_id"
	OIDCConfigColumnClientSecret             = "client_secret"
	OIDCConfigColumnRedirectUris             = "redirect_uris"
	OIDCConfigColumnResponseTypes            = "response_types"
	OIDCConfigColumnGrantTypes               = "grant_types"
	OIDCConfigColumnApplicationType          = "application_type"
	OIDCConfigColumnAuthMethodType           = "auth_method_type"
	OIDCConfigColumnPostLogoutRedirectUris   = "post_logout_redirect_uris"
	OIDCConfigColumnDevMode                  = "is_dev_mode"
	OIDCConfigColumnAccessTokenType          = "access_token_type"
	OIDCConfigColumnAccessTokenRoleAssertion = "access_token_role_assertion"
	OIDCConfigColumnIDTokenRoleAssertion     = "id_token_role_assertion"
	OIDCConfigColumnIDTokenUserinfoAssertion = "id_token_userinfo_assertion"
	OIDCConfigColumnClockSkew                = "clock_skew"
	OIDCConfigColumnAdditionalOrigins        = "additional_origins"
)

func (p *AppProjection) reduceAppAdded(event eventstore.EventReader) (*handler.Statement, error) {
	e, ok := event.(*project.ApplicationAddedEvent)
	if !ok {
		logging.LogWithFields("HANDL-OzK4m", "seq", event.Sequence(), "expectedType", project.ApplicationAddedType).Error("wrong event type")
		return nil, errors.ThrowInvalidArgument(nil, "HANDL-1xYE6", "reduce.wrong.event.type")
	}
	return crdb.NewCreateStatement(
		e,
		[]handler.Column{
			handler.NewCol(AppColumnID, e.AppID),
			handler.NewCol(AppColumnName, e.Name),
			handler.NewCol(AppColumnProjectID, e.Aggregate().ID),
			handler.NewCol(AppColumnCreationDate, e.CreationDate()),
			handler.NewCol(AppColumnChangeDate, e.CreationDate()),
			handler.NewCol(AppColumnResourceOwner, e.Aggregate().ResourceOwner),
			handler.NewCol(AppColumnState, domain.AppStateActive),
			handler.NewCol(AppColumnSequence, e.Sequence()),
		},
	), nil
}

func (p *AppProjection) reduceAppChanged(event eventstore.EventReader) (*handler.Statement, error) {
	e, ok := event.(*project.ApplicationChangedEvent)
	if !ok {
		logging.LogWithFields("HANDL-4Fjh2", "seq", event.Sequence(), "expectedType", project.ApplicationChangedType).Error("wrong event type")
		return nil, errors.ThrowInvalidArgument(nil, "HANDL-ZJ8JA", "reduce.wrong.event.type")
	}
	return crdb.NewUpdateStatement(
		e,
		[]handler.Column{
			handler.NewCol(AppColumnName, e.Name),
			handler.NewCol(AppColumnChangeDate, e.CreationDate()),
			handler.NewCol(AppColumnSequence, e.Sequence()),
		},
		[]handler.Condition{
			handler.NewCond(AppColumnID, e.AppID),
		},
	), nil
}

func (p *AppProjection) reduceAppDeactivated(event eventstore.EventReader) (*handler.Statement, error) {
	e, ok := event.(*project.ApplicationDeactivatedEvent)
	if !ok {
		logging.LogWithFields("HANDL-hZ9to", "seq", event.Sequence(), "expectedType", project.ApplicationDeactivatedType).Error("wrong event type")
		return nil, errors.ThrowInvalidArgument(nil, "HANDL-MVWxZ", "reduce.wrong.event.type")
	}
	return crdb.NewUpdateStatement(
		e,
		[]handler.Column{
			handler.NewCol(AppColumnState, domain.AppStateInactive),
			handler.NewCol(AppColumnChangeDate, e.CreationDate()),
			handler.NewCol(AppColumnSequence, e.Sequence()),
		},
		[]handler.Condition{
			handler.NewCond(AppColumnID, e.AppID),
		},
	), nil
}

func (p *AppProjection) reduceAppReactivated(event eventstore.EventReader) (*handler.Statement, error) {
	e, ok := event.(*project.ApplicationReactivatedEvent)
	if !ok {
		logging.LogWithFields("HANDL-AbK3B", "seq", event.Sequence(), "expectedType", project.ApplicationReactivatedType).Error("wrong event type")
		return nil, errors.ThrowInvalidArgument(nil, "HANDL-D0HZO", "reduce.wrong.event.type")
	}
	return crdb.NewUpdateStatement(
		e,
		[]handler.Column{
			handler.NewCol(AppColumnState, domain.AppStateActive),
			handler.NewCol(AppColumnChangeDate, e.CreationDate()),
			handler.NewCol(AppColumnSequence, e.Sequence()),
		},
		[]handler.Condition{
			handler.NewCond(AppColumnID, e.AppID),
		},
	), nil
}

func (p *AppProjection) reduceAppRemoved(event eventstore.EventReader) (*handler.Statement, error) {
	e, ok := event.(*project.ApplicationRemovedEvent)
	if !ok {
		logging.LogWithFields("HANDL-tdRId", "seq", event.Sequence(), "expectedType", project.ApplicationRemovedType).Error("wrong event type")
		return nil, errors.ThrowInvalidArgument(nil, "HANDL-Y99aq", "reduce.wrong.event.type")
	}
	return crdb.NewDeleteStatement(
		e,
		[]handler.Condition{
			handler.NewCond(AppColumnID, e.AppID),
		},
	), nil
}

func (p *AppProjection) reduceAPIConfigAdded(event eventstore.EventReader) (*handler.Statement, error) {
	e, ok := event.(*project.APIConfigAddedEvent)
	if !ok {
		logging.LogWithFields("HANDL-tdRId", "seq", event.Sequence(), "expectedType", project.APIConfigAddedType).Error("wrong event type")
		return nil, errors.ThrowInvalidArgument(nil, "HANDL-Y99aq", "reduce.wrong.event.type")
	}
	return crdb.NewMultiStatement(
		e,
		crdb.AddCreateStatement(
			[]handler.Column{
				handler.NewCol(APIConfigColumnAppID, e.AppID),
				handler.NewCol(APIConfigColumnClientID, e.ClientID),
				handler.NewCol(APIConfigColumnClientSecret, e.ClientSecret),
				handler.NewCol(APIConfigColumnAuthMethod, e.AuthMethodType),
			},
			crdb.WithTableSuffix(AppAPISuffix),
		),
		crdb.AddUpdateStatement(
			[]handler.Column{
				handler.NewCol(AppColumnChangeDate, e.CreationDate()),
				handler.NewCol(AppColumnSequence, e.Sequence()),
			},
			[]handler.Condition{
				handler.NewCond(AppColumnID, e.AppID),
			},
		),
	), nil
}

func (p *AppProjection) reduceAPIConfigChanged(event eventstore.EventReader) (*handler.Statement, error) {
	e, ok := event.(*project.APIConfigChangedEvent)
	if !ok {
		logging.LogWithFields("HANDL-C6b4f", "seq", event.Sequence(), "expectedType", project.APIConfigChangedType).Error("wrong event type")
		return nil, errors.ThrowInvalidArgument(nil, "HANDL-vnZKi", "reduce.wrong.event.type")
	}
	cols := make([]handler.Column, 0, 2)
	if e.ClientSecret != nil {
		cols = append(cols, handler.NewCol(APIConfigColumnClientSecret, e.ClientSecret))
	}
	if e.AuthMethodType != nil {
		cols = append(cols, handler.NewCol(APIConfigColumnAuthMethod, e.AuthMethodType))
	}
	if len(cols) == 0 {
		return crdb.NewNoOpStatement(e), nil
	}
	return crdb.NewMultiStatement(
		e,
		crdb.AddUpdateStatement(
			cols,
			[]handler.Condition{
				handler.NewCond(APIConfigColumnAppID, e.AppID),
			},
			crdb.WithTableSuffix(AppAPISuffix),
		),
		crdb.AddUpdateStatement(
			[]handler.Column{
				handler.NewCol(AppColumnChangeDate, e.CreationDate()),
				handler.NewCol(AppColumnSequence, e.Sequence()),
			},
			[]handler.Condition{
				handler.NewCond(AppColumnID, e.AppID),
			},
		),
	), nil
}

func (p *AppProjection) reduceAPIConfigSecretChanged(event eventstore.EventReader) (*handler.Statement, error) {
	e, ok := event.(*project.APIConfigSecretChangedEvent)
	if !ok {
		logging.LogWithFields("HANDL-dssSI", "seq", event.Sequence(), "expectedType", project.APIConfigSecretChangedType).Error("wrong event type")
		return nil, errors.ThrowInvalidArgument(nil, "HANDL-ttb0I", "reduce.wrong.event.type")
	}
	cols := make([]handler.Column, 0, 1)
	if e.ClientSecret != nil {
		cols = append(cols, handler.NewCol(APIConfigColumnClientSecret, e.ClientSecret))
	}
	if len(cols) == 0 {
		return crdb.NewNoOpStatement(e), nil
	}
	return crdb.NewMultiStatement(
		e,
		crdb.AddUpdateStatement(
			cols,
			[]handler.Condition{
				handler.NewCond(APIConfigColumnAppID, e.AppID),
			},
			crdb.WithTableSuffix(AppAPISuffix),
		),
		crdb.AddUpdateStatement(
			[]handler.Column{
				handler.NewCol(AppColumnChangeDate, e.CreationDate()),
				handler.NewCol(AppColumnSequence, e.Sequence()),
			},
			[]handler.Condition{
				handler.NewCond(AppColumnID, e.AppID),
			},
		),
	), nil
}

func (p *AppProjection) reduceOIDCConfigAdded(event eventstore.EventReader) (*handler.Statement, error) {
	e, ok := event.(*project.OIDCConfigAddedEvent)
	if !ok {
		logging.LogWithFields("HANDL-nlDQv", "seq", event.Sequence(), "expectedType", project.OIDCConfigAddedType).Error("wrong event type")
		return nil, errors.ThrowInvalidArgument(nil, "HANDL-GNHU1", "reduce.wrong.event.type")
	}
	return crdb.NewMultiStatement(
		e,
		crdb.AddCreateStatement(
			[]handler.Column{
				handler.NewCol(OIDCConfigColumnAppID, e.AppID),
				handler.NewCol(OIDCConfigColumnVersion, e.Version),
				handler.NewCol(OIDCConfigColumnClientID, e.ClientID),
				handler.NewCol(OIDCConfigColumnClientSecret, e.ClientSecret),
				handler.NewCol(OIDCConfigColumnRedirectUris, pq.StringArray(e.RedirectUris)),
				handler.NewCol(OIDCConfigColumnResponseTypes, pq.Array(e.ResponseTypes)),
				handler.NewCol(OIDCConfigColumnGrantTypes, pq.Array(e.GrantTypes)),
				handler.NewCol(OIDCConfigColumnApplicationType, e.ApplicationType),
				handler.NewCol(OIDCConfigColumnAuthMethodType, e.AuthMethodType),
				handler.NewCol(OIDCConfigColumnPostLogoutRedirectUris, pq.StringArray(e.PostLogoutRedirectUris)),
				handler.NewCol(OIDCConfigColumnDevMode, e.DevMode),
				handler.NewCol(OIDCConfigColumnAccessTokenType, e.AccessTokenType),
				handler.NewCol(OIDCConfigColumnAccessTokenRoleAssertion, e.AccessTokenRoleAssertion),
				handler.NewCol(OIDCConfigColumnIDTokenRoleAssertion, e.IDTokenRoleAssertion),
				handler.NewCol(OIDCConfigColumnIDTokenUserinfoAssertion, e.IDTokenUserinfoAssertion),
				handler.NewCol(OIDCConfigColumnClockSkew, e.ClockSkew),
				handler.NewCol(OIDCConfigColumnAdditionalOrigins, pq.StringArray(e.AdditionalOrigins)),
			},
			crdb.WithTableSuffix(AppOIDCSuffix),
		),
		crdb.AddUpdateStatement(
			[]handler.Column{
				handler.NewCol(AppColumnChangeDate, e.CreationDate()),
				handler.NewCol(AppColumnSequence, e.Sequence()),
			},
			[]handler.Condition{
				handler.NewCond(AppColumnID, e.AppID),
			},
		),
	), nil
}

func (p *AppProjection) reduceOIDCConfigChanged(event eventstore.EventReader) (*handler.Statement, error) {
	e, ok := event.(*project.OIDCConfigChangedEvent)
	if !ok {
		logging.LogWithFields("HANDL-nlDQv", "seq", event.Sequence(), "expectedType", project.OIDCConfigChangedType).Error("wrong event type")
		return nil, errors.ThrowInvalidArgument(nil, "HANDL-GNHU1", "reduce.wrong.event.type")
	}

	cols := make([]handler.Column, 0, 15)
	if e.Version != nil {
		cols = append(cols, handler.NewCol(OIDCConfigColumnVersion, *e.Version))
	}
	if e.RedirectUris != nil {
		cols = append(cols, handler.NewCol(OIDCConfigColumnRedirectUris, pq.StringArray(*e.RedirectUris)))
	}
	if e.ResponseTypes != nil {
		cols = append(cols, handler.NewCol(OIDCConfigColumnResponseTypes, pq.Array(*e.ResponseTypes)))
	}
	if e.GrantTypes != nil {
		cols = append(cols, handler.NewCol(OIDCConfigColumnGrantTypes, pq.Array(*e.GrantTypes)))
	}
	if e.ApplicationType != nil {
		cols = append(cols, handler.NewCol(OIDCConfigColumnApplicationType, *e.ApplicationType))
	}
	if e.AuthMethodType != nil {
		cols = append(cols, handler.NewCol(OIDCConfigColumnAuthMethodType, *e.AuthMethodType))
	}
	if e.PostLogoutRedirectUris != nil {
		cols = append(cols, handler.NewCol(OIDCConfigColumnPostLogoutRedirectUris, pq.StringArray(*e.PostLogoutRedirectUris)))
	}
	if e.DevMode != nil {
		cols = append(cols, handler.NewCol(OIDCConfigColumnDevMode, *e.DevMode))
	}
	if e.AccessTokenType != nil {
		cols = append(cols, handler.NewCol(OIDCConfigColumnAccessTokenType, *e.AccessTokenType))
	}
	if e.AccessTokenRoleAssertion != nil {
		cols = append(cols, handler.NewCol(OIDCConfigColumnAccessTokenRoleAssertion, *e.AccessTokenRoleAssertion))
	}
	if e.IDTokenRoleAssertion != nil {
		cols = append(cols, handler.NewCol(OIDCConfigColumnIDTokenRoleAssertion, *e.IDTokenRoleAssertion))
	}
	if e.IDTokenUserinfoAssertion != nil {
		cols = append(cols, handler.NewCol(OIDCConfigColumnIDTokenUserinfoAssertion, *e.IDTokenUserinfoAssertion))
	}
	if e.ClockSkew != nil {
		cols = append(cols, handler.NewCol(OIDCConfigColumnClockSkew, *e.ClockSkew))
	}
	if e.AdditionalOrigins != nil {
		cols = append(cols, handler.NewCol(OIDCConfigColumnAdditionalOrigins, pq.StringArray(*e.AdditionalOrigins)))
	}

	if len(cols) == 0 {
		return crdb.NewNoOpStatement(e), nil
	}

	return crdb.NewMultiStatement(
		e,
		crdb.AddUpdateStatement(
			cols,
			[]handler.Condition{
				handler.NewCond(OIDCConfigColumnAppID, e.AppID),
			},
			crdb.WithTableSuffix(AppOIDCSuffix),
		),
		crdb.AddUpdateStatement(
			[]handler.Column{
				handler.NewCol(AppColumnChangeDate, e.CreationDate()),
				handler.NewCol(AppColumnSequence, e.Sequence()),
			},
			[]handler.Condition{
				handler.NewCond(AppColumnID, e.AppID),
			},
		),
	), nil
}

func (p *AppProjection) reduceOIDCConfigSecretChanged(event eventstore.EventReader) (*handler.Statement, error) {
	e, ok := event.(*project.OIDCConfigSecretChangedEvent)
	if !ok {
		logging.LogWithFields("HANDL-nlDQv", "seq", event.Sequence(), "expectedType", project.OIDCConfigSecretChangedType).Error("wrong event type")
		return nil, errors.ThrowInvalidArgument(nil, "HANDL-GNHU1", "reduce.wrong.event.type")
	}
	return crdb.NewMultiStatement(
		e,
		crdb.AddUpdateStatement(
			[]handler.Column{
				handler.NewCol(OIDCConfigColumnClientSecret, e.ClientSecret),
			},
			[]handler.Condition{
				handler.NewCond(OIDCConfigColumnAppID, e.AppID),
			},
			crdb.WithTableSuffix(AppOIDCSuffix),
		),
		crdb.AddUpdateStatement(
			[]handler.Column{
				handler.NewCol(AppColumnChangeDate, e.CreationDate()),
				handler.NewCol(AppColumnSequence, e.Sequence()),
			},
			[]handler.Condition{
				handler.NewCond(AppColumnID, e.AppID),
			},
		),
	), nil
}
