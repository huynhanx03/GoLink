package http

import (
	"context"

	"go-link/common/pkg/common/http/handler"
	commonDto "go-link/common/pkg/dto"
	"go-link/notification/internal/core/dto"
	"go-link/notification/internal/ports"
)

// NotificationHandler defines the HTTP handler interface for notifications.
type NotificationHandler interface {
	Find(ctx context.Context, req *dto.GetNotificationsRequest) (*commonDto.Paginated[*dto.NotificationResponse], error)
	GetUnreadCount(ctx context.Context, _ *struct{}) (*dto.GetUnreadCountResponse, error)
	MarkAsRead(ctx context.Context, req *dto.MarkAsReadRequest) (*struct{}, error)
	MarkAllAsRead(ctx context.Context, _ *struct{}) (*struct{}, error)
}

type notificationHandler struct {
	handler.BaseHandler
	service ports.NotificationService
}

// NewNotificationHandler creates a new NotificationHandler.
func NewNotificationHandler(service ports.NotificationService) NotificationHandler {
	return &notificationHandler{service: service}
}

// Find retrieves a paginated list of notifications.
func (h *notificationHandler) Find(ctx context.Context, req *dto.GetNotificationsRequest) (*commonDto.Paginated[*dto.NotificationResponse], error) {
	return h.service.Find(ctx, req)
}

// GetUnreadCount retrieves the number of unread notifications.
func (h *notificationHandler) GetUnreadCount(ctx context.Context, _ *struct{}) (*dto.GetUnreadCountResponse, error) {
	return h.service.GetUnreadCount(ctx)
}

// MarkAsRead marks a specific notification as read.
func (h *notificationHandler) MarkAsRead(ctx context.Context, req *dto.MarkAsReadRequest) (*struct{}, error) {
	return &struct{}{}, h.service.MarkAsRead(ctx, req.ID)
}

// MarkAllAsRead marks all user notifications as read.
func (h *notificationHandler) MarkAllAsRead(ctx context.Context, _ *struct{}) (*struct{}, error) {
	return &struct{}{}, h.service.MarkAllAsRead(ctx)
}
