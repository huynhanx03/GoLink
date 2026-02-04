package dto

import "time"

// CreatePlanRequest represents request to create a plan.
type CreatePlanRequest struct {
	Name        string                 `json:"name" validate:"required,min=2,max=100"`
	Description string                 `json:"description" validate:"max=200"`
	BasePrice   float64                `json:"base_price" validate:"required,gte=0"`
	Period      string                 `json:"period" validate:"required,oneof=month year"`
	Features    map[string]interface{} `json:"features"`
	IsActive    bool                   `json:"is_active"`
}

// UpdatePlanRequest represents request to update a plan.
type UpdatePlanRequest struct {
	ID          int                     `json:"-" uri:"id"`
	Name        *string                 `json:"name" validate:"omitempty,min=2,max=100"`
	Description *string                 `json:"description" validate:"omitempty,max=200"`
	BasePrice   *float64                `json:"base_price" validate:"omitempty,gte=0"`
	Features    *map[string]interface{} `json:"features"`
	IsActive    *bool                   `json:"is_active"`
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
	ID          int                    `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	BasePrice   float64                `json:"base_price"`
	Period      string                 `json:"period"`
	Features    map[string]interface{} `json:"features,omitempty"`
	IsActive    bool                   `json:"is_active"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}
