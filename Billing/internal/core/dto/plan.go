package dto

import "time"

// CreatePlanRequest represents request to create a plan.
type CreatePlanRequest struct {
	Name      string                 `json:"name" validate:"required,min=2,max=100"`
	BasePrice float64                `json:"base_price" validate:"required,gte=0"`
	Period    string                 `json:"period" validate:"required,oneof=month year"`
	Limits    map[string]interface{} `json:"limits"`
	IsActive  bool                   `json:"is_active"`
}

// UpdatePlanRequest represents request to update a plan.
type UpdatePlanRequest struct {
	ID        int                     `json:"-" uri:"id"`
	Name      *string                 `json:"name" validate:"omitempty,min=2,max=100"`
	BasePrice *float64                `json:"base_price" validate:"omitempty,gte=0"`
	Limits    *map[string]interface{} `json:"limits"`
	IsActive  *bool                   `json:"is_active"`
}

// FindPlanRequest represents request to find plans.
type FindPlanRequest struct{}

// FindActivePlanRequest represents request to find active plans.
type FindActivePlanRequest struct{}

// GetPlanRequest represents request to get a plan by ID.
type GetPlanRequest struct {
	ID int `uri:"id" validate:"required"`
}

// DeletePlanRequest represents request to delete a plan.
type DeletePlanRequest struct {
	ID int `uri:"id" validate:"required"`
}

// PlanResponse represents plan data in API response.
type PlanResponse struct {
	ID        int                    `json:"id"`
	Name      string                 `json:"name"`
	BasePrice float64                `json:"base_price"`
	Period    string                 `json:"period"`
	Limits    map[string]interface{} `json:"limits,omitempty"`
	IsActive  bool                   `json:"is_active"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}
