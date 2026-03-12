package di

import (
	"go-link/common/pkg/mq/kafka"
	httpAdapter "go-link/notification/internal/adapters/driver/http"
	"go-link/notification/internal/ports"
)

type NotificationContainer struct {
	Service       ports.NotificationService
	Repo          ports.NotificationRepository
	Consumer      ports.NotificationConsumer
	Handler       httpAdapter.NotificationHandler
	UnreadCounter ports.UnreadCounter
	DigestService ports.DigestService
	DigestWorker  ports.DigestWorker
	Producer      kafka.SyncProducer
}
