package entity

import "time"

// Invoice represents a billing invoice.
type Invoice struct {
	ID             int       `json:"id"`
	SubscriptionID int       `json:"subscription_id"`
	TenantID       int       `json:"tenant_id"`
	Amount         float64   `json:"amount"`
	Currency       string    `json:"currency"`
	Status         string    `json:"status"`
	PaymentID      *string   `json:"payment_id,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
