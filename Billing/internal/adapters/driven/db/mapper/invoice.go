package mapper

import (
	"go-link/billing/internal/adapters/driven/db/ent/generate"
	"go-link/billing/internal/core/entity"
)

func ToInvoiceEntity(e *generate.Invoice) *entity.Invoice {
	if e == nil {
		return nil
	}

	extID := e.PaymentID
	var extIDPtr *string
	if extID != "" {
		extIDPtr = &extID
	}

	return &entity.Invoice{
		ID:             e.ID,
		SubscriptionID: e.SubscriptionID,
		TenantID:       e.TenantID,
		Amount:         e.Amount,
		Currency:       e.Currency,
		Status:         e.Status.String(),
		PaymentID:      extIDPtr,
		CreatedAt:      e.CreatedAt,
		UpdatedAt:      e.UpdatedAt,
	}
}
