package dto

type ProcessPaymentRequest struct {
	InvoiceID int64   `json:"invoice_id" validate:"required"`
	TenantID  int64   `json:"tenant_id" validate:"required"`
	Amount    float64 `json:"amount" validate:"required,gt=0"`
	Currency  string  `json:"currency" validate:"required,len=3"`
}

type PaymentResponse struct {
	PaymentID    int64  `json:"payment_id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message,omitempty"`
}
