package consumer

import (
	"context"
	"encoding/json"
	"time"

	"go-link/common/pkg/common/workerpool"
	"go-link/common/pkg/mq/kafka"
	"go-link/common/pkg/utils"
	"go-link/notification/global"
	"go-link/notification/internal/constant"
	"go-link/notification/internal/ports"

	notificationv1 "github.com/huynhanx03/GoLink/events-contract/notification/v1"
	"github.com/huynhanx03/GoLink/events-contract/topics"
	"go.uber.org/zap"
)

const (
	poolSize = 10
)

// notificationConsumer processes events in batches
type notificationConsumer struct {
	consumer            kafka.ConsumerGroup
	notificationService ports.NotificationService
	workerPool          *workerpool.GenericPool[*notificationv1.NotificationSendEvent]
	producer            kafka.SyncProducer
}

// NewNotificationConsumer creates consumer
func NewNotificationConsumer(
	cfg *kafka.Config,
	svc ports.NotificationService,
	producer kafka.SyncProducer,
) (ports.NotificationConsumer, error) {
	c, err := kafka.NewConsumer(cfg, constant.ConsumerGroupNotification, kafka.Recovery)
	if err != nil {
		return nil, err
	}

	nc := &notificationConsumer{
		consumer:            c,
		notificationService: svc,
		producer:            producer,
	}

	taskFunc := func(evt *notificationv1.NotificationSendEvent) {
		if err := svc.ProcessNotification(context.Background(), evt); err != nil {
			nc.handleFailure(context.Background(), evt, err)
		}
	}

	pool, err := workerpool.NewGenericPool(poolSize, taskFunc)
	if err != nil {
		return nil, err
	}
	nc.workerPool = pool

	return nc, nil
}

// Start begins consuming
func (nc *notificationConsumer) Start(ctx context.Context) error {
	batchSize := global.Config.Kafka.ConsumerBatchSize
	if batchSize <= 0 {
		batchSize = 100
	}

	batchInterval := utils.ToDurationMs(global.Config.Kafka.ConsumerBatchInterval)
	if batchInterval <= 0 {
		batchInterval = 500 * time.Millisecond
	}

	batchChan := make(chan *notificationv1.NotificationSendEvent, batchSize*2)

	go nc.processBatchLoop(ctx, batchChan, batchSize, batchInterval)

	handler := func(ctx context.Context, key, value []byte) error {
		var evt notificationv1.NotificationSendEvent
		if err := json.Unmarshal(value, &evt); err != nil {
			global.LoggerZap.Error("Failed to unmarshal notification event", zap.Error(err))
			return nil
		}

		select {
		case batchChan <- &evt:
		case <-ctx.Done():
			return ctx.Err()
		}

		return nil
	}

	errHandler := func(err error) {
		global.LoggerZap.Error("NotificationConsumer error", zap.Error(err))
	}

	global.LoggerZap.Info("Starting Notification Consumer (Batch Mode)",
		zap.String("topic", topics.NotificationSend))

	return nc.consumer.Start(ctx, []string{topics.NotificationSend}, handler, errHandler)
}

// Stop closes consumer
func (nc *notificationConsumer) Stop() error {
	return nc.consumer.Close()
}

func (nc *notificationConsumer) processBatchLoop(
	ctx context.Context,
	batchChan <-chan *notificationv1.NotificationSendEvent,
	batchSize int,
	batchInterval time.Duration,
) {
	ticker := time.NewTicker(batchInterval)
	defer ticker.Stop()

	for {
		select {
		case evt := <-batchChan:
			if err := nc.workerPool.Invoke(evt); err != nil {
				global.LoggerZap.Error("Failed to submit notification to worker pool", zap.Error(err))
			}
		case <-ticker.C:
		case <-ctx.Done():
			return
		}
	}
}

func (nc *notificationConsumer) handleFailure(ctx context.Context, evt *notificationv1.NotificationSendEvent, err error) {
	evt.RetryCount++
	evt.LastError = err.Error()

	global.LoggerZap.Error("Failed to process notification",
		zap.Error(err),
		zap.String("type", evt.Type),
		zap.Int("retry_count", evt.RetryCount),
	)

	payload, _ := json.Marshal(evt)
	topic := topics.NotificationRetry
	if evt.RetryCount >= 3 {
		topic = topics.NotificationDLQ
		global.LoggerZap.Warn("Notification moved to DLQ", zap.String("idempotency_key", evt.IdempotencyKey))
	}

	if _, _, err := nc.producer.Publish(ctx, topic, []byte(evt.IdempotencyKey), payload); err != nil {
		global.LoggerZap.Error("Failed to republish notification for retry/dlq", zap.Error(err))
	}
}
