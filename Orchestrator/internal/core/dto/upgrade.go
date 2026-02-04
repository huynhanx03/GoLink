package dto

// UpgradeSubscriptionRequest represents the request to upgrade a subscription.
type UpgradeSubscriptionRequest struct {
	TenantID       int64 `json:"tenant_id" validate:"required"`
	SubscriptionID int64 `json:"subscription_id" validate:"required"`
	NewPlanID      int   `json:"new_plan_id" validate:"required"`
}

// UpgradeSubscriptionResponse represents the response after upgrading a subscription.
type UpgradeSubscriptionResponse struct {
	InvoiceID int64 `json:"invoice_id"`
}
