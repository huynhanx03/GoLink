package worker

import (
	"context"
	"strings"
	"time"

	"go-link/notification/global"
	"go-link/notification/internal/ports"

	"go.uber.org/zap"
)

type digestWorker struct {
	notificationService ports.NotificationService
	digestService       ports.DigestService
	interval            time.Duration
	stopChan            chan struct{}
}

// NewDigestWorker creates a new worker to process digests.
func NewDigestWorker(
	notificationService ports.NotificationService,
	digestService ports.DigestService,
) ports.DigestWorker {
	return &digestWorker{
		notificationService: notificationService,
		digestService:       digestService,
		interval:            1 * time.Minute, // Default interval
		stopChan:            make(chan struct{}),
	}
}

func (w *digestWorker) Start(ctx context.Context) error {
	global.LoggerZap.Info("Starting digest worker", zap.Duration("interval", w.interval))

	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			w.run(ctx)
		case <-w.stopChan:
			global.LoggerZap.Info("Digest worker stopped")
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (w *digestWorker) Stop() {
	close(w.stopChan)
}

func (w *digestWorker) run(ctx context.Context) {
	keys, err := w.digestService.ScanPendingDigests(ctx)
	if err != nil {
		global.LoggerZap.Error("Failed to scan pending digests", zap.Error(err))
		return
	}

	for _, key := range keys {
		// Key format: notification:digest:userID:collapseKey
		parts := strings.Split(key, ":")
		if len(parts) < 4 {
			continue
		}
		userID := parts[2]
		collapseKey := parts[3]

		notificationIDs, err := w.digestService.ConsumeDigest(ctx, key)
		if err != nil {
			global.LoggerZap.Error("Failed to consume digest", zap.String("key", key), zap.Error(err))
			continue
		}

		if len(notificationIDs) > 0 {
			err = w.notificationService.ProcessDigest(ctx, userID, collapseKey, notificationIDs)
			if err != nil {
				global.LoggerZap.Error("Failed to process digest",
					zap.String("userID", userID),
					zap.String("collapseKey", collapseKey),
					zap.Error(err),
				)
			}
		}
	}
}
