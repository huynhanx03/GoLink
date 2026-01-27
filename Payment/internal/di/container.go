package di

type Container struct {
	PaymentContainer *PaymentContainer
}

var GlobalContainer *Container
