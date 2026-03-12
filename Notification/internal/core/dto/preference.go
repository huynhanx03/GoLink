package dto

// UpdatePreferenceRequest is the request body for updating notification preferences.
type UpdatePreferenceRequest struct {
	EmailEnabled    bool   `json:"email_enabled"`
	InAppEnabled    bool   `json:"in_app_enabled"`
	WebhookEnabled  bool   `json:"webhook_enabled"`
	QuietHoursStart *int   `json:"quiet_hours_start,omitempty"` // Hour 0-23
	QuietHoursEnd   *int   `json:"quiet_hours_end,omitempty"`
	Timezone        string `json:"timezone"`
}

// PreferenceResponse is the response payload for a user's notification preferences.
type PreferenceResponse struct {
	UserID          string `json:"user_id"`
	EmailEnabled    bool   `json:"email_enabled"`
	InAppEnabled    bool   `json:"in_app_enabled"`
	WebhookEnabled  bool   `json:"webhook_enabled"`
	QuietHoursStart *int   `json:"quiet_hours_start,omitempty"`
	QuietHoursEnd   *int   `json:"quiet_hours_end,omitempty"`
	Timezone        string `json:"timezone"`
}
