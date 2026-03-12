package consumer

import (
	"context"
	"encoding/json"
	"time"

	"go-link/notification/global"

	notificationv1 "github.com/huynhanx03/GoLink/events-contract/notification/v1"
	"github.com/huynhanx03/GoLink/events-contract/topics"
	"go.uber.org/zap"
)

// StartRetryConsumer for retry topic
func (nc *notificationConsumer) StartRetryConsumer(ctx context.Context) error {
	handler := func(ctx context.Context, key, value []byte) error {
		var evt notificationv1.NotificationSendEvent
		if err := json.Unmarshal(value, &evt); err != nil {
			global.LoggerZap.Error("Failed to unmarshal retry event", zap.Error(err))
			return nil
		}

		// Exponential Backoff: basic sleep (1m, 4m, 9m...)
		backoff := time.Minute * time.Duration(evt.RetryCount*evt.RetryCount)
		if backoff > 30*time.Minute {
			backoff = 30 * time.Minute
		}

		time.Sleep(backoff)

		return nc.workerPool.Invoke(&evt)
	}

	errHandler := func(err error) {
		global.LoggerZap.Error("RetryConsumer error", zap.Error(err))
	}

	global.LoggerZap.Info("Starting Notification Retry Consumer",
		zap.String("topic", topics.NotificationRetry))

	return nc.consumer.Start(ctx, []string{topics.NotificationRetry}, handler, errHandler)
}
