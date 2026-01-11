package link

import (
	"context"
	"time"

	"go-link/common/pkg/cdc"
	"go-link/common/pkg/common/workerpool"
	"go-link/common/pkg/mq/kafka"
	"go-link/common/pkg/utils"

	"go.uber.org/zap"

	"go-link/redirection/global"
	"go-link/redirection/internal/constant"
	"go-link/redirection/internal/core/entity"
	"go-link/redirection/internal/ports"
)

const (
	poolSize = 10
)

type CDCConsumer struct {
	consumer    kafka.ConsumerGroup
	linkService ports.LinkService
	workerPool  *workerpool.GenericPool[[]*cdc.DebeziumPayload[entity.Link]]
}

func NewCDCConsumer(cfg *kafka.Config, linkService ports.LinkService) (ports.LinkConsumer, error) {
	c, err := kafka.NewConsumer(cfg, constant.ConsumerGroupLinkCDC, kafka.Recovery)
	if err != nil {
		return nil, err
	}

	taskFunc := func(batch []*cdc.DebeziumPayload[entity.Link]) {
		if err := linkService.HandleLinkBatchChange(context.Background(), batch); err != nil {
			global.Logger.Error("Failed to process link batch", zap.Error(err))
		}
	}

	pool, err := workerpool.NewGenericPool(poolSize, taskFunc)
	if err != nil {
		return nil, err
	}

	return &CDCConsumer{
		consumer:    c,
		linkService: linkService,
		workerPool:  pool,
	}, nil
}

// Start starts the consumer
func (c *CDCConsumer) Start(ctx context.Context) error {
	batchSize := global.Config.Kafka.ConsumerBatchSize
	if batchSize <= 0 {
		batchSize = 100
	}

	batchInterval := utils.ToDurationMs(global.Config.Kafka.ConsumerBatchInterval)
	if batchInterval <= 0 {
		batchInterval = 500 * time.Millisecond
	}

	batchChan := make(chan *cdc.DebeziumPayload[entity.Link], batchSize*2)

	go c.processBatchLoop(ctx, batchChan, batchSize, batchInterval)

	handler := func(ctx context.Context, key, value []byte) error {
		_, payload, err := c.extractPayload(key, value)
		if err != nil {
			global.Logger.Warn("Skipping malformed CDC message", zap.Error(err))
			return nil
		}

		if payload == nil {
			return nil
		}

		select {
		case batchChan <- payload:
		case <-ctx.Done():
			return ctx.Err()
		}

		return nil
	}

	errHandler := func(err error) {
		global.Logger.Error("LinkCDCConsumer error", zap.Error(err))
	}

	global.Logger.Info("Starting Link CDC Consumer (Batch Mode)", zap.String("topic", constant.TopicLinkCDC))
	return c.consumer.Start(ctx, []string{constant.TopicLinkCDC}, handler, errHandler)
}

func (c *CDCConsumer) Stop() error {
	return c.consumer.Close()
}

func (c *CDCConsumer) processBatchLoop(
	ctx context.Context,
	batchChan <-chan *cdc.DebeziumPayload[entity.Link],
	batchSize int,
	batchInterval time.Duration,
) {
	batch := make([]*cdc.DebeziumPayload[entity.Link], 0, batchSize)
	ticker := time.NewTicker(batchInterval)
	defer ticker.Stop()

	flush := func() {
		size := len(batch)
		if size == 0 {
			return
		}

		finalBatch := make([]*cdc.DebeziumPayload[entity.Link], size)
		copy(finalBatch, batch)

		if err := c.workerPool.Invoke(finalBatch); err != nil {
			global.Logger.Error("Failed to submit batch to worker pool", zap.Error(err))
		}

		batch = batch[:0]
	}

	for {
		select {
		case msg := <-batchChan:
			batch = append(batch, msg)
			if len(batch) >= batchSize {
				flush()
			}
		case <-ticker.C:
			flush()
		case <-ctx.Done():
			flush()
			return
		}
	}
}

func (c *CDCConsumer) extractPayload(key, value []byte) (string, *cdc.DebeziumPayload[entity.Link], error) {
	if len(value) == 0 {
		return "", nil, nil
	}

	cdcPayload, err := cdc.ParseDebeziumMessage[CDCLink](value)
	if err != nil {
		return "", nil, err
	}

	payload := &cdc.DebeziumPayload[entity.Link]{
		Source: cdcPayload.Source,
		Op:     cdcPayload.Op,
		TsMs:   cdcPayload.TsMs,
	}

	if cdcPayload.After != nil {
		payload.After = cdcPayload.After.ToEntity()
	}

	if cdcPayload.Before != nil {
		payload.Before = cdcPayload.Before.ToEntity()
	}

	// For Delete, if Before is present, we have ID.
	var id string
	if payload.Before != nil {
		id = payload.Before.ID
	} else if payload.After != nil {
		id = payload.After.ID
	}

	return id, payload, nil
}
