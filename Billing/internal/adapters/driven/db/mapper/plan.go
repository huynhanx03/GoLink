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
		ID:        e.ID,
		Name:      e.Name,
		BasePrice: e.BasePrice,
		Period:    e.Period,
		Limits:    e.Limits,
		IsActive:  e.IsActive,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
