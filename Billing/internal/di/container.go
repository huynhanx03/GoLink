package di

// Container holds all dependency containers for the Billing service.
type Container struct {
	InvoiceContainer      InvoiceContainer
	PlanContainer         PlanContainer
	SubscriptionContainer SubscriptionContainer
}

// GlobalContainer is the global instance of Container.
var GlobalContainer *Container
