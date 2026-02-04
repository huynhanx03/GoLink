package grpc

import (
	"context"

	paymentv1 "go-link/common/gen/go/payment/v1"
	"go-link/common/pkg/grpc"
	"go-link/common/pkg/settings"
	"go-link/orchestrator/internal/core/dto"
	"go-link/orchestrator/internal/ports"
)

type paymentClientAdapter struct {
	client paymentv1.PaymentServiceClient
}

func NewPaymentClient(cfg settings.GRPCService) (ports.PaymentClient, error) {
	conn, err := grpc.NewClientConn(cfg)
	if err != nil {
		return nil, err
	}
	return &paymentClientAdapter{
		client: paymentv1.NewPaymentServiceClient(conn),
	}, nil
}

func (a *paymentClientAdapter) ProcessPayment(ctx context.Context, req dto.ProcessPaymentRequest) (*dto.PaymentResponse, error) {
	protoReq := &paymentv1.ProcessPaymentRequest{
		InvoiceId: req.InvoiceID,
		TenantId:  req.TenantID,
		Amount:    req.Amount,
		Currency:  req.Currency,
	}

	resp, err := a.client.ProcessPayment(ctx, protoReq)
	if err != nil {
		return nil, err
	}

	return &dto.PaymentResponse{
		PaymentID:    resp.PaymentId,
		Status:       resp.Status,
		ErrorMessage: resp.ErrorMessage,
	}, nil
}
