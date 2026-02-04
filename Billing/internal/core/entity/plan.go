package entity

import "time"

// Plan represents a subscription plan.
type Plan struct {
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
