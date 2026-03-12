package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go-link/common/pkg/database/mongodb"
	"go-link/notification/internal/adapters/driven/db/mapper"
	"go-link/notification/internal/adapters/driven/db/models"
	"go-link/notification/internal/core/entity"
	"go-link/notification/internal/ports"
)

const (
	UserPreferenceCollection = "user_preferences"
)

type userPreferenceRepo struct {
	repo *mongodb.BaseRepository[models.UserPreference]
}

// NewUserPreferenceRepo creates a MongoDB-backed UserPreferenceRepository.
func NewUserPreferenceRepo(db *mongo.Database) ports.UserPreferenceRepository {
	collection := db.Collection(UserPreferenceCollection)
	return &userPreferenceRepo{
		repo: mongodb.NewBaseRepository[models.UserPreference](collection),
	}
}

// GetByUserID retrieves the preference for a user, returning the default if not found.
func (r *userPreferenceRepo) GetByUserID(ctx context.Context, userID string) (*entity.UserPreference, error) {
	var prefModel models.UserPreference
	err := r.repo.GetCollection().FindOne(ctx, bson.M{"user_id": userID}).Decode(&prefModel)
	if err == mongo.ErrNoDocuments {
		return entity.DefaultPreference(userID), nil
	}
	if err != nil {
		return nil, err
	}
	return mapper.ToUserPreferenceEntity(&prefModel), nil
}

// Upsert inserts or updates the user preference document using user_id as the key.
func (r *userPreferenceRepo) Upsert(ctx context.Context, pref *entity.UserPreference) error {
	now := time.Now()

	model := mapper.ToUserPreferenceModel(pref)

	filter := bson.M{"user_id": model.UserID}
	update := bson.M{
		"$set": bson.M{
			"email_enabled":     model.EmailEnabled,
			"in_app_enabled":    model.InAppEnabled,
			"webhook_enabled":   model.WebhookEnabled,
			"quiet_hours_start": model.QuietHoursStart,
			"quiet_hours_end":   model.QuietHoursEnd,
			"updated_at":        now,
		},
		"$setOnInsert": bson.M{
			"_id":        primitive.NewObjectID(),
			"user_id":    model.UserID,
			"created_at": now,
		},
	}

	opts := options.Update().SetUpsert(true)
	_, err := r.repo.GetCollection().UpdateOne(ctx, filter, update, opts)
	return err
}
