package notificationv1

// NotificationSendEvent represents the payload for sending a notification.
type NotificationSendEvent struct {
	IdempotencyKey string            `json:"idempotency_key"`
	Type           string            `json:"type"`
	Channel        string            `json:"channel"`
	Channels       []string          `json:"channels"` // New field for multi-channel support
	Priority       string            `json:"priority"`
	Recipient      Recipient         `json:"recipient"`
	TemplateData   map[string]string `json:"template_data"`
	RetryCount     int               `json:"retry_count"`
	LastError      string            `json:"last_error"`
}

// Recipient defines the target user domain.
type Recipient struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
}
