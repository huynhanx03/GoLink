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
	NotificationCollection = "notifications"
)

type notificationMongoRepo struct {
	repo *mongodb.BaseRepository[models.Notification]
}

// NewNotificationMongoRepo creates a MongoDB-backed NotificationRepository.
func NewNotificationMongoRepo(db *mongo.Database) ports.NotificationRepository {
	collection := db.Collection(NotificationCollection)
	return &notificationMongoRepo{
		repo: mongodb.NewBaseRepository[models.Notification](collection),
	}
}

// Create inserts a new notification record.
func (r *notificationMongoRepo) Create(ctx context.Context, notification *entity.Notification) error {
	now := time.Now()

	// Set TTL expiry to 90 days from creation for automatic cleanup.
	if notification.ExpiresAt.IsZero() {
		notification.ExpiresAt = now.Add(90 * 24 * time.Hour)
	}

	model := mapper.ToNotificationModel(notification)

	err := r.repo.Create(ctx, model)
	if err == nil {
		notification.ID = model.ID.Hex()
	}
	return err
}

// Get retrieves a notification by its hex ID string.
func (r *notificationMongoRepo) Get(ctx context.Context, id string) (*entity.Notification, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	model, err := r.repo.Get(ctx, oid)
	if err != nil {
		return nil, err
	}

	return mapper.ToNotificationEntity(model), nil
}

// FindByUserID returns a paginated list of notifications for a given user, newest first.
func (r *notificationMongoRepo) FindByUserID(ctx context.Context, userID string, page, pageSize int) ([]*entity.Notification, int64, error) {
	filter := bson.M{"recipient.user_id": userID}

	total, err := r.repo.GetCollection().CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}

	skip := int64((page - 1) * pageSize)
	opts := options.Find().
		SetSort(bson.D{{Key: "created_at", Value: -1}}).
		SetSkip(skip).
		SetLimit(int64(pageSize))

	cursor, err := r.repo.GetCollection().Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var modelsList []*models.Notification
	if err := cursor.All(ctx, &modelsList); err != nil {
		return nil, 0, err
	}

	notifications := make([]*entity.Notification, 0, len(modelsList))
	for _, m := range modelsList {
		notifications = append(notifications, mapper.ToNotificationEntity(m))
	}

	return notifications, total, nil
}

// MarkAsRead sets is_read=true on a single notification by its hex ID.
func (r *notificationMongoRepo) MarkAsRead(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"is_read":    true,
			"updated_at": time.Now(),
		},
	}

	_, err = r.repo.GetCollection().UpdateOne(ctx, bson.M{"_id": oid}, update)
	return err
}

// MarkAllAsRead sets is_read=true on all unread notifications for a user.
func (r *notificationMongoRepo) MarkAllAsRead(ctx context.Context, userID string) error {
	filter := bson.M{
		"recipient.user_id": userID,
		"is_read":           false,
	}
	update := bson.M{
		"$set": bson.M{
			"is_read":    true,
			"updated_at": time.Now(),
		},
	}

	_, err := r.repo.GetCollection().UpdateMany(ctx, filter, update)
	return err
}

// CountUnread returns the count of unread notifications for a user.
func (r *notificationMongoRepo) CountUnread(ctx context.Context, userID string) (int64, error) {
	filter := bson.M{
		"recipient.user_id": userID,
		"is_read":           false,
	}
	return r.repo.GetCollection().CountDocuments(ctx, filter)
}
