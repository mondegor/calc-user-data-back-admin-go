package controls

import (
	"context"

	"github.com/mondegor/print-shop-back/internal/app"
	"github.com/mondegor/print-shop-back/internal/controls/submitform/module"
	"github.com/mondegor/print-shop-back/internal/controls/submitform/shared/validate"
	"github.com/mondegor/print-shop-back/internal/factory/controls/elementtemplate/api/header"
	"github.com/mondegor/print-shop-back/internal/factory/controls/submitform"
	pkgvalidate "github.com/mondegor/print-shop-back/pkg/controls/validate"

	"github.com/mondegor/go-webcore/mrcore/mrinit"
)

// NewSubmitFormModuleOptions - создаёт объект submitform.Options.
func NewSubmitFormModuleOptions(_ context.Context, opts app.Options) (submitform.Options, error) {
	return submitform.Options{
		EventEmitter:  opts.EventEmitter,
		UsecaseHelper: opts.UsecaseErrorWrapper,
		DBConnManager: opts.PostgresConnManager,
		Locker:        opts.Locker,
		RequestParsers: submitform.RequestParsers{
			// Parser:       opts.RequestParsers.Parser,
			// ExtendParser: opts.RequestParsers.ExtendParser,
			ModuleParser: validate.NewParser(
				opts.RequestParsers.ExtendParser,
				opts.RequestParsers.FileJson,
				pkgvalidate.NewDetailingParser(),
			),
		},
		ResponseSender: opts.ResponseSenders.FileSender,

		ElementTemplateAPI: header.NewElementTemplate(opts.PostgresConnManager, opts.UsecaseErrorWrapper),
		OrdererAPI:         opts.OrdererAPI,

		PageSizeMax:     opts.Cfg.General.PageSizeMax,
		PageSizeDefault: opts.Cfg.General.PageSizeDefault,
	}, nil
}

// RegisterSubmitFormErrors - comment func.
func RegisterSubmitFormErrors(em *mrinit.ErrorManager) {
	em.RegisterList(mrinit.WrapProtoList(module.Errors()))
}
