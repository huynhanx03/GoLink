package entity

import (
	"time"
)

// UserPreference stores per-user notification channel opt-in settings and quiet hours.
type UserPreference struct {
	ID              string
	UserID          string
	EmailEnabled    bool
	InAppEnabled    bool
	WebhookEnabled  bool
	QuietHoursStart *int
	QuietHoursEnd   *int
	Timezone        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// DefaultPreference returns a preference with all channels enabled for the given user.
func DefaultPreference(userID string) *UserPreference {
	now := time.Now()
	return &UserPreference{
		UserID:         userID,
		EmailEnabled:   true,
		InAppEnabled:   true,
		WebhookEnabled: true,
		Timezone:       "UTC",
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}
