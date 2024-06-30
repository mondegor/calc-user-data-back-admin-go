package entity

import (
	"github.com/mondegor/go-webcore/mrtype"
)

const (
	ModelNamePaperFacture = "public-api.Dictionaries.PaperFacture" // ModelNamePaperFacture - название сущности
)

type (
	// PaperFacture - comment struct.
	PaperFacture struct { // DB: printshop_dictionaries.paper_factures
		ID      mrtype.KeyInt32 `json:"id"` // facture_id
		Caption string          `json:"caption"`
	}

	// PaperFactureParams - comment struct.
	PaperFactureParams struct{}
)
