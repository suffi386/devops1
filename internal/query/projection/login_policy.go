package projection

import (
	"context"

	"github.com/caos/logging"

	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore"
	"github.com/caos/zitadel/internal/eventstore/handler"
	"github.com/caos/zitadel/internal/eventstore/handler/crdb"
	"github.com/caos/zitadel/internal/repository/instance"
	"github.com/caos/zitadel/internal/repository/org"
	"github.com/caos/zitadel/internal/repository/policy"
)

type LoginPolicyProjection struct {
	crdb.StatementHandler
}

const (
	LoginPolicyTable = "zitadel.projections.login_policies"
)

func NewLoginPolicyProjection(ctx context.Context, config crdb.StatementHandlerConfig) *LoginPolicyProjection {
	p := new(LoginPolicyProjection)
	config.ProjectionName = LoginPolicyTable
	config.Reducers = p.reducers()
	p.StatementHandler = crdb.NewStatementHandler(ctx, config)
	return p
}

func (p *LoginPolicyProjection) reducers() []handler.AggregateReducer {
	return []handler.AggregateReducer{
		{
			Aggregate: org.AggregateType,
			EventRedusers: []handler.EventReducer{
				{
					Event:  org.LoginPolicyAddedEventType,
					Reduce: p.reduceLoginPolicyAdded,
				},
				{
					Event:  org.LoginPolicyChangedEventType,
					Reduce: p.reduceLoginPolicyChanged,
				},
				{
					Event:  org.LoginPolicyMultiFactorAddedEventType,
					Reduce: p.reduceMFAAdded,
				},
				{
					Event:  org.LoginPolicyMultiFactorRemovedEventType,
					Reduce: p.reduceMFARemoved,
				},
				{
					Event:  org.LoginPolicyRemovedEventType,
					Reduce: p.reduceLoginPolicyRemoved,
				},
				{
					Event:  org.LoginPolicySecondFactorAddedEventType,
					Reduce: p.reduce2FAAdded,
				},
				{
					Event:  org.LoginPolicySecondFactorRemovedEventType,
					Reduce: p.reduce2FARemoved,
				},
			},
		},
		{
			Aggregate: instance.AggregateType,
			EventRedusers: []handler.EventReducer{
				{
					Event:  instance.LoginPolicyAddedEventType,
					Reduce: p.reduceLoginPolicyAdded,
				},
				{
					Event:  instance.LoginPolicyChangedEventType,
					Reduce: p.reduceLoginPolicyChanged,
				},
				{
					Event:  instance.LoginPolicyMultiFactorAddedEventType,
					Reduce: p.reduceMFAAdded,
				},
				{
					Event:  instance.LoginPolicyMultiFactorRemovedEventType,
					Reduce: p.reduceMFARemoved,
				},
				{
					Event:  instance.LoginPolicySecondFactorAddedEventType,
					Reduce: p.reduce2FAAdded,
				},
				{
					Event:  instance.LoginPolicySecondFactorRemovedEventType,
					Reduce: p.reduce2FARemoved,
				},
			},
		},
	}
}

const (
	LoginPolicyIDCol                    = "aggregate_id"
	LoginPolicyCreationDateCol          = "creation_date"
	LoginPolicyChangeDateCol            = "change_date"
	LoginPolicySequenceCol              = "sequence"
	LoginPolicyAllowRegisterCol         = "allow_register"
	LoginPolicyAllowUsernamePasswordCol = "allow_username_password"
	LoginPolicyAllowExternalIDPsCol     = "allow_external_idps"
	LoginPolicyForceMFACol              = "force_mfa"
	LoginPolicy2FAsCol                  = "second_factors"
	LoginPolicyMFAsCol                  = "multi_factors"
	LoginPolicyPasswordlessTypeCol      = "passwordless_type"
	LoginPolicyIsDefaultCol             = "is_default"
	LoginPolicyHidePWResetCol           = "hide_password_reset"
	PasswordCheckLifetimeCol            = "password_check_lifetime"
	ExternalLoginCheckLifetimeCol       = "external_login_check_lifetime"
	MFAInitSkipLifetimeCol              = "mfa_init_skip_lifetime"
	SecondFactorCheckLifetimeCol        = "second_factor_check_lifetime"
	MultiFactorCheckLifetimeCol         = "multi_factor_check_lifetime"
)

