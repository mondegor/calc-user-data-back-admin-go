package usecase

import (
	"context"

	"github.com/mondegor/go-sysmess/mrmsg"
	"github.com/mondegor/go-webcore/mrcore"
	"github.com/mondegor/go-webcore/mrenum"
	"github.com/mondegor/go-webcore/mrsender"
	"github.com/mondegor/go-webcore/mrsender/decorator"
	"github.com/mondegor/go-webcore/mrstatus"
	"github.com/mondegor/go-webcore/mrstatus/mrflow"

	"github.com/mondegor/print-shop-back/internal/catalog/laminate/module"
	"github.com/mondegor/print-shop-back/internal/catalog/laminate/section/adm"
	"github.com/mondegor/print-shop-back/internal/catalog/laminate/section/adm/entity"
	"github.com/mondegor/print-shop-back/pkg/dictionaries/api"
)

type (
	// Laminate - comment struct.
	Laminate struct {
		storage         adm.LaminateStorage
		materialTypeAPI api.MaterialTypeAvailability
		eventEmitter    mrsender.EventEmitter
		errorWrapper    mrcore.UseCaseErrorWrapper
		statusFlow      mrstatus.Flow
	}
)

// NewLaminate - создаёт объект NewLaminate.
func NewLaminate(
	storage adm.LaminateStorage,
	materialTypeAPI api.MaterialTypeAvailability,
	eventEmitter mrsender.EventEmitter,
	errorWrapper mrcore.UseCaseErrorWrapper,
) *Laminate {
	return &Laminate{
		storage:         storage,
		materialTypeAPI: materialTypeAPI,
		eventEmitter:    decorator.NewSourceEmitter(eventEmitter, entity.ModelNameLaminate),
		errorWrapper:    errorWrapper,
		statusFlow:      mrflow.ItemStatusFlow(),
	}
}

// GetList - comment method.
func (uc *Laminate) GetList(ctx context.Context, params entity.LaminateParams) (items []entity.Laminate, countItems uint64, err error) {
	items, countItems, err = uc.storage.FetchWithTotal(ctx, params)
	if err != nil {
		return nil, 0, uc.errorWrapper.WrapErrorFailed(err, entity.ModelNameLaminate)
	}

	if countItems == 0 {
		return make([]entity.Laminate, 0), 0, nil
	}

	return items, countItems, nil
}

// GetItem - comment method.
func (uc *Laminate) GetItem(ctx context.Context, itemID uint64) (entity.Laminate, error) {
	if itemID == 0 {
		return entity.Laminate{}, mrcore.ErrUseCaseEntityNotFound.New()
	}

	item, err := uc.storage.FetchOne(ctx, itemID)
	if err != nil {
		return entity.Laminate{}, uc.errorWrapper.WrapErrorEntityNotFoundOrFailed(err, entity.ModelNameLaminate, itemID)
	}

	return item, nil
}

// Create - comment method.
func (uc *Laminate) Create(ctx context.Context, item entity.Laminate) (itemID uint64, err error) {
	if err = uc.checkItem(ctx, &item); err != nil {
		return 0, err
	}

	item.Status = mrenum.ItemStatusDraft

	itemID, err = uc.storage.Insert(ctx, item)
	if err != nil {
		return 0, uc.errorWrapper.WrapErrorFailed(err, entity.ModelNameLaminate)
	}

	uc.eventEmitter.Emit(ctx, "Create", mrmsg.Data{"id": itemID})

	return itemID, nil
}

// Store - comment method.
func (uc *Laminate) Store(ctx context.Context, item entity.Laminate) error {
	if item.ID == 0 {
		return mrcore.ErrUseCaseEntityNotFound.New()
	}

	if item.TagVersion == 0 {
		return mrcore.ErrUseCaseEntityVersionInvalid.New()
	}

	// предварительная проверка существования записи нужна для того,
	// чтобы при Update быть уверенным, что отсутствие записи из-за ошибки VersionInvalid
	if _, err := uc.storage.FetchStatus(ctx, item.ID); err != nil {
		return uc.errorWrapper.WrapErrorEntityNotFoundOrFailed(err, entity.ModelNameLaminate, item.ID)
	}

	if err := uc.checkItem(ctx, &item); err != nil {
		return err
	}

	tagVersion, err := uc.storage.Update(ctx, item)
	if err != nil {
		if uc.errorWrapper.IsNotFoundError(err) {
			return mrcore.ErrUseCaseEntityVersionInvalid.Wrap(err)
		}

		return uc.errorWrapper.WrapErrorFailed(err, entity.ModelNameLaminate)
	}

	uc.eventEmitter.Emit(ctx, "Store", mrmsg.Data{"id": item.ID, "ver": tagVersion})

	return nil
}

// ChangeStatus - comment method.
func (uc *Laminate) ChangeStatus(ctx context.Context, item entity.Laminate) error {
	if item.ID == 0 {
		return mrcore.ErrUseCaseEntityNotFound.New()
	}

	if item.TagVersion == 0 {
		return mrcore.ErrUseCaseEntityVersionInvalid.New()
	}

	currentStatus, err := uc.storage.FetchStatus(ctx, item.ID)
	if err != nil {
		return uc.errorWrapper.WrapErrorEntityNotFoundOrFailed(err, entity.ModelNameLaminate, item.ID)
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

		return uc.errorWrapper.WrapErrorFailed(err, entity.ModelNameLaminate)
	}

	uc.eventEmitter.Emit(ctx, "ChangeStatus", mrmsg.Data{"id": item.ID, "ver": tagVersion, "status": item.Status})

	return nil
}

// Remove - comment method.
func (uc *Laminate) Remove(ctx context.Context, itemID uint64) error {
	if itemID == 0 {
		return mrcore.ErrUseCaseEntityNotFound.New()
	}

	if err := uc.storage.Delete(ctx, itemID); err != nil {
		return uc.errorWrapper.WrapErrorEntityNotFoundOrFailed(err, entity.ModelNameLaminate, itemID)
	}

	uc.eventEmitter.Emit(ctx, "Remove", mrmsg.Data{"id": itemID})

	return nil
}

func (uc *Laminate) checkItem(ctx context.Context, item *entity.Laminate) error {
	if err := uc.checkArticle(ctx, item); err != nil {
		return err
	}

	if item.ID == 0 || item.TypeID > 0 {
		if err := uc.materialTypeAPI.CheckingAvailability(ctx, item.TypeID); err != nil {
			return err
		}
	}

	return nil
}

func (uc *Laminate) checkArticle(ctx context.Context, item *entity.Laminate) error {
	id, err := uc.storage.FetchIDByArticle(ctx, item.Article)
	if err != nil {
		if uc.errorWrapper.IsNotFoundError(err) {
			return nil
		}

		return uc.errorWrapper.WrapErrorFailed(err, entity.ModelNameLaminate)
	}

	if item.ID != id {
		return module.ErrLaminateArticleAlreadyExists.New(item.Article)
	}

	return nil
}
