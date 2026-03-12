package entity

import (
	"time"
)

// WebhookConfig holds the configuration for a tenant's outbound webhook endpoint.
type WebhookConfig struct {
	ID         string
	TenantID   string
	URL        string
	Secret     string // HMAC-SHA256 signing secret
	EventTypes []string
	IsActive   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
