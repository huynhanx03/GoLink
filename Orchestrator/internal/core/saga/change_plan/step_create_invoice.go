package change_plan

import (
	"context"

	"go-link/orchestrator/internal/constant"
	"go-link/orchestrator/internal/core/dto"
	"go-link/orchestrator/internal/ports"
)

type StepCreateInvoice struct {
	State         *State
	BillingClient ports.BillingClient
}

func (s *StepCreateInvoice) Name() string { return "CreateInvoice" }

func (s *StepCreateInvoice) Execute(ctx context.Context) error {
	req := dto.CreateInvoiceRequest{
		SubscriptionID: s.State.SubscriptionID,
		TenantID:       s.State.TenantID,
		Amount:         s.State.PlanPrice,
		Currency:       s.State.Currency,
		Status:         constant.InvoiceStatusPending,
	}

	invoice, err := s.BillingClient.CreateInvoice(ctx, req)
	if err != nil {
		return err
	}

	s.State.InvoiceID = invoice.ID
	return nil
}

func (s *StepCreateInvoice) Compensate(ctx context.Context) error {
	if s.State.InvoiceID == 0 {
		return nil
	}

	req := dto.UpdateInvoiceRequest{
		Status: constant.InvoiceStatusCancelled,
	}
	return s.BillingClient.UpdateInvoice(ctx, s.State.InvoiceID, req)
}
