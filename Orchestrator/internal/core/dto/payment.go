package dto

type ProcessPaymentRequest struct {
	InvoiceID int64   `json:"invoice_id"`
	TenantID  int64   `json:"tenant_id"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
}

type PaymentResponse struct {
	PaymentID    int64  `json:"payment_id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message,omitempty"`
}
