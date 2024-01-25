package dictionaries

import (
	"context"

	"github.com/mondegor/go-webcore/mrtype"
)

type (
	PrintFormatAPI interface {
		// CheckingAvailability - error: FactoryErrPrintFormatNotFound or Failed
		CheckingAvailability(ctx context.Context, id mrtype.KeyInt32) error
	}
)