func (p *LoginPolicyProjection) reduceLoginPolicyAdded(event eventstore.Event) (*handler.Statement, error) {
	var policyEvent policy.LoginPolicyAddedEvent
	var isDefault bool
	switch e := event.(type) {
	case *instance.LoginPolicyAddedEvent:
		policyEvent = e.LoginPolicyAddedEvent
		isDefault = true
	case *org.LoginPolicyAddedEvent:
		policyEvent = e.LoginPolicyAddedEvent
		isDefault = false
	default:
		logging.LogWithFields("HANDL-IW6So", "seq", event.Sequence(), "expectedTypes", []eventstore.EventType{org.LoginPolicyAddedEventType, instance.LoginPolicyAddedEventType}).Error("wrong event type")
		return nil, errors.ThrowInvalidArgument(nil, "HANDL-pYPxS", "reduce.wrong.event.type")
	}

	return crdb.NewCreateStatement(&policyEvent, []handler.Column{
		handler.NewCol(LoginPolicyIDCol, policyEvent.Aggregate().ID),
		handler.NewCol(LoginPolicyCreationDateCol, policyEvent.CreationDate()),
		handler.NewCol(LoginPolicyChangeDateCol, policyEvent.CreationDate()),
		handler.NewCol(LoginPolicySequenceCol, policyEvent.Sequence()),
		handler.NewCol(LoginPolicyAllowRegisterCol, policyEvent.AllowRegister),
		handler.NewCol(LoginPolicyAllowUsernamePasswordCol, policyEvent.AllowUserNamePassword),
		handler.NewCol(LoginPolicyAllowExternalIDPsCol, policyEvent.AllowExternalIDP),
		handler.NewCol(LoginPolicyForceMFACol, policyEvent.ForceMFA),
		handler.NewCol(LoginPolicyPasswordlessTypeCol, policyEvent.PasswordlessType),
		handler.NewCol(LoginPolicyIsDefaultCol, isDefault),
		handler.NewCol(LoginPolicyHidePWResetCol, policyEvent.HidePasswordReset),
		handler.NewCol(PasswordCheckLifetimeCol, policyEvent.PasswordCheckLifetime),
		handler.NewCol(ExternalLoginCheckLifetimeCol, policyEvent.ExternalLoginCheckLifetime),
		handler.NewCol(MFAInitSkipLifetimeCol, policyEvent.MFAInitSkipLifetime),
		handler.NewCol(SecondFactorCheckLifetimeCol, policyEvent.SecondFactorCheckLifetime),
		handler.NewCol(MultiFactorCheckLifetimeCol, policyEvent.MultiFactorCheckLifetime),
	}), nil
}

