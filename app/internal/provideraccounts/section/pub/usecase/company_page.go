package usecase

import (
	"context"

	"github.com/mondegor/print-shop-back/internal/provideraccounts/section/pub"
	"github.com/mondegor/print-shop-back/internal/provideraccounts/section/pub/entity"

	"github.com/mondegor/go-webcore/mrcore"
)

type (
	// CompanyPage - comment struct.
	CompanyPage struct {
		storage      pub.CompanyPageStorage
		errorWrapper mrcore.UsecaseErrorWrapper
	}
)

// NewCompanyPage - создаёт объект CompanyPage.
func NewCompanyPage(storage pub.CompanyPageStorage, errorWrapper mrcore.UsecaseErrorWrapper) *CompanyPage {
	return &CompanyPage{
		storage:      storage,
		errorWrapper: errorWrapper,
	}
}

// GetItemByRewriteName - comment method.
func (uc *CompanyPage) GetItemByRewriteName(ctx context.Context, rewriteName string) (entity.CompanyPage, error) {
	if rewriteName == "" {
		return entity.CompanyPage{}, mrcore.ErrUseCaseEntityNotFound.New()
	}

	item, err := uc.storage.FetchByRewriteName(ctx, rewriteName)
	if err != nil {
		return entity.CompanyPage{}, uc.errorWrapper.WrapErrorEntityNotFoundOrFailed(err, entity.ModelNameCompanyPage, rewriteName)
	}

	return item, nil
}
