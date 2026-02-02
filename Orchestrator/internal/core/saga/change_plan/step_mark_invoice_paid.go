package change_plan

import (
	"context"
	"fmt"

	"go-link/orchestrator/internal/constant"
	"go-link/orchestrator/internal/core/dto"
	"go-link/orchestrator/internal/ports"
)

type StepMarkInvoicePaid struct {
	State         *State
	BillingClient ports.BillingClient
}

func (s *StepMarkInvoicePaid) Name() string { return "MarkInvoicePaid" }

func (s *StepMarkInvoicePaid) Execute(ctx context.Context) error {
	req := dto.UpdateInvoiceRequest{
		Status: constant.InvoiceStatusPaid,
		// Convert int64 PaymentID to string for Billing Service
		PaymentID: fmt.Sprintf("%d", s.State.PaymentID),
	}
	return s.BillingClient.UpdateInvoice(ctx, s.State.InvoiceID, req)
}

func (s *StepMarkInvoicePaid) Compensate(ctx context.Context) error {
	return nil
}
