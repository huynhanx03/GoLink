package di

import (
	"context"
	"go-link/common/pkg/mq/kafka"
	"go-link/notification/global"
	notificationCache "go-link/notification/internal/adapters/driven/cache"
	channelAdapter "go-link/notification/internal/adapters/driven/channel"
	dbAdapter "go-link/notification/internal/adapters/driven/db"
	consumerAdapter "go-link/notification/internal/adapters/driver/consumer"
	httpAdapter "go-link/notification/internal/adapters/driver/http"
	workerAdapter "go-link/notification/internal/adapters/driver/worker"
	"go-link/notification/internal/core/entity"
	"go-link/notification/internal/core/service"
	"go-link/notification/internal/ports"

	"go.uber.org/zap"
)

// SetupDependencies wires all concrete implementations into the dependency container.
func SetupDependencies(ctx context.Context) *Container {
	// Root components / Repos
	notificationRepo := dbAdapter.NewNotificationMongoRepo(global.MongoDB.DB)
	webhookConfigRepo := dbAdapter.NewWebhookConfigRepo(global.MongoDB.DB)
	userPreferenceRepo := dbAdapter.NewUserPreferenceRepo(global.MongoDB.DB)
	deliveryLogRepo := dbAdapter.NewDeliveryLogRepo(global.MongoDB.DB)

	// Cache / Rate limit
	idempotencyChecker := notificationCache.NewRedisIdempotencyChecker(global.Redis)
	unreadCounter := notificationCache.NewUnreadCounter(global.Redis)
	rateLimiter := service.NewRateLimiter(global.Redis)
	digestService := service.NewDigestService(global.Redis)

	// Sub-domains
	sseContainer := InitSSEDependencies()
	preferenceContainer := InitPreferenceDependencies(userPreferenceRepo)
	webhookContainer := InitWebhookDependencies(webhookConfigRepo)

	// Channel adapters
	inAppAdapter := channelAdapter.NewInAppAdapter(notificationRepo, sseContainer.Hub, unreadCounter)
	emailAdapter := channelAdapter.NewResendEmailAdapter(
		global.Config.Resend.APIKey,
		global.Config.Resend.FromEmail,
		global.Config.Resend.FromName,
	)
	webhookAdapter := channelAdapter.NewWebhookAdapter(webhookConfigRepo)
	fcmAdapter, err := channelAdapter.NewFCMAdapter(ctx)
	if err != nil {
		global.LoggerZap.Warn("Failed to initialize FCM adapter", zap.Error(err))
	}

	channels := map[string]ports.ChannelAdapter{
		entity.ChannelEmail:   emailAdapter,
		entity.ChannelInApp:   inAppAdapter,
		entity.ChannelWebhook: webhookAdapter,
		entity.ChannelFCM:     fcmAdapter,
	}

	// Notification core Service
	notificationService := service.NewNotificationService(
		notificationRepo,
		idempotencyChecker,
		channels,
		preferenceContainer.Service,
		rateLimiter,
		deliveryLogRepo,
		digestService,
	)

	notificationHandler := httpAdapter.NewNotificationHandler(notificationService)

	// Kafka consumer
	kafkaCfg := &kafka.Config{
		Brokers:  global.Config.Kafka.Brokers,
		ClientID: "notification-service",
		ProducerInfo: kafka.ProducerConfig{
			MaxRetries:   global.Config.Kafka.MaxRetries,
			RetryBackoff: global.Config.Kafka.RetryBackoff,
		},
	}

	syncProducer, err := kafka.NewSyncProducer(kafkaCfg)
	if err != nil {
		global.LoggerZap.Fatal("Failed to create kafka sync producer", zap.Error(err))
	}

	digestWorker := workerAdapter.NewDigestWorker(notificationService, digestService)

	notificationConsumer, err := consumerAdapter.NewNotificationConsumer(kafkaCfg, notificationService, syncProducer)
	if err != nil {
		global.LoggerZap.Fatal("Failed to create notification consumer", zap.Error(err))
	}

	notificationContainer := NotificationContainer{
		Service:       notificationService,
		Repo:          notificationRepo,
		Consumer:      notificationConsumer,
		Handler:       notificationHandler,
		UnreadCounter: unreadCounter,
		DigestService: digestService,
		DigestWorker:  digestWorker,
		Producer:      syncProducer,
	}

	container := &Container{
		NotificationContainer: notificationContainer,
		PreferenceContainer:   preferenceContainer,
		WebhookContainer:      webhookContainer,
		SSEContainer:          sseContainer,
	}

	GlobalContainer = container
	return container
}
