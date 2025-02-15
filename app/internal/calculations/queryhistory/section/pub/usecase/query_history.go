package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/mondegor/go-sysmess/mrmsg"
	"github.com/mondegor/go-webcore/mrcore"
	"github.com/mondegor/go-webcore/mrlog"
	"github.com/mondegor/go-webcore/mrsender"
	"github.com/mondegor/go-webcore/mrsender/decorator"

	"github.com/mondegor/print-shop-back/internal/calculations/queryhistory/section/pub"
	"github.com/mondegor/print-shop-back/internal/calculations/queryhistory/section/pub/entity"
)

type (
	// QueryHistory - comment struct.
	QueryHistory struct {
		storage      pub.QueryResultStorage
		eventEmitter mrsender.EventEmitter
		errorWrapper mrcore.UseCaseErrorWrapper
	}
)

// NewQueryHistory - создаёт объект QueryHistory.
func NewQueryHistory(storage pub.QueryResultStorage, eventEmitter mrsender.EventEmitter, errorWrapper mrcore.UseCaseErrorWrapper) *QueryHistory {
	return &QueryHistory{
		storage:      storage,
		eventEmitter: decorator.NewSourceEmitter(eventEmitter, entity.ModelNameQueryHistory),
		errorWrapper: errorWrapper,
	}
}

// GetItem - comment method.
func (uc *QueryHistory) GetItem(ctx context.Context, itemID uuid.UUID) (entity.QueryHistoryItem, error) {
	if itemID == uuid.Nil {
		return entity.QueryHistoryItem{}, mrcore.ErrUseCaseEntityNotFound.New()
	}

	item, err := uc.storage.FetchOne(ctx, itemID)
	if err != nil {
		return entity.QueryHistoryItem{}, uc.errorWrapper.WrapErrorEntityNotFoundOrFailed(err, entity.ModelNameQueryHistory, itemID)
	}

	// обновление счётчика посещений
	// TODO: send to queue
	go func() {
		if err := uc.storage.UpdateQuantity(ctx, itemID); err != nil {
			mrlog.Ctx(ctx).Error().Err(err).Send()
		}
	}()

	return item, nil
}

// Create - comment method.
func (uc *QueryHistory) Create(ctx context.Context, item entity.QueryHistoryItem) (itemID uuid.UUID, err error) {
	itemID, err = uc.storage.Insert(ctx, item)
	if err != nil {
		return uuid.Nil, uc.errorWrapper.WrapErrorFailed(err, entity.ModelNameQueryHistory)
	}

	uc.eventEmitter.Emit(ctx, "Create", mrmsg.Data{"id": itemID})

	return itemID, nil
}
