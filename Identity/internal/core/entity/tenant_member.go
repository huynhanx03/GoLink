package entity

import "time"

// TenantMember represents a user's membership in a tenant.
type TenantMember struct {
	ID        int       `json:"id"`
	TenantID  int       `json:"tenant_id"`
	UserID    int       `json:"user_id"`
	RoleID    int       `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
