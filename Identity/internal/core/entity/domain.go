package entity

import "time"

// Domain represents a custom domain for a tenant.
type Domain struct {
	ID         int       `json:"id"`
	Domain     string    `json:"domain"`
	IsVerified bool      `json:"is_verified"`
	TenantID   int       `json:"tenant_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
