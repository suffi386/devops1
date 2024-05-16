package handler

import (
	"context"
	"time"

	auth_view "github.com/zitadel/zitadel/internal/auth/repository/eventsourcing/view"
	"github.com/zitadel/zitadel/internal/crypto"
	"github.com/zitadel/zitadel/internal/eventstore"
	"github.com/zitadel/zitadel/internal/eventstore/handler/v2"
	query2 "github.com/zitadel/zitadel/internal/query"
	"github.com/zitadel/zitadel/internal/repository/instance"
	"github.com/zitadel/zitadel/internal/repository/org"
	user_repo "github.com/zitadel/zitadel/internal/repository/user"
	"github.com/zitadel/zitadel/internal/zerrors"
)

const (
	userTable = "auth.users3"
)

type User struct {
	view    *auth_view.View
	queries *query2.Queries
	es      handler.EventStore
}

var _ handler.Projection = (*User)(nil)

func newUser(
	ctx context.Context,
	config handler.Config,
	view *auth_view.View,
	queries *query2.Queries,
) *handler.Handler {
	return handler.NewHandler(
		ctx,
		&config,
		&User{
			view:    view,
			queries: queries,
			es:      config.Eventstore,
		},
	)
}

func (*User) Name() string {
	return userTable
}
func (u *User) Reducers() []handler.AggregateReducer {
	return []handler.AggregateReducer{
		{
			Aggregate: user_repo.AggregateType,
			EventReducers: []handler.EventReducer{
				{
					Event:  user_repo.HumanOTPSMSRemovedType,
					Reduce: u.ProcessUser,
				},
				{
					Event:  user_repo.HumanOTPEmailRemovedType,
					Reduce: u.ProcessUser,
				},
				{
					Event:  user_repo.HumanAddedType,
					Reduce: u.ProcessUser,
				},
				{
					Event:  user_repo.UserV1AddedType,
					Reduce: u.ProcessUser,
				},
				{
					Event:  user_repo.UserV1RegisteredType,
					Reduce: u.ProcessUser,
				},
				{
					Event:  user_repo.HumanRegisteredType,
					Reduce: u.ProcessUser,
				},
				{
					Event:  user_repo.UserV1PhoneRemovedType,
					Reduce: u.ProcessUser,
				},
				{
					Event:  user_repo.UserV1MFAOTPVerifiedType,
					Reduce: u.ProcessUser,
				},
				{
					Event:  user_repo.UserV1MFAInitSkippedType,
					Reduce: u.ProcessUser,
				},
				{
					Event:  user_repo.UserV1PasswordChangedType,
					Reduce: u.ProcessUser,
				},
				{
					Event:  user_repo.HumanPhoneRemovedType,
					Reduce: u.ProcessUser,
				},
				{
					Event:  user_repo.HumanMFAOTPVerifiedType,
					Reduce: u.ProcessUser,
				},
				{
					Event:  user_repo.HumanU2FTokenVerifiedType,
					Reduce: u.ProcessUser,
				},
				{
					Event:  user_repo.HumanMFAInitSkippedType,
					Reduce: u.ProcessUser,
				},
				{
					Event:  user_repo.HumanPasswordChangedType,
					Reduce: u.ProcessUser,
				},
				{
					Event:  user_repo.HumanInitialCodeAddedType,
					Reduce: u.ProcessUser,
				},
				{
					Event:  user_repo.UserV1InitialCodeAddedType,
					Reduce: u.ProcessUser,
				},
				{
					Event:  user_repo.UserV1InitializedCheckSucceededType,
					Reduce: u.ProcessUser,
				},
				{
					Event:  user_repo.HumanInitializedCheckSucceededType,
					Reduce: u.ProcessUser,
				},
				{
					Event:  user_repo.HumanPasswordlessInitCodeAddedType,
					Reduce: u.ProcessUser,
				},
				{
					Event:  user_repo.HumanPasswordlessInitCodeRequestedType,
					Reduce: u.ProcessUser,
				},
				{
					Event:  user_repo.UserRemovedType,
					Reduce: u.ProcessUser,
				},
			},
		},
		{
			Aggregate: org.AggregateType,
			EventReducers: []handler.EventReducer{
				{
					Event:  org.OrgRemovedEventType,
					Reduce: u.ProcessOrg,
				},
			},
		},
		{
			Aggregate: instance.AggregateType,
			EventReducers: []handler.EventReducer{
				{
					Event:  instance.InstanceRemovedEventType,
					Reduce: u.ProcessInstance,
				},
			},
		},
	}
}

func (u *User) setPasswordData(event eventstore.Event, secret *crypto.CryptoValue, hash string, changeRequired bool) *handler.Statement {
	set := secret != nil || hash != ""
	columns := []handler.Column{
		handler.NewCol(instanceIDCol, event.Aggregate().InstanceID),
		handler.NewCol("id", event.Aggregate().ID),
		handler.NewCol("password_set", set),
		handler.NewCol("password_init_required", !set),
		handler.NewCol("password_change", event.CreatedAt()),
		//handler.NewCol("password_change_required", changeRequired),
	}
	return handler.NewUpsertStatement(event, columns[0:2], columns)
}

