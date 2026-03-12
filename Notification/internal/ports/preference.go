package ports

import (
	"context"

	"go-link/notification/internal/core/dto"
	"go-link/notification/internal/core/entity"
)

// UserPreferenceRepository defines persistence operations for user notification preferences.
type UserPreferenceRepository interface {
	GetByUserID(ctx context.Context, userID string) (*entity.UserPreference, error)
	Upsert(ctx context.Context, pref *entity.UserPreference) error
}

// PreferenceService defines the business logic contract for preferences.
type PreferenceService interface {
	IsChannelAllowed(ctx context.Context, channel string) (bool, error)
	Get(ctx context.Context) (*dto.PreferenceResponse, error)
	Update(ctx context.Context, req *dto.UpdatePreferenceRequest) (*dto.PreferenceResponse, error)
}