func (p *LoginPolicyProjection) reduceLoginPolicyChanged(event eventstore.Event) (*handler.Statement, error) {
	var policyEvent policy.LoginPolicyChangedEvent
	switch e := event.(type) {
	case *instance.LoginPolicyChangedEvent:
		policyEvent = e.LoginPolicyChangedEvent
	case *org.LoginPolicyChangedEvent:
		policyEvent = e.LoginPolicyChangedEvent
	default:
		logging.LogWithFields("HANDL-NIvFo", "seq", event.Sequence(), "expectedTypes", []eventstore.EventType{org.LoginPolicyChangedEventType, instance.LoginPolicyChangedEventType}).Error("wrong event type")
		return nil, errors.ThrowInvalidArgument(nil, "HANDL-BpaO6", "reduce.wrong.event.type")
	}

	cols := []handler.Column{
		handler.NewCol(LoginPolicyChangeDateCol, policyEvent.CreationDate()),
		handler.NewCol(LoginPolicySequenceCol, policyEvent.Sequence()),
	}
	if policyEvent.AllowRegister != nil {
		cols = append(cols, handler.NewCol(LoginPolicyAllowRegisterCol, *policyEvent.AllowRegister))
	}
	if policyEvent.AllowUserNamePassword != nil {
		cols = append(cols, handler.NewCol(LoginPolicyAllowUsernamePasswordCol, *policyEvent.AllowUserNamePassword))
	}
	if policyEvent.AllowExternalIDP != nil {
		cols = append(cols, handler.NewCol(LoginPolicyAllowExternalIDPsCol, *policyEvent.AllowExternalIDP))
	}
	if policyEvent.ForceMFA != nil {
		cols = append(cols, handler.NewCol(LoginPolicyForceMFACol, *policyEvent.ForceMFA))
	}
	if policyEvent.PasswordlessType != nil {
		cols = append(cols, handler.NewCol(LoginPolicyPasswordlessTypeCol, *policyEvent.PasswordlessType))
	}
	if policyEvent.HidePasswordReset != nil {
		cols = append(cols, handler.NewCol(LoginPolicyHidePWResetCol, *policyEvent.HidePasswordReset))
	}
	if policyEvent.PasswordCheckLifetime != nil {
		cols = append(cols, handler.NewCol(PasswordCheckLifetimeCol, *policyEvent.PasswordCheckLifetime))
	}
	if policyEvent.ExternalLoginCheckLifetime != nil {
		cols = append(cols, handler.NewCol(ExternalLoginCheckLifetimeCol, *policyEvent.ExternalLoginCheckLifetime))
	}
	if policyEvent.MFAInitSkipLifetime != nil {
		cols = append(cols, handler.NewCol(MFAInitSkipLifetimeCol, *policyEvent.MFAInitSkipLifetime))
	}
	if policyEvent.SecondFactorCheckLifetime != nil {
		cols = append(cols, handler.NewCol(SecondFactorCheckLifetimeCol, *policyEvent.SecondFactorCheckLifetime))
	}
	if policyEvent.MultiFactorCheckLifetime != nil {
		cols = append(cols, handler.NewCol(MultiFactorCheckLifetimeCol, *policyEvent.MultiFactorCheckLifetime))
	}

	return crdb.NewUpdateStatement(
		&policyEvent,
		cols,
		[]handler.Condition{
			handler.NewCond(LoginPolicyIDCol, policyEvent.Aggregate().ID),
		},
	), nil
}

func (p *LoginPolicyProjection) reduceMFAAdded(event eventstore.Event) (*handler.Statement, error) {
	var policyEvent policy.MultiFactorAddedEvent
	switch e := event.(type) {
	case *instance.LoginPolicyMultiFactorAddedEvent:
		policyEvent = e.MultiFactorAddedEvent
	case *org.LoginPolicyMultiFactorAddedEvent:
		policyEvent = e.MultiFactorAddedEvent
	default:
		logging.LogWithFields("HANDL-fYAHO", "seq", event.Sequence(), "expectedTypes", []eventstore.EventType{org.LoginPolicyMultiFactorAddedEventType, instance.LoginPolicyMultiFactorAddedEventType}).Error("wrong event type")
		return nil, errors.ThrowInvalidArgument(nil, "HANDL-WMhAV", "reduce.wrong.event.type")
	}

	return crdb.NewUpdateStatement(
		&policyEvent,
		[]handler.Column{
			handler.NewCol(LoginPolicyChangeDateCol, policyEvent.CreationDate()),
			handler.NewCol(LoginPolicySequenceCol, policyEvent.Sequence()),
			crdb.NewArrayAppendCol(LoginPolicyMFAsCol, policyEvent.MFAType),
		},
		[]handler.Condition{
			handler.NewCond(LoginPolicyIDCol, policyEvent.Aggregate().ID),
		},
	), nil
}

func (p *LoginPolicyProjection) reduceMFARemoved(event eventstore.Event) (*handler.Statement, error) {
	var policyEvent policy.MultiFactorRemovedEvent
	switch e := event.(type) {
	case *instance.LoginPolicyMultiFactorRemovedEvent:
		policyEvent = e.MultiFactorRemovedEvent
	case *org.LoginPolicyMultiFactorRemovedEvent:
		policyEvent = e.MultiFactorRemovedEvent
	default:
		logging.LogWithFields("HANDL-vtC31", "seq", event.Sequence(), "expectedTypes", []eventstore.EventType{org.LoginPolicyMultiFactorRemovedEventType, instance.LoginPolicyMultiFactorRemovedEventType}).Error("wrong event type")
		return nil, errors.ThrowInvalidArgument(nil, "HANDL-czU7n", "reduce.wrong.event.type")
	}

	return crdb.NewUpdateStatement(
		&policyEvent,
		[]handler.Column{
			handler.NewCol(LoginPolicyChangeDateCol, policyEvent.CreationDate()),
			handler.NewCol(LoginPolicySequenceCol, policyEvent.Sequence()),
			crdb.NewArrayRemoveCol(LoginPolicyMFAsCol, policyEvent.MFAType),
		},
		[]handler.Condition{
			handler.NewCond(LoginPolicyIDCol, policyEvent.Aggregate().ID),
		},
	), nil
}

