package dto

// SendNotificationRequest is the request payload for sending a notification.
type SendNotificationRequest struct {
	Type           string         `json:"type" binding:"required"`
	Channel        string         `json:"channel" binding:"required"`
	Priority       string         `json:"priority"`
	RecipientEmail string         `json:"recipient_email" binding:"required"`
	RecipientName  string         `json:"recipient_name"`
	TemplateData   map[string]any `json:"template_data"`
}

// SendNotificationResponse is the response payload after sending a notification.
type SendNotificationResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// GetNotificationsRequest is the request for listing paginated notifications.
type GetNotificationsRequest struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

// NotificationResponse is the response payload for a single notification.
type NotificationResponse struct {
	ID           string         `json:"id"`
	Type         string         `json:"type"`
	Channel      string         `json:"channel"`
	Priority     string         `json:"priority"`
	Status       string         `json:"status"`
	Subject      string         `json:"subject,omitempty"`
	Body         string         `json:"body,omitempty"`
	TemplateData map[string]any `json:"template_data,omitempty"`
	CollapseKey  string         `json:"collapse_key,omitempty"`
	IsRead       bool           `json:"is_read"`
	ErrorMessage string         `json:"error_message,omitempty"`
	CreatedAt    string         `json:"created_at"`
	SentAt       *string        `json:"sent_at,omitempty"`
}

// MarkAsReadRequest is the request for marking a single notification as read.
type MarkAsReadRequest struct {
	ID string `uri:"id" binding:"required"`
}

// GetUnreadCountResponse holds the unread notification count.
type GetUnreadCountResponse struct {
	Count int64 `json:"count"`
}

// MarkAllAsReadRequest is an empty request for marking all notifications as read.
type MarkAllAsReadRequest struct{}
