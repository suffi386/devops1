package user

import (
	"github.com/zitadel/zitadel/internal/v2/eventstore"
	"github.com/zitadel/zitadel/internal/zerrors"
)

type ReactivatedEvent eventstore.Event[eventstore.EmptyPayload]

const ReactivatedType = AggregateType + ".reactivated"

var _ eventstore.TypeChecker = (*ReactivatedEvent)(nil)

// ActionType implements eventstore.Typer.
func (c *ReactivatedEvent) ActionType() string {
	return ReactivatedType
}

func ReactivatedEventFromStorage(event *eventstore.StorageEvent) (e *ReactivatedEvent, _ error) {
	if event.Type != e.ActionType() {
		return nil, zerrors.ThrowInvalidArgument(nil, "ORG-jeeON", "Errors.Invalid.Event.Type")
	}

	return &ReactivatedEvent{
		StorageEvent: event,
	}, nil
}
