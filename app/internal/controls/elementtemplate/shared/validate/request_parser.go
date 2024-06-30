package validate

import (
	"github.com/mondegor/print-shop-back/pkg/controls/validate"
	pkgvalidate "github.com/mondegor/print-shop-back/pkg/validate"

	"github.com/mondegor/go-webcore/mrserver"
	"github.com/mondegor/go-webcore/mrserver/mrparser"
)

type (
	// RequestElementTemplateParser - comment interface.
	RequestElementTemplateParser interface {
		pkgvalidate.RequestExtendParser
		mrserver.RequestParserFile
		validate.RequestDetailingParser
	}

	// Parser - comment struct.
	Parser struct {
		*pkgvalidate.ExtendParser
		*mrparser.File
		*validate.DetailingParser
	}
)

// NewParser - создаёт объект Parser.
func NewParser(
	p1 *pkgvalidate.ExtendParser,
	p2 *mrparser.File,
	p3 *validate.DetailingParser,
) *Parser {
	return &Parser{
		ExtendParser:    p1,
		File:            p2,
		DetailingParser: p3,
	}
}
