package papercolor

import (
	"github.com/mondegor/go-storage/mrstorage"
	"github.com/mondegor/go-sysmess/mrlang"
	"github.com/mondegor/go-webcore/mrcore"
	"github.com/mondegor/go-webcore/mrsender"
	"github.com/mondegor/go-webcore/mrserver"

	"github.com/mondegor/print-shop-back/pkg/validate"
)

type (
	// Options - comment struct.
	Options struct {
		EventEmitter        mrsender.EventEmitter
		UseCaseErrorWrapper mrcore.UseCaseErrorWrapper
		DBConnManager       mrstorage.DBConnManager
		RequestParsers      RequestParsers
		ResponseSender      mrserver.ResponseSender

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
