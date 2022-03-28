package command

import (
	"context"

	"github.com/caos/zitadel/internal/eventstore"

	"github.com/caos/zitadel/internal/repository/instance"
	"github.com/caos/zitadel/internal/repository/policy"
)

type InstancePasswordAgePolicyWriteModel struct {
	PasswordAgePolicyWriteModel
}

func NewInstancePasswordAgePolicyWriteModel(instanceID string) *InstancePasswordAgePolicyWriteModel {
	return &InstancePasswordAgePolicyWriteModel{
		PasswordAgePolicyWriteModel{
			WriteModel: eventstore.WriteModel{
				AggregateID:   instanceID,
				ResourceOwner: instanceID,
			},
		},
	}
}

func (wm *InstancePasswordAgePolicyWriteModel) AppendEvents(events ...eventstore.Event) {
	for _, event := range events {
		switch e := event.(type) {
		case *instance.PasswordAgePolicyAddedEvent:
			wm.PasswordAgePolicyWriteModel.AppendEvents(&e.PasswordAgePolicyAddedEvent)
		case *instance.PasswordAgePolicyChangedEvent:
			wm.PasswordAgePolicyWriteModel.AppendEvents(&e.PasswordAgePolicyChangedEvent)
		}
	}
}

func (wm *InstancePasswordAgePolicyWriteModel) Reduce() error {
	return wm.PasswordAgePolicyWriteModel.Reduce()
}

func (wm *InstancePasswordAgePolicyWriteModel) Query() *eventstore.SearchQueryBuilder {
	return eventstore.NewSearchQueryBuilder(eventstore.ColumnsEvent).
		ResourceOwner(wm.ResourceOwner).
		AddQuery().
		AggregateTypes(instance.AggregateType).
		AggregateIDs(wm.PasswordAgePolicyWriteModel.AggregateID).
		EventTypes(
			instance.PasswordAgePolicyAddedEventType,
			instance.PasswordAgePolicyChangedEventType).
		Builder()
}

func (wm *InstancePasswordAgePolicyWriteModel) NewChangedEvent(
	ctx context.Context,
	aggregate *eventstore.Aggregate,
	expireWarnDays,
	maxAgeDays uint64) (*instance.PasswordAgePolicyChangedEvent, bool) {
	changes := make([]policy.PasswordAgePolicyChanges, 0)
	if wm.ExpireWarnDays != expireWarnDays {
		changes = append(changes, policy.ChangeExpireWarnDays(expireWarnDays))
	}
	if wm.MaxAgeDays != maxAgeDays {
		changes = append(changes, policy.ChangeMaxAgeDays(maxAgeDays))
	}
	if len(changes) == 0 {
		return nil, false
	}
	changedEvent, err := instance.NewPasswordAgePolicyChangedEvent(ctx, aggregate, changes)
	if err != nil {
		return nil, false
	}
	return changedEvent, true
}