func (p *LoginPolicyProjection) reduceLoginPolicyRemoved(event eventstore.Event) (*handler.Statement, error) {
	e, ok := event.(*org.LoginPolicyRemovedEvent)
	if !ok {
		logging.LogWithFields("HANDL-gF5q6", "seq", event.Sequence(), "expectedType", org.LoginPolicyRemovedEventType).Error("wrong event type")
		return nil, errors.ThrowInvalidArgument(nil, "HANDL-oRSvD", "reduce.wrong.event.type")
	}
	return crdb.NewDeleteStatement(
		e,
		[]handler.Condition{
			handler.NewCond(LoginPolicyIDCol, e.Aggregate().ID),
		},
	), nil
}

func (p *LoginPolicyProjection) reduce2FAAdded(event eventstore.Event) (*handler.Statement, error) {
	var policyEvent policy.SecondFactorAddedEvent
	switch e := event.(type) {
	case *instance.LoginPolicySecondFactorAddedEvent:
		policyEvent = e.SecondFactorAddedEvent
	case *org.LoginPolicySecondFactorAddedEvent:
		policyEvent = e.SecondFactorAddedEvent
	default:
		logging.LogWithFields("HANDL-dwadQ", "seq", event.Sequence(), "expectedTypes", []eventstore.EventType{org.LoginPolicySecondFactorAddedEventType, instance.LoginPolicySecondFactorAddedEventType}).Error("wrong event type")
		return nil, errors.ThrowInvalidArgument(nil, "HANDL-agB2E", "reduce.wrong.event.type")
	}

	return crdb.NewUpdateStatement(
		&policyEvent,
		[]handler.Column{
			handler.NewCol(LoginPolicyChangeDateCol, policyEvent.CreationDate()),
			handler.NewCol(LoginPolicySequenceCol, policyEvent.Sequence()),
			crdb.NewArrayAppendCol(LoginPolicy2FAsCol, policyEvent.MFAType),
		},
		[]handler.Condition{
			handler.NewCond(LoginPolicyIDCol, policyEvent.Aggregate().ID),
		},
	), nil
}

func (p *LoginPolicyProjection) reduce2FARemoved(event eventstore.Event) (*handler.Statement, error) {
	var policyEvent policy.SecondFactorRemovedEvent
	switch e := event.(type) {
	case *instance.LoginPolicySecondFactorRemovedEvent:
		policyEvent = e.SecondFactorRemovedEvent
	case *org.LoginPolicySecondFactorRemovedEvent:
		policyEvent = e.SecondFactorRemovedEvent
	default:
		logging.LogWithFields("HANDL-2IE8Y", "seq", event.Sequence(), "expectedTypes", []eventstore.EventType{org.LoginPolicySecondFactorRemovedEventType, instance.LoginPolicySecondFactorRemovedEventType}).Error("wrong event type")
		return nil, errors.ThrowInvalidArgument(nil, "HANDL-KYJvA", "reduce.wrong.event.type")
	}

	return crdb.NewUpdateStatement(
		&policyEvent,
		[]handler.Column{
			handler.NewCol(LoginPolicyChangeDateCol, policyEvent.CreationDate()),
			handler.NewCol(LoginPolicySequenceCol, policyEvent.Sequence()),
			crdb.NewArrayRemoveCol(LoginPolicy2FAsCol, policyEvent.MFAType),
		},
		[]handler.Condition{
			handler.NewCond(LoginPolicyIDCol, policyEvent.Aggregate().ID),
		},
	), nil
}
