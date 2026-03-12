package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"go-link/common/pkg/database/mongodb"
	"go-link/notification/internal/adapters/driven/db/mapper"
	"go-link/notification/internal/adapters/driven/db/models"
	"go-link/notification/internal/core/entity"
	"go-link/notification/internal/ports"
)

const (
	DeliveryLogCollection = "delivery_logs"
)

type deliveryLogRepo struct {
	repo *mongodb.BaseRepository[models.DeliveryLog]
}

// NewDeliveryLogRepo creates a MongoDB-backed DeliveryLogRepository.
func NewDeliveryLogRepo(db *mongo.Database) ports.DeliveryLogRepository {
	collection := db.Collection(DeliveryLogCollection)
	return &deliveryLogRepo{
		repo: mongodb.NewBaseRepository[models.DeliveryLog](collection),
	}
}

// Create inserts a new delivery log record using the generic base repository.
func (r *deliveryLogRepo) Create(ctx context.Context, log *entity.DeliveryLog) error {
	log.CreatedAt = time.Now()
	model := mapper.ToDeliveryLogModel(log)

	err := r.repo.Create(ctx, model)
	if err == nil {
		log.ID = model.ID.Hex()
	}
	return err
}
