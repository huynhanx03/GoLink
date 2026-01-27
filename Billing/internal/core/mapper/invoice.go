package mapper

import (
	"go-link/billing/internal/core/dto"
	"go-link/billing/internal/core/entity"
)

// ToInvoiceResponse converts Invoice entity to InvoiceResponse DTO.
func ToInvoiceResponse(e *entity.Invoice) *dto.InvoiceResponse {
	if e == nil {
		return nil
	}
	return &dto.InvoiceResponse{
		ID:             e.ID,
		SubscriptionID: e.SubscriptionID,
		TenantID:       e.TenantID,
		Amount:         e.Amount,
		Currency:       e.Currency,
		Status:         e.Status,
		PaymentID:      e.PaymentID,
		CreatedAt:      e.CreatedAt,
		UpdatedAt:      e.UpdatedAt,
	}
}

// ToInvoiceEntityFromCreate converts CreateInvoiceRequest to Invoice entity.
func ToInvoiceEntityFromCreate(req *dto.CreateInvoiceRequest) *entity.Invoice {
	return &entity.Invoice{
		SubscriptionID: req.SubscriptionID,
		TenantID:       req.TenantID,
		Amount:         req.Amount,
		Currency:       req.Currency,
		Status:         req.Status,
		PaymentID:      req.PaymentID,
	}
}
