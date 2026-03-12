package service

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"

	"go-link/common/pkg/constraints"
	"go-link/notification/global"
	"go-link/notification/internal/core/dto"
	"go-link/notification/internal/core/entity"
	"go-link/notification/internal/ports"
	"go-link/notification/internal/templates"

	notificationv1 "github.com/huynhanx03/GoLink/events-contract/notification/v1"
)

const (
	rateLimitPerHour int64 = 100
	rateLimitWindow        = time.Hour
)

// templateMapping for types
var templateMapping = map[string]struct {
	Template string
	Subject  string
}{
	"forgot-password-otp": {Template: "forgot-password-otp", Subject: "Password Reset - GoLink"},
	"account-alert":       {Subject: "Security Alert", Template: "Your account was accessed from a new device."},
	"welcome-email":       {Subject: "Welcome to GoLink!", Template: "Hi {{name}}, thanks for joining us!"},
	"digest-summary":      {Subject: "Notification Summary", Template: "You have {{notification_count}} new updates for {{collapse_key}}."},
}

type notificationService struct {
	repo              ports.NotificationRepository
	idempotency       ports.IdempotencyChecker
	channels          map[string]ports.ChannelAdapter
	preferenceService ports.PreferenceService
	rateLimiter       *RateLimiter
	deliveryLogRepo   ports.DeliveryLogRepository
	digestService     ports.DigestService
}

// NewNotificationService creates service
func NewNotificationService(
	repo ports.NotificationRepository,
	idempotency ports.IdempotencyChecker,
	channels map[string]ports.ChannelAdapter,
	preferenceService ports.PreferenceService,
	rateLimiter *RateLimiter,
	deliveryLogRepo ports.DeliveryLogRepository,
	digestService ports.DigestService,
) ports.NotificationService {
	return &notificationService{
		repo:              repo,
		idempotency:       idempotency,
		channels:          channels,
		preferenceService: preferenceService,
		rateLimiter:       rateLimiter,
		deliveryLogRepo:   deliveryLogRepo,
		digestService:     digestService,
	}
}

// ProcessNotification executes send lifecycle
func (s *notificationService) ProcessNotification(ctx context.Context, evt *notificationv1.NotificationSendEvent) error {
	// Idempotency check
	if evt.IdempotencyKey != "" {
		acquired, err := s.idempotency.TryAcquire(ctx, evt.IdempotencyKey)
		if err != nil {
			global.LoggerZap.Warn("Idempotency check failed, proceeding", zap.Error(err))
		}
		if !acquired {
			global.LoggerZap.Info("Duplicate notification skipped",
				zap.String("idempotency_key", evt.IdempotencyKey))
			return nil
		}
	}

	// Prepare channels to send
	channels := evt.Channels
	if len(channels) == 0 && evt.Channel != "" {
		channels = []string{evt.Channel}
	}

	if len(channels) == 0 {
		return fmt.Errorf("no channels specified")
	}

	// Render template once
	subject, body := "", ""
	if mapping, ok := templateMapping[evt.Type]; ok {
		data := make(map[string]any)
		for k, v := range evt.TemplateData {
			data[k] = v
		}

		rendered, err := templates.Render(mapping.Template, data)
		if err != nil {
			global.LoggerZap.Error("Template render failed", zap.Error(err), zap.String("type", evt.Type))
			return fmt.Errorf("render template: %w", err)
		}
		body = rendered
		subject = mapping.Subject
	}

	// Send to each channel
	var lastErr error
	for _, channel := range channels {
		// User preference check
		prefCtx := context.WithValue(ctx, constraints.ContextKeyUserID, evt.Recipient.UserID)
		if allowed, _ := s.preferenceService.IsChannelAllowed(prefCtx, channel); !allowed {
			global.LoggerZap.Info("Channel skipped by preference",
				zap.String("user_id", evt.Recipient.UserID),
				zap.String("channel", channel),
			)
			continue
		}

		// Rate limit check
		if allowed, _ := s.rateLimiter.Allow(ctx, evt.Recipient.UserID, channel, rateLimitPerHour, rateLimitWindow); !allowed {
			global.LoggerZap.Warn("Notification rate limited",
				zap.String("user_id", evt.Recipient.UserID),
				zap.String("channel", channel),
			)
			continue
		}

		adapter, ok := s.channels[channel]
		if !ok {
			global.LoggerZap.Error("Unsupported channel", zap.String("channel", channel))
			continue
		}

		// Build entity
		notification := &entity.Notification{
			IdempotencyKey: evt.IdempotencyKey,
			Type:           evt.Type,
			Channel:        channel,
			Priority:       evt.Priority,
			Status:         entity.StatusPending,
			Recipient: entity.Recipient{
				UserID: evt.Recipient.UserID,
				Email:  evt.Recipient.Email,
				Name:   evt.Recipient.Name,
			},
			Subject:      subject,
			Body:         body,
			TemplateData: make(map[string]any),
		}
		for k, v := range evt.TemplateData {
			notification.TemplateData[k] = v
		}

		startTime := time.Now()
		err := adapter.Send(ctx, notification)
		duration := time.Since(startTime)

		log := &entity.DeliveryLog{
			NotificationID: notification.ID,
			Channel:        channel,
			Attempt:        evt.RetryCount + 1,
			Duration:       duration,
			CreatedAt:      time.Now(),
		}

		if err != nil {
			notification.Status = entity.StatusFailed
			notification.ErrorMessage = err.Error()
			_ = s.repo.Create(ctx, notification)

			log.Status = "failed"
			log.ErrorMessage = err.Error()
			_ = s.deliveryLogRepo.Create(ctx, log)

			lastErr = err
			global.LoggerZap.Error("Failed to send on channel", zap.String("channel", channel), zap.Error(err))
			continue
		}

		// Persist success
		log.Status = "success"
		_ = s.deliveryLogRepo.Create(ctx, log)

		if channel != entity.ChannelInApp {
			now := time.Now()
			notification.Status = entity.StatusSent
			notification.SentAt = &now
			if err := s.repo.Create(ctx, notification); err != nil {
				global.LoggerZap.Error("Failed to store record", zap.Error(err))
			}
		}

		global.LoggerZap.Info("Notification sent",
			zap.String("type", evt.Type),
			zap.String("channel", channel),
			zap.String("recipient", evt.Recipient.Email),
		)
	}

	return lastErr
}

