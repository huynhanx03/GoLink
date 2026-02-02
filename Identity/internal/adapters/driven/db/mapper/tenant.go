package mapper

import (
	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/core/entity"
)

// ToTenantEntity converts Ent Tenant model to domain entity.
func ToTenantEntity(m *generate.Tenant) *entity.Tenant {
	if m == nil {
		return nil
	}
	return &entity.Tenant{
		ID:        m.ID,
		Name:      m.Name,
		PlanID:    m.PlanID,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

// ToTenantModel converts domain entity to Ent Tenant model.
func ToTenantModel(e *entity.Tenant) *generate.Tenant {
	if e == nil {
		return nil
	}
	return &generate.Tenant{
		ID:        e.ID,
		Name:      e.Name,
		PlanID:    e.PlanID,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
