package dto

import "time"

// CreateSubscriptionRequest represents the request to create a subscription.
type CreateSubscriptionRequest struct {
	UserID int64
	PlanID int
}

// CreateSubscriptionResponse represents the response after creating a subscription.
type CreateSubscriptionResponse struct {
	SubscriptionID int64
}

// SubscriptionResponse represents subscription data.
type SubscriptionResponse struct {
	ID        int64     `json:"id"`
	TenantID  int64     `json:"tenant_id"`
	PlanID    int       `json:"plan_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// PlanResponse represents plan data.
type PlanResponse struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	BasePrice float64 `json:"base_price"`
	Period    string  `json:"period"`
}

// CreateInvoiceRequest represents request to create an invoice.
type CreateInvoiceRequest struct {
	SubscriptionID int64   `json:"subscription_id"`
	TenantID       int64   `json:"tenant_id"`
	Amount         float64 `json:"amount"`
	Currency       string  `json:"currency"`
	Status         string  `json:"status"`
}

// InvoiceResponse represents invoice data.
type InvoiceResponse struct {
	ID             int64     `json:"id"`
	SubscriptionID int64     `json:"subscription_id"`
	TenantID       int64     `json:"tenant_id"`
	Amount         float64   `json:"amount"`
	Currency       string    `json:"currency"`
	Status         string    `json:"status"`
	PaymentID      *string   `json:"payment_id,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
}

// UpdateInvoiceRequest represents request to update an invoice.
type UpdateInvoiceRequest struct {
	Status    string `json:"status"`
	PaymentID string `json:"payment_id,omitempty"`
}

// UpdateSubscriptionRequest represents request to update a subscription.
type UpdateSubscriptionRequest struct {
	PlanID int `json:"plan_id"`
}
