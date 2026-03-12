package channel

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"go.uber.org/zap"
	"google.golang.org/api/option"

	"go-link/notification/global"
	"go-link/notification/internal/core/entity"
)

type fcmAdapter struct {
	client *messaging.Client
}

// NewFCMAdapter creates a new FCM adapter.
func NewFCMAdapter(ctx context.Context) (*fcmAdapter, error) {
	// If no config, return nil (don't fail the whole app startup)
	if global.Config.FCM.ServiceAccountPath == "" {
		return &fcmAdapter{}, nil
	}

	opt := option.WithServiceAccountFile(global.Config.FCM.ServiceAccountPath)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase app: %w", err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting messaging client: %w", err)
	}

	return &fcmAdapter{client: client}, nil
}

// Type returns the channel type.
func (a *fcmAdapter) Type() string {
	return entity.ChannelFCM
}

// Send push notification via FCM.
func (a *fcmAdapter) Send(ctx context.Context, n *entity.Notification) error {
	if a.client == nil {
		return fmt.Errorf("fcm client not initialized")
	}

	// device_token expected in TemplateData
	token, ok := n.TemplateData["device_token"].(string)
	if !ok || token == "" {
		return fmt.Errorf("device_token missing in template data")
	}

	message := &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title: n.Subject,
			Body:  n.Body,
		},
		Data: map[string]string{
			"type":            n.Type,
			"idempotency_key": n.IdempotencyKey,
		},
	}

	_, err := a.client.Send(ctx, message)
	if err != nil {
		global.LoggerZap.Error("FCM send failed", zap.Error(err), zap.String("id", n.ID))
		return err
	}

	return nil
}
