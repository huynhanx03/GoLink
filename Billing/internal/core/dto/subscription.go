package dto

import "time"

// CreateSubscriptionRequest represents request to create a subscription.
type CreateSubscriptionRequest struct {
	TenantID           int       `json:"tenant_id" validate:"required"`
	PlanID             int       `json:"plan_id" validate:"required"`
	Status             string    `json:"status" validate:"required,oneof=ACTIVE PENDING PAID PAST_DUE CANCELED"`
	CurrentPeriodStart time.Time `json:"current_period_start" validate:"required"`
	CurrentPeriodEnd   time.Time `json:"current_period_end" validate:"required,gtfield=CurrentPeriodStart"`
	CancelAtPeriodEnd  bool      `json:"cancel_at_period_end"`
}

// UpdateSubscriptionRequest represents request to update a subscription.
type UpdateSubscriptionRequest struct {
	ID                int        `json:"-" uri:"id"`
	PlanID            *int       `json:"plan_id"`
	Status            *string    `json:"status" validate:"omitempty,oneof=ACTIVE PENDING PAID PAST_DUE CANCELED"`
	CurrentPeriodEnd  *time.Time `json:"current_period_end"`
	CancelAtPeriodEnd *bool      `json:"cancel_at_period_end"`
}

// GetSubscriptionRequest represents request to get a subscription by ID.
type GetSubscriptionRequest struct {
	ID int `uri:"id" validate:"required"`
}

// DeleteSubscriptionRequest represents request to delete a subscription.
type DeleteSubscriptionRequest struct {
	ID int `uri:"id" validate:"required"`
}

// SubscriptionResponse represents subscription data in API response.
type SubscriptionResponse struct {
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
