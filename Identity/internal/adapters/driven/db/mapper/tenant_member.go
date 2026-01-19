package mapper

import (
	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/core/entity"
)

// ToTenantMemberEntity converts Ent TenantMember model to domain entity.
func ToTenantMemberEntity(m *generate.TenantMember) *entity.TenantMember {
	if m == nil {
		return nil
	}
	return &entity.TenantMember{
		ID:        m.ID,
		TenantID:  m.TenantID,
		UserID:    m.UserID,
		RoleID:    m.RoleID,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

// ToTenantMemberModel converts domain entity to Ent TenantMember model.
func ToTenantMemberModel(e *entity.TenantMember) *generate.TenantMember {
	if e == nil {
		return nil
	}
	return &generate.TenantMember{
		ID:        e.ID,
		TenantID:  e.TenantID,
		UserID:    e.UserID,
		RoleID:    e.RoleID,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
