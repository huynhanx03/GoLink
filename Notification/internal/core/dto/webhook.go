package dto

// CreateWebhookRequest is the request body for creating a webhook config.
type CreateWebhookRequest struct {
	URL        string   `json:"url" binding:"required,url"`
	EventTypes []string `json:"event_types"`
	IsActive   bool     `json:"is_active"`
}

// UpdateWebhookRequest is the request body for updating a webhook config.
type UpdateWebhookRequest struct {
	ID         string   `uri:"id" binding:"required"`
	URL        string   `json:"url" binding:"required,url"`
	EventTypes []string `json:"event_types"`
	IsActive   bool     `json:"is_active"`
}

// GetWebhookRequest is the request for getting a single webhook config.
type GetWebhookRequest struct {
	ID string `uri:"id" binding:"required"`
}

// DeleteWebhookRequest is the request for deleting a single webhook config.
type DeleteWebhookRequest struct {
	ID string `uri:"id" binding:"required"`
}

// WebhookResponse is the response payload for a webhook configuration.
type WebhookResponse struct {
	ID         string   `json:"id"`
	TenantID   string   `json:"tenant_id"`
	URL        string   `json:"url"`
	EventTypes []string `json:"event_types"`
	IsActive   bool     `json:"is_active"`
}
