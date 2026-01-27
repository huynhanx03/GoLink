package entity

import "time"

// Subscription represents a tenant subscription.
type Subscription struct {
	ID                 int       `json:"id"`
	TenantID           int       `json:"tenant_id"`
	PlanID             int       `json:"plan_id"`
	Status             string    `json:"status"`
	CurrentPeriodStart time.Time `json:"current_period_start"`
	CurrentPeriodEnd   time.Time `json:"current_period_end"`
	CancelAtPeriodEnd  bool      `json:"cancel_at_period_end"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
