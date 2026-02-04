package service

import (
	"context"
	"math/rand"

	"go-link/common/pkg/runtime" // Import fast_rand wrapper
	"go-link/payment/global"
	"go-link/payment/internal/constant"
	"go-link/payment/internal/core/dto"
	"go-link/payment/internal/ports"

	"go.uber.org/zap"
)

type paymentService struct{}

func NewPaymentService() ports.PaymentService {
	return &paymentService{}
}

func (s *paymentService) ProcessPayment(ctx context.Context, req *dto.ProcessPaymentRequest) (*dto.PaymentResponse, error) {
	// Generate int64 Payment ID using fastrand
	paymentID := int64(runtime.Unit64())

	randomValue := rand.Intn(100)
	if randomValue < constant.MockFailThreshold {
		global.LoggerZap.Warn("Payment failed for invoice %d, random=%d", zap.Int64("invoice_id", req.InvoiceID), zap.Int("random", randomValue))
		return &dto.PaymentResponse{
			PaymentID:    paymentID,
			Status:       constant.PaymentStatusFailed,
			ErrorMessage: "payment declined by mock processor",
		}, nil
	}

	global.LoggerZap.Info("Payment success for invoice %d, amount=%.2f %s", zap.Int64("invoice_id", req.InvoiceID), zap.Float64("amount", req.Amount), zap.String("currency", req.Currency))
	return &dto.PaymentResponse{
		PaymentID: paymentID,
		Status:    constant.PaymentStatusSuccess,
	}, nil
}
