package entity

import (
	"time"

	"github.com/mondegor/go-webcore/mrenum"
	"github.com/mondegor/go-webcore/mrtype"
)

const (
	ModelNamePaperFacture = "admin-api.Dictionaries.PaperFacture"
)

type (
	PaperFacture struct { // DB: printshop_dictionaries.paper_factures
		ID         mrtype.KeyInt32   `json:"id"` // facture_id
		TagVersion int32             `json:"tagVersion"`
		Caption    string            `json:"caption" sort:"caption,default"`
		Status     mrenum.ItemStatus `json:"status"`
		CreatedAt  time.Time         `json:"createdAt" sort:"createdAt"`
		UpdatedAt  *time.Time        `json:"updatedAt,omitempty" sort:"updatedAt"`
	}

	PaperFactureParams struct {
		Filter PaperFactureListFilter
		Sorter mrtype.SortParams
		Pager  mrtype.PageParams
	}

	PaperFactureListFilter struct {
		SearchText string
		Statuses   []mrenum.ItemStatus
	}
)
