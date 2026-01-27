package builder

import (
	"context"

	"go-link/billing/global"
	"go-link/billing/internal/adapters/driven/db/ent/generate"
	"go-link/billing/internal/adapters/driven/db/ent/generate/invoice"
	"go-link/billing/internal/core/entity"
)

func BuildCreateInvoice(ctx context.Context, e *entity.Invoice) *generate.InvoiceCreate {
	client := global.EntClient.DB(ctx)
	create := client.Invoice.Create().
		SetSubscriptionID(e.SubscriptionID).
		SetTenantID(e.TenantID).
		SetAmount(e.Amount).
		SetCurrency(e.Currency).
		SetStatus(invoice.Status(e.Status))

	if e.PaymentID != nil {
		create.SetPaymentID(*e.PaymentID)
	}

	return create
}

func BuildUpdateInvoice(ctx context.Context, e *entity.Invoice) *generate.InvoiceUpdateOne {
	client := global.EntClient.DB(ctx)
	update := client.Invoice.UpdateOneID(e.ID).
		SetStatus(invoice.Status(e.Status))

	if e.PaymentID != nil {
		update.SetPaymentID(*e.PaymentID)
	} else {
		update.ClearPaymentID()
	}

	return update
}
