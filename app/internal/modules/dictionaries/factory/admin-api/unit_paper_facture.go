package factory

import (
	module "print-shop-back/internal/modules/dictionaries"
	http_v1 "print-shop-back/internal/modules/dictionaries/controller/http_v1/admin-api"
	entity "print-shop-back/internal/modules/dictionaries/entity/admin-api"
	"print-shop-back/internal/modules/dictionaries/factory"
	repository "print-shop-back/internal/modules/dictionaries/infrastructure/repository/admin-api"
	usecase "print-shop-back/internal/modules/dictionaries/usecase/admin-api"

	"github.com/mondegor/go-storage/mrpostgres"
	"github.com/mondegor/go-storage/mrsql"
	"github.com/mondegor/go-webcore/mrserver"
)

func createUnitPaperFacture(opts *factory.Options) ([]mrserver.HttpController, error) {
	var list []mrserver.HttpController

	if c, err := newUnitPaperFacture(opts); err != nil {
		return nil, err
	} else {
		list = append(list, c)
	}

	return list, nil
}

func newUnitPaperFacture(opts *factory.Options) (*http_v1.PaperFacture, error) {
	metaOrderBy, err := mrsql.NewEntityMetaOrderBy(entity.PaperFacture{})

	if err != nil {
		return nil, err
	}

	storage := repository.NewPaperFacturePostgres(
		opts.PostgresAdapter,
		mrsql.NewBuilderSelect(
			mrpostgres.NewSqlBuilderWhere(),
			mrpostgres.NewSqlBuilderOrderByWithDefaultSort(metaOrderBy.DefaultSort()),
			mrpostgres.NewSqlBuilderPager(module.PageSizeMax),
		),
	)
	service := usecase.NewPaperFacture(storage, opts.EventBox, opts.ServiceHelper)
	controller := http_v1.NewPaperFacture(
		opts.RequestParser,
		opts.ResponseSender,
		service,
		metaOrderBy,
	)

	return controller, nil
}
