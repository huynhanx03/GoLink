package grpc

import (
	"context"

	paymentv1 "go-link/common/gen/go/payment/v1"
	"go-link/payment/internal/core/dto"
	"go-link/payment/internal/ports"
)

type PaymentServer struct {
	paymentv1.UnimplementedPaymentServiceServer
	service ports.PaymentService
}

func NewPaymentServer(service ports.PaymentService) *PaymentServer {
	return &PaymentServer{service: service}
}

func (s *PaymentServer) ProcessPayment(ctx context.Context, req *paymentv1.ProcessPaymentRequest) (*paymentv1.ProcessPaymentResponse, error) {
	dtoReq := dto.ProcessPaymentRequest{
		InvoiceID: req.InvoiceId,
		TenantID:  req.TenantId,
		Amount:    req.Amount,
		Currency:  req.Currency,
	}

	res, err := s.service.ProcessPayment(ctx, &dtoReq)
	if err != nil {
		return nil, err
	}

	return &paymentv1.ProcessPaymentResponse{
		PaymentId:    res.PaymentID,
		Status:       res.Status,
		ErrorMessage: res.ErrorMessage,
	}, nil
}
