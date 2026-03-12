package models

import (
	"go-link/common/pkg/database/mongodb"
)

// UserPreference MongoDB model
type UserPreference struct {
	*mongodb.BaseModel `bson:",inline"`
	UserID             string `bson:"user_id"`
	EmailEnabled       bool   `bson:"email_enabled"`
	InAppEnabled       bool   `bson:"in_app_enabled"`
	WebhookEnabled     bool   `bson:"webhook_enabled"`
	QuietHoursStart    *int   `bson:"quiet_hours_start,omitempty"` // Hour 0-23
	QuietHoursEnd      *int   `bson:"quiet_hours_end,omitempty"`
	Timezone           string `bson:"timezone"`
}
