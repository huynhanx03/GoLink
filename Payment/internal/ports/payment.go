package ports

import (
	"context"

	"go-link/payment/internal/core/dto"
)

type PaymentService interface {
	ProcessPayment(ctx context.Context, req *dto.ProcessPaymentRequest) (*dto.PaymentResponse, error)
}
