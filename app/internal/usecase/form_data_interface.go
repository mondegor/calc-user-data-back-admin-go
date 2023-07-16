package usecase

import (
    "context"
    "print-shop-back/internal/entity"
    "print-shop-back/pkg/mrentity"
)

type (
    FormDataService interface {
        GetList(ctx context.Context, listFilter *entity.FormDataListFilter) ([]entity.FormData, error)
        GetItem(ctx context.Context, id mrentity.KeyInt32) (*entity.FormData, error)
        CheckAvailability(ctx context.Context, id mrentity.KeyInt32) error
        Create(ctx context.Context, item *entity.FormData) error
        Store(ctx context.Context, item *entity.FormData) error
        ChangeStatus(ctx context.Context, item *entity.FormData) error
        Remove(ctx context.Context, id mrentity.KeyInt32) error
    }

    UIFormDataService interface {
        CompileForm(ctx context.Context, id mrentity.KeyInt32) (*entity.UIForm, error)
    }

    FormDataStorage interface {
        LoadAll(ctx context.Context, listFilter *entity.FormDataListFilter, rows *[]entity.FormData) error
        LoadOne(ctx context.Context, row *entity.FormData) error
        FetchIdByName(ctx context.Context, row *entity.FormData) (mrentity.KeyInt32, error)
        FetchStatus(ctx context.Context, row *entity.FormData) (entity.ItemStatus, error)
        IsExists(ctx context.Context, id mrentity.KeyInt32) error
        Insert(ctx context.Context, row *entity.FormData) error
        Update(ctx context.Context, row *entity.FormData) error
        UpdateStatus(ctx context.Context, row *entity.FormData) error
        Delete(ctx context.Context, id mrentity.KeyInt32) error
    }
)
