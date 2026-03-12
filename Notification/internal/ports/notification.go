package ports

import (
	"context"

	"go-link/notification/internal/core/dto"
	"go-link/notification/internal/core/entity"

	commonDto "go-link/common/pkg/dto"

	notificationv1 "github.com/huynhanx03/GoLink/events-contract/notification/v1"
)

// NotificationRepository defines the notification persistence contract.
type NotificationRepository interface {
	Create(ctx context.Context, notification *entity.Notification) error
	Get(ctx context.Context, id string) (*entity.Notification, error)
	FindByUserID(ctx context.Context, userID string, page, pageSize int) ([]*entity.Notification, int64, error)
	MarkAsRead(ctx context.Context, id string) error
	MarkAllAsRead(ctx context.Context, userID string) error
	CountUnread(ctx context.Context, userID string) (int64, error)
}

// NotificationService defines the notification business logic contract.
type NotificationService interface {
	ProcessNotification(ctx context.Context, evt *notificationv1.NotificationSendEvent) error
	Find(ctx context.Context, req *dto.GetNotificationsRequest) (*commonDto.Paginated[*dto.NotificationResponse], error)
	GetUnreadCount(ctx context.Context) (*dto.GetUnreadCountResponse, error)
	MarkAsRead(ctx context.Context, id string) error
	MarkAllAsRead(ctx context.Context) error
	ProcessDigest(ctx context.Context, userID, collapseKey string, notificationIDs []string) error
}

// NotificationConsumer defines the contract for consuming notification events.
type NotificationConsumer interface {
	Start(ctx context.Context) error
	StartRetryConsumer(ctx context.Context) error
	Stop() error
}

// DigestService defines the contract for notification aggregation.
type DigestService interface {
	ShouldDigest(notification *entity.Notification) bool
	AddToDigest(ctx context.Context, notification *entity.Notification) error
	ScanPendingDigests(ctx context.Context) ([]string, error)
	ConsumeDigest(ctx context.Context, key string) ([]string, error)
}

// DigestWorker defines the contract for periodic digest processing.
type DigestWorker interface {
	Start(ctx context.Context) error
	Stop()
}
