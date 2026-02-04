package change_plan

// State holds the state during the Upgrade Subscription Saga execution.
type State struct {
	// Input
	TenantID       int64
	SubscriptionID int64
	NewPlanID      int
	OldPlanID      int // For compensation

	// Context/Temp
	PlanPrice float64
	Currency  string

	// Output
	InvoiceID int64
	PaymentID int64 // Changed from string to int64
	Status    string
}
