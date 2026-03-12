package ports

import (
	"context"

	"go-link/notification/internal/core/entity"
)

// ChannelAdapter defines the contract for notification delivery channels.
type ChannelAdapter interface {
	Send(ctx context.Context, notification *entity.Notification) error
	Type() string
}
