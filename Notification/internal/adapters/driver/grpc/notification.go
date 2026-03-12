package grpc

import (
	"context"

	notificationv1 "go-link/common/gen/go/notification/v1"
	"go-link/notification/internal/ports"

	eventv1 "github.com/huynhanx03/GoLink/events-contract/notification/v1"
)

// NotificationServer implements the gRPC NotificationServiceServer.
type NotificationServer struct {
	notificationv1.UnimplementedNotificationServiceServer
	service ports.NotificationService
	repo    ports.NotificationRepository
	counter ports.UnreadCounter
}

// NewNotificationServer creates a new gRPC NotificationServer.
func NewNotificationServer(
	svc ports.NotificationService,
	repo ports.NotificationRepository,
	counter ports.UnreadCounter,
) *NotificationServer {
	return &NotificationServer{
		service: svc,
		repo:    repo,
		counter: counter,
	}
}

// SendNotification handles an RPC call to send a notification through the service.
func (s *NotificationServer) SendNotification(ctx context.Context, req *notificationv1.SendNotificationRequest) (*notificationv1.SendNotificationResponse, error) {
	recipient := eventv1.Recipient{}
	if req.Recipient != nil {
		recipient = eventv1.Recipient{
			UserID: req.Recipient.UserId,
			Email:  req.Recipient.Email,
			Name:   req.Recipient.Name,
		}
	}

	// Convert map[string]string to map[string]string for TemplateData.
	templateData := make(map[string]string, len(req.TemplateData))
	for k, v := range req.TemplateData {
		templateData[k] = v
	}

	evt := &eventv1.NotificationSendEvent{
		IdempotencyKey: req.IdempotencyKey,
		Type:           req.Type,
		Channel:        req.Channel,
		Priority:       req.Priority,
		Recipient:      recipient,
		TemplateData:   templateData,
	}

	if err := s.service.ProcessNotification(ctx, evt); err != nil {
		return &notificationv1.SendNotificationResponse{Success: false}, err
	}

	return &notificationv1.SendNotificationResponse{Success: true}, nil
}

// GetUnreadCount returns the unread notification count for a user.
func (s *NotificationServer) GetUnreadCount(ctx context.Context, req *notificationv1.GetUnreadCountRequest) (*notificationv1.GetUnreadCountResponse, error) {
	count, err := s.counter.Get(ctx, req.UserId)
	if err != nil {
		// Fall back to MongoDB count.
		count, err = s.repo.CountUnread(ctx, req.UserId)
		if err != nil {
			return nil, err
		}
	}

	return &notificationv1.GetUnreadCountResponse{Count: count}, nil
}

// GetNotifications returns a paginated list of notifications for a user.
func (s *NotificationServer) GetNotifications(ctx context.Context, req *notificationv1.GetNotificationsRequest) (*notificationv1.GetNotificationsResponse, error) {
	page := int(req.Page)
	pageSize := int(req.PageSize)
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}

	notifications, total, err := s.repo.FindByUserID(ctx, req.UserId, page, pageSize)
	if err != nil {
		return nil, err
	}

	items := make([]*notificationv1.NotificationItem, 0, len(notifications))
	for _, n := range notifications {
		items = append(items, &notificationv1.NotificationItem{
			Id:        n.ID,
			Type:      n.Type,
			Channel:   n.Channel,
			Status:    n.Status,
			Subject:   n.Subject,
			Body:      n.Body,
			IsRead:    n.IsRead,
			CreatedAt: n.CreatedAt.UTC().String(),
		})
	}

	return &notificationv1.GetNotificationsResponse{
		Notifications: items,
		Total:         total,
	}, nil
}
