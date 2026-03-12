package models

import (
	"go-link/common/pkg/database/mongodb"
)

// WebhookConfig MongoDB model
type WebhookConfig struct {
	*mongodb.BaseModel `bson:",inline"`
	TenantID           string   `bson:"tenant_id"`
	URL                string   `bson:"url"`
	Secret             string   `bson:"secret"` // HMAC-SHA256 signing secret
	EventTypes         []string `bson:"event_types"`
	IsActive           bool     `bson:"is_active"`
}
