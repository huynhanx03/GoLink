package dto

// CreateSubscriptionRequest represents the request to create a subscription.
type CreateSubscriptionRequest struct {
	UserID int64
	PlanID int
}

// CreateSubscriptionResponse represents the response after creating a subscription.
type CreateSubscriptionResponse struct {
	SubscriptionID int64
}
