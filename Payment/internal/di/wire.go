package di

func SetupDependencies() *Container {
	container := &Container{
		PaymentContainer: InitPaymentContainer(),
	}
	return container
}
