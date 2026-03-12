package models

import (
	"time"

	"go-link/common/pkg/database/mongodb"
)

// DeliveryLog MongoDB model
type DeliveryLog struct {
	*mongodb.BaseModel `bson:",inline"`
	NotificationID     string        `bson:"notification_id"`
	Channel            string        `bson:"channel"`
	Status             string        `bson:"status"` // "success", "failed", "skipped"
	Attempt            int           `bson:"attempt"`
	ErrorMessage       string        `bson:"error_message,omitempty"`
	ResponseCode       int           `bson:"response_code,omitempty"`
	Duration           time.Duration `bson:"duration"` // delivery latency in nanoseconds
}
