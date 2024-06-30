package papercolor

import (
	"github.com/mondegor/print-shop-back/pkg/validate"

	"github.com/mondegor/go-storage/mrstorage"
	"github.com/mondegor/go-sysmess/mrlang"
	"github.com/mondegor/go-webcore/mrcore"
	"github.com/mondegor/go-webcore/mrsender"
	"github.com/mondegor/go-webcore/mrserver"
)

type (
	// Options - comment struct.
	Options struct {
		EventEmitter   mrsender.EventEmitter
		UsecaseHelper  mrcore.UsecaseErrorWrapper
		DBConnManager  mrstorage.DBConnManager
		RequestParsers RequestParsers
		ResponseSender mrserver.ResponseSender

		UnitPaperColor UnitPaperColorOptions

		PageSizeMax     uint64
		PageSizeDefault uint64
	}

	// UnitPaperColorOptions - comment struct.
	UnitPaperColorOptions struct {
		Dictionary *mrlang.MultiLangDictionary
	}

	// RequestParsers - comment struct.
	RequestParsers struct {
		Parser       *validate.Parser
		ExtendParser *validate.ExtendParser
	}
)
