package di

import (
	"go-link/common/pkg/mq/kafka"

	"go.uber.org/zap"

	"go-link/redirection/global"
	"go-link/redirection/internal/adapters/driven/cache"
	db "go-link/redirection/internal/adapters/driven/db"
	linkconsumer "go-link/redirection/internal/adapters/driver/consumer/link"
	driverHttp "go-link/redirection/internal/adapters/driver/http"
	"go-link/redirection/internal/core/service"
	"go-link/redirection/internal/ports"
)

type LinkContainer struct {
	Repository ports.LinkRepository
	Service    ports.LinkService
	Consumer   ports.LinkConsumer
	Handler    driverHttp.LinkHandler
}

func InitLinkDependencies() *LinkContainer {
	// Cache
	cache := cache.NewLink(global.Redis)

	// Repository
	repository := db.NewLinkRepository()

	// Service
	service := service.NewLinkService(repository, cache)

	// Handler
	handler := driverHttp.NewLinkHandler(service)

	// Consumer
	kafkaCfg := &kafka.Config{
		Brokers:  global.Config.Kafka.Brokers,
		ClientID: "link-cdc",
		ProducerInfo: kafka.ProducerConfig{
			FlushFrequency:  global.Config.Kafka.FlushFrequency,
			FlushBytes:      global.Config.Kafka.FlushBytes,
			MaxMessageBytes: global.Config.Kafka.MaxMessageBytes,
			MaxRetries:      global.Config.Kafka.MaxRetries,
			RetryBackoff:    global.Config.Kafka.RetryBackoff,
			ReturnSuccesses: true,
		},
		ConsumerInfo: kafka.ConsumerConfig{
			SessionTimeout:    global.Config.Kafka.Timeout * 1000,
			MaxProcessingTime: global.Config.Kafka.MaxProcessingTime,
		},
	}

	consumer, err := linkconsumer.NewCDCConsumer(kafkaCfg, service)
	if err != nil {
		global.LoggerZap.Fatal("failed to create link cdc consumer", zap.Error(err))
	}

	return &LinkContainer{
		Repository: repository,
		Service:    service,
		Consumer:   consumer,
		Handler:    handler,
	}
}