//nolint:gocognit
func (u *User) ProcessUser(event eventstore.Event) (_ *handler.Statement, err error) {
	switch event.Type() {
	case user_repo.UserV1AddedType,
		user_repo.HumanAddedType:
		e, ok := event.(*user_repo.HumanAddedEvent)
		if !ok {
			return nil, zerrors.ThrowInvalidArgumentf(nil, "MODEL-SDAGF", "reduce.wrong.event.type %s", user_repo.HumanAddedType)
		}
		return u.setPasswordData(event, e.Secret, e.EncodedHash, e.ChangeRequired), nil
	case user_repo.UserV1RegisteredType,
		user_repo.HumanRegisteredType:
		e, ok := event.(*user_repo.HumanRegisteredEvent)
		if !ok {
			return nil, zerrors.ThrowInvalidArgumentf(nil, "MODEL-AS1hz", "reduce.wrong.event.type %s", user_repo.HumanRegisteredType)
		}
		return u.setPasswordData(event, e.Secret, e.EncodedHash, e.ChangeRequired), nil
	case user_repo.UserV1PasswordChangedType,
		user_repo.HumanPasswordChangedType:
		e, ok := event.(*user_repo.HumanPasswordChangedEvent)
		if !ok {
			return nil, zerrors.ThrowInvalidArgumentf(nil, "MODEL-Gd31w", "reduce.wrong.event.type %s", user_repo.HumanPasswordChangedType)
		}
		return u.setPasswordData(event, e.Secret, e.EncodedHash, e.ChangeRequired), nil
	case user_repo.UserV1PhoneRemovedType,
		user_repo.HumanPhoneRemovedType,
		user_repo.UserV1MFAOTPVerifiedType,
		user_repo.HumanMFAOTPVerifiedType,
		user_repo.HumanOTPSMSRemovedType,
		user_repo.HumanOTPEmailRemovedType,
		user_repo.HumanU2FTokenVerifiedType:
		return handler.NewUpdateStatement(event,
			[]handler.Column{
				handler.NewCol("mfa_init_skipped", time.Time{}),
			},
			[]handler.Condition{
				handler.NewCond(instanceIDCol, event.Aggregate().InstanceID),
				handler.NewCond("id", event.Aggregate().ID),
			}), nil
	case user_repo.UserV1MFAInitSkippedType,
		user_repo.HumanMFAInitSkippedType:
		return handler.NewUpdateStatement(event,
			[]handler.Column{
				handler.NewCol("mfa_init_skipped", event.CreatedAt()),
			},
			[]handler.Condition{
				handler.NewCond(instanceIDCol, event.Aggregate().InstanceID),
				handler.NewCond("id", event.Aggregate().ID),
			}), nil
	case user_repo.UserV1InitialCodeAddedType,
		user_repo.HumanInitialCodeAddedType:
		return handler.NewUpdateStatement(event,
			[]handler.Column{
				handler.NewCol("init_required", true),
			},
			[]handler.Condition{
				handler.NewCond(instanceIDCol, event.Aggregate().InstanceID),
				handler.NewCond("id", event.Aggregate().ID),
			}), nil
	case user_repo.UserV1InitializedCheckSucceededType,
		user_repo.HumanInitializedCheckSucceededType:
		return handler.NewUpdateStatement(event,
			[]handler.Column{
				handler.NewCol("init_required", false),
			},
			[]handler.Condition{
				handler.NewCond(instanceIDCol, event.Aggregate().InstanceID),
				handler.NewCond("id", event.Aggregate().ID),
			}), nil
	case user_repo.HumanPasswordlessInitCodeAddedType,
		user_repo.HumanPasswordlessInitCodeRequestedType:
		return handler.NewUpdateStatement(event,
			[]handler.Column{
				handler.NewCol("passwordless_init_required", true),
				handler.NewCol("password_init_required", false),
			},
			[]handler.Condition{
				handler.NewCond(instanceIDCol, event.Aggregate().InstanceID),
				handler.NewCond("id", event.Aggregate().ID),
				handler.NewCond("password_set", false),
			}), nil
	case user_repo.UserRemovedType:
		return handler.NewDeleteStatement(event,
			[]handler.Condition{
				handler.NewCond(instanceIDCol, event.Aggregate().InstanceID),
				handler.NewCond("id", event.Aggregate().ID),
			}), nil
	default:
		return handler.NewNoOpStatement(event), nil
	}
}

func (u *User) ProcessOrg(event eventstore.Event) (_ *handler.Statement, err error) {
	switch event.Type() {
	case org.OrgRemovedEventType:
		return handler.NewDeleteStatement(event,
			[]handler.Condition{
				handler.NewCond(instanceIDCol, event.Aggregate().InstanceID),
				handler.NewCond(resourceOwnerCol, event.Aggregate().ID),
			},
		), nil
	default:
		return handler.NewNoOpStatement(event), nil
	}
}

func (u *User) ProcessInstance(event eventstore.Event) (_ *handler.Statement, err error) {
	switch event.Type() {
	case instance.InstanceRemovedEventType:
		return handler.NewDeleteStatement(event,
			[]handler.Condition{
				handler.NewCond(instanceIDCol, event.Aggregate().InstanceID),
			},
		), nil
	default:
		return handler.NewNoOpStatement(event), nil
	}
}
