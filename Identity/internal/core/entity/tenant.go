package entity

import "time"

// Tenant represents a tenant organization.
type Tenant struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	PlanID    int       `json:"plan_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
