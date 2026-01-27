package entity

import "time"

// Tenant represents a tenant organization.
type Tenant struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	TierID    int       `json:"tier_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
