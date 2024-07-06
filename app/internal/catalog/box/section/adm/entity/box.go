package entity

import (
	"time"

	"github.com/mondegor/print-shop-back/pkg/libs/measure"

	"github.com/mondegor/go-webcore/mrenum"
	"github.com/mondegor/go-webcore/mrtype"
)

const (
	ModelNameBox = "admin-api.Catalog.Box" // ModelNameBox - название сущности
)

type (
	// Box - comment struct.
	Box struct { // DB: printshop_catalog.boxes
		ID         mrtype.KeyInt32   `json:"id"` // box_id
		TagVersion int32             `json:"tagVersion"`
		Article    string            `json:"article" sort:"article" upd:"box_article"`
		Caption    string            `json:"caption" sort:"caption,default" upd:"box_caption"`
		Length     measure.Meter     `json:"length" sort:"length" upd:"box_length"`
		Width      measure.Meter     `json:"width" sort:"width" upd:"box_width"`
		Height     measure.Meter     `json:"height" sort:"height" upd:"box_height"`
		Weight     measure.Kilogram  `json:"weight" sort:"weight" upd:"box_weight"`
		Status     mrenum.ItemStatus `json:"status"`
		CreatedAt  time.Time         `json:"createdAt" sort:"createdAt"`
		UpdatedAt  *time.Time        `json:"updatedAt,omitempty" sort:"updatedAt"`
	}

	// BoxParams - comment struct.
	BoxParams struct {
		Filter BoxListFilter
		Sorter mrtype.SortParams
		Pager  mrtype.PageParams
	}

	// BoxListFilter - comment struct.
	BoxListFilter struct {
		SearchText string
		Length     mrtype.RangeInt64
		Width      mrtype.RangeInt64
		Height     mrtype.RangeInt64
		Weight     mrtype.RangeInt64
		Statuses   []mrenum.ItemStatus
	}
)
