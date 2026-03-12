package ports

import (
	"context"
	"go-link/notification/internal/core/entity"
)

// DeliveryLogRepository defines the persistence contract for delivery logs.
type DeliveryLogRepository interface {
	Create(ctx context.Context, log *entity.DeliveryLog) error
}
