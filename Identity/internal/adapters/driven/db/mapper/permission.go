package mapper

import (
	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/core/entity"
)

// ToPermissionEntity converts Ent Permission model to domain entity.
func ToPermissionEntity(m *generate.Permission) *entity.Permission {
	if m == nil {
		return nil
	}
	return &entity.Permission{
		ID:          m.ID,
		RoleID:      m.RoleID,
		ResourceID:  m.ResourceID,
		Description: m.Description,
		Scopes:      m.Scopes,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

// ToPermissionModel converts domain entity to Ent Permission model.
func ToPermissionModel(e *entity.Permission) *generate.Permission {
	if e == nil {
		return nil
	}
	return &generate.Permission{
		ID:          e.ID,
		RoleID:      e.RoleID,
		ResourceID:  e.ResourceID,
		Description: e.Description,
		Scopes:      e.Scopes,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}