// Find user notifications
func (s *notificationService) Find(ctx context.Context, req *dto.GetNotificationsRequest) (*commonDto.Paginated[*dto.NotificationResponse], error) {
	userID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}

	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	notifications, total, err := s.repo.FindByUserID(ctx, userID, page, pageSize)
	if err != nil {
		return nil, err
	}

	items := notificationMapper.ToNotificationResponseList(notifications)
	return &commonDto.Paginated[*dto.NotificationResponse]{
		Records:    &items,
		Pagination: commonDto.CalculatePagination(page, pageSize, total),
	}, nil
}

// GetUnreadCount
func (s *notificationService) GetUnreadCount(ctx context.Context) (*dto.GetUnreadCountResponse, error) {
	userID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}
	count, err := s.repo.CountUnread(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &dto.GetUnreadCountResponse{Count: count}, nil
}

// MarkAsRead
func (s *notificationService) MarkAsRead(ctx context.Context, id string) error {
	userID, err := getUserID(ctx)
	if err != nil {
		return err
	}

	notification, err := s.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if notification.Recipient.UserID != userID {
		return fmt.Errorf("unauthorized access to notification")
	}

	return s.repo.MarkAsRead(ctx, id)
}

// MarkAllAsRead
func (s *notificationService) MarkAllAsRead(ctx context.Context) error {
	userID, err := getUserID(ctx)
	if err != nil {
		return err
	}
	return s.repo.MarkAllAsRead(ctx, userID)
}

// ProcessDigest sends a summary of aggregated notifications.
func (s *notificationService) ProcessDigest(ctx context.Context, userID, collapseKey string, notificationIDs []string) error {
	if len(notificationIDs) == 0 {
		return nil
	}

	// 1. Fetch first notification to get context (type, etc.)
	firstID := notificationIDs[0]
	n, err := s.repo.Get(ctx, firstID)
	if err != nil {
		return err
	}

	// 2. Prepare summary event
	summaryEvt := &notificationv1.NotificationSendEvent{
		Type:     "digest-summary",
		Channel:  entity.ChannelEmail, // Default to email for digests
		Priority: entity.PriorityNormal,
		Recipient: notificationv1.Recipient{
			UserID: userID,
			Email:  n.Recipient.Email,
			Name:   n.Recipient.Name,
		},
		TemplateData: map[string]string{
			"collapse_key":       collapseKey,
			"notification_count": fmt.Sprintf("%d", len(notificationIDs)),
		},
	}

	// 3. Process as a regular notification
	// NOTE: We recursively call ProcessNotification which is fine here
	return s.ProcessNotification(ctx, summaryEvt)
}
