package di

import "go-link/billing/global"

// SetupDependencies initializes all dependencies and returns the container.
func SetupDependencies() *Container {
	client := global.EntClient
	log := global.LoggerZap

	invoiceContainer := InitInvoiceDependencies(client, global.Redis, log)
	planContainer := InitPlanDependencies(client, global.Tinylfu, log)
	subscriptionContainer := InitSubscriptionDependencies(client, planContainer.Repository, global.Redis, log)

	container := &Container{
		InvoiceContainer:      invoiceContainer,
		PlanContainer:         planContainer,
		SubscriptionContainer: subscriptionContainer,
	}

	GlobalContainer = container
	return container
}
