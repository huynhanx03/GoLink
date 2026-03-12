package entity

import (
	"time"
)

// Notification is the core domain model for a notification record.
type Notification struct {
	ID             string
	IdempotencyKey string
	Type           string
	Channel        string
	Priority       string
	Status         string
	Recipient      Recipient
	Subject        string
	Body           string
	TemplateData   map[string]any
	CollapseKey    string
	IsRead         bool
	ErrorMessage   string
	RetryCount     int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	SentAt         *time.Time
	ExpiresAt      time.Time
}

// Recipient holds the target recipient information.
type Recipient struct {
	UserID string
	Email  string
	Name   string
}

// Status constants for Notification.Status.
const (
	StatusPending = "pending"
	StatusSent    = "sent"
	StatusFailed  = "failed"
	StatusSkipped = "skipped"
)

// Channel constants for Notification.Channel.
const (
	ChannelEmail   = "email"
	ChannelInApp   = "in_app"
	ChannelWebhook = "webhook"
	ChannelFCM     = "fcm"
)

// Priority constants for Notification.Priority.
const (
	PriorityUrgent = "urgent"
	PriorityHigh   = "high"
	PriorityNormal = "normal"
	PriorityLow    = "low"
)
