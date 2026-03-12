package entity

import (
	"time"
)

// DeliveryLog records the outcome of a single notification delivery attempt per channel.
type DeliveryLog struct {
	ID             string
	NotificationID string
	Channel        string
	Status         string // "success", "failed", "skipped"
	Attempt        int
	ErrorMessage   string
	ResponseCode   int
	Duration       time.Duration // delivery latency in nanoseconds
	CreatedAt      time.Time
}
