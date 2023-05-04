package command

import (
	"bytes"
	"context"
	"time"

	"github.com/zitadel/zitadel/internal/crypto"
	"github.com/zitadel/zitadel/internal/domain"
	caos_errs "github.com/zitadel/zitadel/internal/errors"
	"github.com/zitadel/zitadel/internal/eventstore"
	"github.com/zitadel/zitadel/internal/repository/session"
)

type SessionWriteModel struct {
	eventstore.WriteModel

	Token             *crypto.CryptoValue
	UserID            string
	UserCheckedAt     time.Time
	PasswordCheckedAt time.Time
	Metadata          map[string][]byte
	State             domain.SessionState

	commands  []eventstore.Command
	aggregate *eventstore.Aggregate
}

func NewSessionWriteModel(sessionID string, resourceOwner string) *SessionWriteModel {
	// var resourceOwner string //TODO: resourceowner?
	return &SessionWriteModel{
		WriteModel: eventstore.WriteModel{
			AggregateID:   sessionID,
			ResourceOwner: resourceOwner,
		},
		Metadata:  make(map[string][]byte),
		aggregate: &session.NewAggregate(sessionID, resourceOwner).Aggregate,
	}
}

func (wm *SessionWriteModel) Reduce() error {
	for _, event := range wm.Events {
		switch e := event.(type) {
		case *session.AddedEvent:
			wm.reduceAdded(e)
		case *session.UserCheckedEvent:
			wm.reduceUserChecked(e)
		case *session.PasswordCheckedEvent:
			wm.reducePasswordChecked(e)
		case *session.TokenSetEvent:
			wm.reduceTokenSet(e)
		case *session.TerminateEvent:
			wm.reduceTerminate()
		}
	}
	return wm.WriteModel.Reduce()
}

func (wm *SessionWriteModel) Query() *eventstore.SearchQueryBuilder {
	query := eventstore.NewSearchQueryBuilder(eventstore.ColumnsEvent).
		AddQuery().
		AggregateTypes(session.AggregateType).
		AggregateIDs(wm.AggregateID).
		EventTypes(
			session.AddedType,
			session.UserCheckedType,
			session.PasswordCheckedType,
			session.TokenSetType,
			session.MetadataSetType,
			session.TerminateType,
		).
		Builder()

	if wm.ResourceOwner != "" {
		query.ResourceOwner(wm.ResourceOwner)
	}
	return query
}

func (wm *SessionWriteModel) reduceAdded(e *session.AddedEvent) {
	wm.State = domain.SessionStateActive
}

func (wm *SessionWriteModel) reduceUserChecked(e *session.UserCheckedEvent) {
	wm.UserID = e.UserID
	wm.UserCheckedAt = e.CheckedAt
}

func (wm *SessionWriteModel) reducePasswordChecked(e *session.PasswordCheckedEvent) {
	wm.PasswordCheckedAt = e.CheckedAt
}

func (wm *SessionWriteModel) reduceTokenSet(e *session.TokenSetEvent) {
	wm.State = domain.SessionStateActive //TODO: ?
	wm.Token = e.Token
}

func (wm *SessionWriteModel) reduceTerminate() {
	wm.State = domain.SessionStateTerminated
}

func (wm *SessionWriteModel) Start(ctx context.Context) {
	wm.commands = append(wm.commands, session.NewAddedEvent(ctx, wm.aggregate))
}

func (wm *SessionWriteModel) UserChecked(ctx context.Context, userID string, checkedAt time.Time) error {
	if wm.UserID != "" && userID != "" && wm.UserID != userID {
		return caos_errs.ThrowInvalidArgument(nil, "", "user change not possible")
	}
	wm.commands = append(wm.commands, session.NewUserCheckedEvent(ctx, wm.aggregate, userID, checkedAt))
	// set the userID so other checks can use it
	wm.UserID = userID
	return nil
}

func (wm *SessionWriteModel) PasswordChecked(ctx context.Context, checkedAt time.Time) {
	wm.commands = append(wm.commands, session.NewPasswordCheckedEvent(ctx, wm.aggregate, checkedAt))
}

func (wm *SessionWriteModel) SetToken(ctx context.Context, token *crypto.CryptoValue) {
	wm.commands = append(wm.commands, session.NewTokenSetEvent(ctx, wm.aggregate, token))
}

func (wm *SessionWriteModel) ChangeMetadata(ctx context.Context, metadata map[string][]byte) {
	var changed bool
	for key, value := range metadata {
		currentValue, exists := wm.Metadata[key]

		if len(value) != 0 {
			// if a value is provided, and it's not equal, change it
			if !bytes.Equal(currentValue, value) {
				wm.Metadata[key] = value
				changed = true
			}
		} else {
			// if there's no / an empty value, we only need to remove it on existing entries
			if exists {
				delete(wm.Metadata, key)
				changed = true
			}
		}
	}
	if changed {
		wm.commands = append(wm.commands, session.NewMetadataSetEvent(ctx, wm.aggregate, wm.Metadata))
	}
}
