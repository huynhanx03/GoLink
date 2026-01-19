package mapper

import (
	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/core/entity"
)

// ToDomainEntity converts Ent Domain model to domain entity.
func ToDomainEntity(m *generate.Domain) *entity.Domain {
	if m == nil {
		return nil
	}
	return &entity.Domain{
		ID:         m.ID,
		Domain:     m.Domain,
		IsVerified: m.IsVerified,
		TenantID:   m.TenantID,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
}

// ToDomainModel converts domain entity to Ent Domain model.
func ToDomainModel(e *entity.Domain) *generate.Domain {
	if e == nil {
		return nil
	}
	return &generate.Domain{
		ID:         e.ID,
		Domain:     e.Domain,
		TenantID:   e.TenantID,
		IsVerified: e.IsVerified,
		CreatedAt:  e.CreatedAt,
		UpdatedAt:  e.UpdatedAt,
	}
}
