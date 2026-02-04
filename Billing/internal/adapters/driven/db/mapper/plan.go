package mapper

import (
	"go-link/billing/internal/adapters/driven/db/ent/generate"
	"go-link/billing/internal/core/entity"
)

func ToPlanEntity(e *generate.Plan) *entity.Plan {
	if e == nil {
		return nil
	}
	return &entity.Plan{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		BasePrice:   e.BasePrice,
		Period:      e.Period,
		Features:    e.Features,
		IsActive:    e.IsActive,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}
