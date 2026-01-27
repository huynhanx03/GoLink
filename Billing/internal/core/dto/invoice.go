package dto

import "time"

// CreateInvoiceRequest represents request to create an invoice.
type CreateInvoiceRequest struct {
	SubscriptionID int     `json:"subscription_id" validate:"required"`
	TenantID       int     `json:"tenant_id" validate:"required"`
	Amount         float64 `json:"amount" validate:"required,gt=0"`
	Currency       string  `json:"currency" validate:"required,len=3"`
	Status         string  `json:"status" validate:"required,oneof=PENDING PAID FAILED REFUNDED"`
	PaymentID      *string `json:"payment_id"`
}

// UpdateInvoiceRequest represents request to update an invoice.
type UpdateInvoiceRequest struct {
	ID        int     `json:"-" uri:"id"`
	Status    *string `json:"status" validate:"omitempty,oneof=PENDING PAID FAILED REFUNDED"`
	PaymentID *string `json:"payment_id"`
}

// GetInvoiceRequest represents request to get an invoice by ID.
type GetInvoiceRequest struct {
	ID int `uri:"id" validate:"required"`
}

// FindMyInvoicesRequest represents request to find invoices for the current user.
type FindMyInvoicesRequest struct{}

// DeleteInvoiceRequest represents request to delete an invoice.
type DeleteInvoiceRequest struct {
	ID int `uri:"id" validate:"required"`
}

// InvoiceResponse represents invoice data in API response.
type InvoiceResponse struct {
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
