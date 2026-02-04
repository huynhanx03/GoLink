package di

import (
	"go-link/payment/internal/core/service"
	"go-link/payment/internal/ports"
)

type PaymentContainer struct {
	Service ports.PaymentService
}

func InitPaymentContainer() *PaymentContainer {
	return &PaymentContainer{
		Service: service.NewPaymentService(),
	}
}
