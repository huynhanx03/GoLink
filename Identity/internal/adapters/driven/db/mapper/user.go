package mapper

import (
	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/core/entity"
)

// ToUserEntity converts Ent User model to domain entity.
func ToUserEntity(m *generate.User) *entity.User {
	if m == nil {
		return nil
	}
	return &entity.User{
		ID:        m.ID,
		Username:  m.Username,
		IsAdmin:   m.IsAdmin,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

// ToUserModel converts domain entity to Ent User model.
func ToUserModel(e *entity.User) *generate.User {
	if e == nil {
		return nil
	}
	return &generate.User{
		ID:        e.ID,
		Username:  e.Username,
		IsAdmin:   e.IsAdmin,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
