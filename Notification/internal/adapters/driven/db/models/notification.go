package models

import (
	"time"

	"go-link/common/pkg/database/mongodb"
)

// Notification represents the MongoDB document for a notification log.
type Notification struct {
	*mongodb.BaseModel `bson:",inline"`
	IdempotencyKey     string         `bson:"idempotency_key"`
	Type               string         `bson:"type"`
	Channel            string         `bson:"channel"`
	Priority           string         `bson:"priority"`
	Status             string         `bson:"status"`
	Recipient          Recipient      `bson:"recipient"`
	Subject            string         `bson:"subject,omitempty"`
	Body               string         `bson:"body,omitempty"`
	TemplateData       map[string]any `bson:"template_data,omitempty"`
	CollapseKey        string         `bson:"collapse_key,omitempty"`
	IsRead             bool           `bson:"is_read"`
	ErrorMessage       string         `bson:"error_message,omitempty"`
	RetryCount         int            `bson:"retry_count"`
	SentAt             *time.Time     `bson:"sent_at,omitempty"`
	ExpiresAt          time.Time      `bson:"expires_at"`
}

// Recipient holds the target recipient information.
type Recipient struct {
	UserID string `bson:"user_id,omitempty"`
	Email  string `bson:"email,omitempty"`
	Name   string `bson:"name,omitempty"`
}
