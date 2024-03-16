package http_v1

import (
	entity "print-shop-back/internal/modules/controls/entity/admin-api"
	entity_shared "print-shop-back/internal/modules/controls/entity/shared"
)

type (
	CreateElementTemplateRequest struct {
		ParamName string                         `json:"paramName" validate:"required,min=4,max=32,tag_variable"`
		Caption   string                         `json:"caption" validate:"required,max=64"`
		Type      entity_shared.ElementType      `json:"elementType" validate:"required"`
		Detailing entity_shared.ElementDetailing `json:"elementDetailing" validate:"required"`
	}

	StoreElementTemplateRequest struct {
		TagVersion int32  `json:"tagVersion" validate:"required,gte=1"`
		ParamName  string `json:"paramName" validate:"omitempty,min=4,max=32,tag_variable"`
		Caption    string `json:"caption" validate:"omitempty,max=64"`
	}

	ElementTemplateListResponse struct {
		Items []entity.ElementTemplate `json:"items"`
		Total int64                    `json:"total"`
	}
)
