package usecase

import (
	"context"

	"github.com/mondegor/print-shop-back/internal/dictionaries/papercolor/section/adm"
	"github.com/mondegor/print-shop-back/internal/dictionaries/papercolor/section/adm/entity"

	"github.com/mondegor/go-sysmess/mrmsg"
	"github.com/mondegor/go-webcore/mrcore"
	"github.com/mondegor/go-webcore/mrenum"
	"github.com/mondegor/go-webcore/mrsender"
	"github.com/mondegor/go-webcore/mrstatus"
	"github.com/mondegor/go-webcore/mrstatus/mrflow"

	"github.com/mondegor/go-webcore/mrtype"
)

type (
	// PaperColor - comment struct.
	PaperColor struct {
		storage      adm.PaperColorStorage
		eventEmitter mrsender.EventEmitter
		errorWrapper mrcore.UsecaseErrorWrapper
		statusFlow   mrstatus.Flow
	}
)

// NewPaperColor - создаёт объект PaperColor.
func NewPaperColor(storage adm.PaperColorStorage, eventEmitter mrsender.EventEmitter, errorWrapper mrcore.UsecaseErrorWrapper) *PaperColor {
	return &PaperColor{
		storage:      storage,
		eventEmitter: eventEmitter,
		errorWrapper: errorWrapper,
		statusFlow:   mrflow.ItemStatusFlow(),
	}
}

// GetList - comment method.
func (uc *PaperColor) GetList(ctx context.Context, params entity.PaperColorParams) ([]entity.PaperColor, int64, error) {
	fetchParams := uc.storage.NewSelectParams(params)

	total, err := uc.storage.FetchTotal(ctx, fetchParams.Where)
	if err != nil {
		return nil, 0, uc.errorWrapper.WrapErrorFailed(err, entity.ModelNamePaperColor)
	}

	if total < 1 {
		return nil, 0, nil
	}

	items, err := uc.storage.Fetch(ctx, fetchParams)
	if err != nil {
		return nil, 0, uc.errorWrapper.WrapErrorFailed(err, entity.ModelNamePaperColor)
	}

	return items, total, nil
}

// GetItem - comment method.
func (uc *PaperColor) GetItem(ctx context.Context, itemID mrtype.KeyInt32) (entity.PaperColor, error) {
	if itemID < 1 {
		return entity.PaperColor{}, mrcore.ErrUseCaseEntityNotFound.New()
	}

	item, err := uc.storage.FetchOne(ctx, itemID)
	if err != nil {
		return entity.PaperColor{}, uc.errorWrapper.WrapErrorEntityNotFoundOrFailed(err, entity.ModelNamePaperColor, itemID)
	}

	return item, nil
}

// Create - comment method.
func (uc *PaperColor) Create(ctx context.Context, item entity.PaperColor) (mrtype.KeyInt32, error) {
	item.Status = mrenum.ItemStatusDraft

	itemID, err := uc.storage.Insert(ctx, item)
	if err != nil {
		return 0, uc.errorWrapper.WrapErrorFailed(err, entity.ModelNamePaperColor)
	}

	uc.emitEvent(ctx, "Create", mrmsg.Data{"id": itemID})

	return itemID, nil
}

// Store - comment method.
func (uc *PaperColor) Store(ctx context.Context, item entity.PaperColor) error {
	if item.ID < 1 {
		return mrcore.ErrUseCaseEntityNotFound.New()
	}

	if item.TagVersion < 1 {
		return mrcore.ErrUseCaseEntityVersionInvalid.New()
	}

	// предварительная проверка существования записи нужна для того,
	// чтобы при Update быть уверенным, что отсутствие записи из-за ошибки VersionInvalid
	if _, err := uc.storage.FetchStatus(ctx, item.ID); err != nil {
		return uc.errorWrapper.WrapErrorEntityNotFoundOrFailed(err, entity.ModelNamePaperColor, item.ID)
	}

	tagVersion, err := uc.storage.Update(ctx, item)
	if err != nil {
		if uc.errorWrapper.IsNotFoundError(err) {
			return mrcore.ErrUseCaseEntityVersionInvalid.Wrap(err)
		}

		return uc.errorWrapper.WrapErrorFailed(err, entity.ModelNamePaperColor)
	}

	uc.emitEvent(ctx, "Store", mrmsg.Data{"id": item.ID, "ver": tagVersion})

	return nil
}

// ChangeStatus - comment method.
func (uc *PaperColor) ChangeStatus(ctx context.Context, item entity.PaperColor) error {
	if item.ID < 1 {
		return mrcore.ErrUseCaseEntityNotFound.New()
	}

	if item.TagVersion < 1 {
		return mrcore.ErrUseCaseEntityVersionInvalid.New()
	}

	currentStatus, err := uc.storage.FetchStatus(ctx, item.ID)
	if err != nil {
		return uc.errorWrapper.WrapErrorEntityNotFoundOrFailed(err, entity.ModelNamePaperColor, item.ID)
	}

	if currentStatus == item.Status {
		return nil
	}

	if !uc.statusFlow.Check(currentStatus, item.Status) {
		return mrcore.ErrUseCaseSwitchStatusRejected.New(currentStatus, item.Status)
	}

	tagVersion, err := uc.storage.UpdateStatus(ctx, item)
	if err != nil {
		if uc.errorWrapper.IsNotFoundError(err) {
			return mrcore.ErrUseCaseEntityVersionInvalid.Wrap(err)
		}

		return uc.errorWrapper.WrapErrorFailed(err, entity.ModelNamePaperColor)
	}

	uc.emitEvent(ctx, "ChangeStatus", mrmsg.Data{"id": item.ID, "ver": tagVersion, "status": item.Status})

	return nil
}

// Remove - comment method.
func (uc *PaperColor) Remove(ctx context.Context, itemID mrtype.KeyInt32) error {
	if itemID < 1 {
		return mrcore.ErrUseCaseEntityNotFound.New()
	}

	if err := uc.storage.Delete(ctx, itemID); err != nil {
		return uc.errorWrapper.WrapErrorEntityNotFoundOrFailed(err, entity.ModelNamePaperColor, itemID)
	}

	uc.emitEvent(ctx, "Remove", mrmsg.Data{"id": itemID})

	return nil
}

func (uc *PaperColor) emitEvent(ctx context.Context, eventName string, data mrmsg.Data) {
	uc.eventEmitter.EmitWithSource(
		ctx,
		eventName,
		entity.ModelNamePaperColor,
		data,
	)
}
