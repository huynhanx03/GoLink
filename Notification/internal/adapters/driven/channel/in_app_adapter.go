package channel

import (
	"context"

	"go-link/notification/internal/adapters/driven/cache"
	"go-link/notification/internal/core/entity"
	"go-link/notification/internal/core/service"
	"go-link/notification/internal/ports"
)

type inAppAdapter struct {
	repo    ports.NotificationRepository
	hub     *service.SSEHub
	counter *cache.UnreadCounter
}

// NewInAppAdapter creates a new in-app notification channel adapter.
// It persists the notification to MongoDB, increments the unread counter,
// and broadcasts to any active SSE connections for the recipient.
func NewInAppAdapter(
	repo ports.NotificationRepository,
	hub *service.SSEHub,
	counter *cache.UnreadCounter,
) ports.ChannelAdapter {
	return &inAppAdapter{
		repo:    repo,
		hub:     hub,
		counter: counter,
	}
}

// Send stores the in-app notification, updates the unread counter, and pushes via SSE.
func (a *inAppAdapter) Send(ctx context.Context, notification *entity.Notification) error {
	// Mark as unread on creation.
	notification.IsRead = false

	if err := a.repo.Create(ctx, notification); err != nil {
		return err
	}

	// Increment Redis unread counter for the recipient.
	if notification.Recipient.UserID != "" {
		_ = a.counter.Increment(ctx, notification.Recipient.UserID)
	}

	// Push to any active SSE connections for this user (non-blocking).
	a.hub.Broadcast(notification.Recipient.UserID, notification)

	return nil
}

// Type returns the channel identifier.
func (a *inAppAdapter) Type() string {
	return entity.ChannelInApp
}
