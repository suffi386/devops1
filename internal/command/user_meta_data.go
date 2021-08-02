package command

import (
	"context"

	"github.com/caos/zitadel/internal/domain"
	caos_errs "github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore"
	"github.com/caos/zitadel/internal/repository/user"
)

func (c *Commands) SetUserMetadata(ctx context.Context, metaData *domain.Metadata, userID, resourceOwner string) (_ *domain.Metadata, err error) {
	err = c.checkUserExists(ctx, userID, resourceOwner)
	if err != nil {
		return nil, err
	}
	setMetadata := NewUserMetadataWriteModel(userID, resourceOwner, metaData.Key)
	userAgg := UserAggregateFromWriteModel(&setMetadata.WriteModel)
	event, err := c.setUserMetadata(ctx, userAgg, metaData)
	if err != nil {
		return nil, err
	}
	pushedEvents, err := c.eventstore.PushEvents(ctx, event)
	if err != nil {
		return nil, err
	}

	err = AppendAndReduce(setMetadata, pushedEvents...)
	if err != nil {
		return nil, err
	}
	return writeModelToUserMetadata(setMetadata), nil
}

func (c *Commands) BulkSetUserMetadata(ctx context.Context, userID, resourceOwner string, metaDatas ...*domain.Metadata) (_ *domain.ObjectDetails, err error) {
	if len(metaDatas) == 0 {
		return nil, caos_errs.ThrowPreconditionFailed(nil, "META-9mm2d", "Errors.Metadata.NoData")
	}
	err = c.checkUserExists(ctx, userID, resourceOwner)
	if err != nil {
		return nil, err
	}

	events := make([]eventstore.EventPusher, len(metaDatas))
	setMetadata := NewUserMetadataListWriteModel(userID, resourceOwner)
	userAgg := UserAggregateFromWriteModel(&setMetadata.WriteModel)
	for i, data := range metaDatas {
		event, err := c.setUserMetadata(ctx, userAgg, data)
		if err != nil {
			return nil, err
		}
		events[i] = event
	}

	pushedEvents, err := c.eventstore.PushEvents(ctx, events...)
	if err != nil {
		return nil, err
	}

	err = AppendAndReduce(setMetadata, pushedEvents...)
	if err != nil {
		return nil, err
	}
	return writeModelToObjectDetails(&setMetadata.WriteModel), nil
}

func (c *Commands) setUserMetadata(ctx context.Context, userAgg *eventstore.Aggregate, metaData *domain.Metadata) (pusher eventstore.EventPusher, err error) {
	if !metaData.IsValid() {
		return nil, caos_errs.ThrowInvalidArgument(nil, "META-2m00f", "Errors.Metadata.Invalid")
	}
	pusher = user.NewMetadataSetEvent(
		ctx,
		userAgg,
		metaData.Key,
		metaData.Value,
	)
	return pusher, nil
}

func (c *Commands) RemoveUserMetadata(ctx context.Context, metaDataKey, userID, resourceOwner string) (_ *domain.ObjectDetails, err error) {
	if metaDataKey == "" {
		return nil, caos_errs.ThrowInvalidArgument(nil, "META-2n0fs", "Errors.Metadata.Invalid")
	}
	err = c.checkUserExists(ctx, userID, resourceOwner)
	if err != nil {
		return nil, err
	}
	removeMetadata, err := c.getUserMetadataModelByID(ctx, userID, resourceOwner, metaDataKey)
	if err != nil {
		return nil, err
	}
	if !removeMetadata.State.Exists() {
		return nil, caos_errs.ThrowNotFound(nil, "META-ncnw3", "Errors.Metadata.NotFound")
	}
	userAgg := UserAggregateFromWriteModel(&removeMetadata.WriteModel)
	event, err := c.removeUserMetadata(ctx, userAgg, metaDataKey)
	if err != nil {
		return nil, err
	}
	pushedEvents, err := c.eventstore.PushEvents(ctx, event)
	if err != nil {
		return nil, err
	}

	err = AppendAndReduce(removeMetadata, pushedEvents...)
	if err != nil {
		return nil, err
	}
	return writeModelToObjectDetails(&removeMetadata.WriteModel), nil
}

func (c *Commands) BulkRemoveUserMetadata(ctx context.Context, userID, resourceOwner string, metaDataKeys ...string) (_ *domain.ObjectDetails, err error) {
	if len(metaDataKeys) == 0 {
		return nil, caos_errs.ThrowPreconditionFailed(nil, "META-9mm2d", "Errors.Metadata.NoData")
	}
	err = c.checkUserExists(ctx, userID, resourceOwner)
	if err != nil {
		return nil, err
	}

	events := make([]eventstore.EventPusher, len(metaDataKeys))
	removeMetadata, err := c.getUserMetadataListModelByID(ctx, userID, resourceOwner)
	if err != nil {
		return nil, err
	}
	userAgg := UserAggregateFromWriteModel(&removeMetadata.WriteModel)
	for i, key := range metaDataKeys {
		if key == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "COMMAND-m29ds", "Errors.Metadata.Invalid")
		}
		if _, found := removeMetadata.metaDataList[key]; !found {
			return nil, caos_errs.ThrowNotFound(nil, "META-2nnds", "Errors.Metadata.KeyNotExisting")
		}
		event, err := c.removeUserMetadata(ctx, userAgg, key)
		if err != nil {
			return nil, err
		}
		events[i] = event
	}

	pushedEvents, err := c.eventstore.PushEvents(ctx, events...)
	if err != nil {
		return nil, err
	}

	err = AppendAndReduce(removeMetadata, pushedEvents...)
	if err != nil {
		return nil, err
	}
	return writeModelToObjectDetails(&removeMetadata.WriteModel), nil
}

func (c *Commands) removeUserMetadata(ctx context.Context, userAgg *eventstore.Aggregate, metaDataKey string) (pusher eventstore.EventPusher, err error) {
	pusher = user.NewMetadataRemovedEvent(
		ctx,
		userAgg,
		metaDataKey,
	)
	return pusher, nil
}

func (c *Commands) getUserMetadataModelByID(ctx context.Context, userID, resourceOwner, key string) (*UserMetadataWriteModel, error) {
	userMetadataWriteModel := NewUserMetadataWriteModel(userID, resourceOwner, key)
	err := c.eventstore.FilterToQueryReducer(ctx, userMetadataWriteModel)
	if err != nil {
		return nil, err
	}
	return userMetadataWriteModel, nil
}

func (c *Commands) getUserMetadataListModelByID(ctx context.Context, userID, resourceOwner string) (*UserMetadataListWriteModel, error) {
	userMetadataWriteModel := NewUserMetadataListWriteModel(userID, resourceOwner)
	err := c.eventstore.FilterToQueryReducer(ctx, userMetadataWriteModel)
	if err != nil {
		return nil, err
	}
	return userMetadataWriteModel, nil
}
