package grpc

import (
	"google.golang.org/grpc"

	paymentv1 "go-link/common/gen/go/payment/v1"
	"go-link/payment/internal/ports"
)

// V1Routes registers the payment service routes
func V1Routes(
	service ports.PaymentService,
) func(srv *grpc.Server) {
	return func(srv *grpc.Server) {
		paymentv1.RegisterPaymentServiceServer(srv, NewPaymentServer(service))
	}
}
