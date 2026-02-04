package change_plan

import (
	"context"
	"net/http"
	"strings"

	"go-link/common/pkg/common/apperr"
	"go-link/common/pkg/common/http/response"
	"go-link/orchestrator/internal/constant"
	"go-link/orchestrator/internal/core/dto"
	"go-link/orchestrator/internal/ports"
)

type StepProcessPayment struct {
	State         *State
	PaymentClient ports.PaymentClient
}

func (s *StepProcessPayment) Name() string { return "ProcessPayment" }

func (s *StepProcessPayment) Execute(ctx context.Context) error {
	req := dto.ProcessPaymentRequest{
		InvoiceID: s.State.InvoiceID,
		TenantID:  s.State.TenantID,
		Amount:    s.State.PlanPrice,
		Currency:  s.State.Currency,
	}

	res, err := s.PaymentClient.ProcessPayment(ctx, req)
	if err != nil {
		return err
	}

	if strings.ToLower(res.Status) != constant.PaymentStatusSuccess {
		return apperr.NewError(
			"StepProcessPayment",
			response.CodeInternalError,
			res.ErrorMessage,
			http.StatusPaymentRequired,
			nil,
		)
	}

	s.State.PaymentID = res.PaymentID
	s.State.Status = res.Status
	return nil
}

func (s *StepProcessPayment) Compensate(ctx context.Context) error {
	// Refund logic
	if s.State.PaymentID != 0 {
		// Log: Payment made but SAGA failed. Refund needed.
	}
	return nil
}
