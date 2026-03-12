package service

import (
	"context"
	"time"

	"go-link/common/pkg/common/apperr"
	"go-link/common/pkg/common/http/response"
	"go-link/common/pkg/constraints"
	"go-link/notification/internal/core/dto"
	"go-link/notification/internal/core/entity"
	"go-link/notification/internal/core/mapper"
	"go-link/notification/internal/ports"
)

func getUserID(ctx context.Context) (string, error) {
	userID, ok := ctx.Value(constraints.ContextKeyUserID).(string)
	if !ok || userID == "" {
		return "", apperr.New(response.CodeUnauthorized, "unauthorized", 0, nil)
	}
	return userID, nil
}

// PreferenceService enforces user-level channel opt-in and quiet hours rules.
type PreferenceService struct {
	repo ports.UserPreferenceRepository
}

// NewPreferenceService creates a new PreferenceService backed by the given repository.
func NewPreferenceService(repo ports.UserPreferenceRepository) ports.PreferenceService {
	return &PreferenceService{repo: repo}
}

// IsChannelAllowed returns true when the user has opted into the given channel
// and is not currently within a configured quiet hours window.
// Defaults to true (allow) on repository error to avoid blocking notifications.
func (s *PreferenceService) IsChannelAllowed(ctx context.Context, channel string) (bool, error) {
	userID, err := getUserID(ctx)
	if err != nil {
		return true, nil // fail open if no user
	}

	pref, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return true, nil // fail open
	}

	// Check per-channel opt-in.
	switch channel {
	case entity.ChannelEmail:
		if !pref.EmailEnabled {
			return false, nil
		}
	case entity.ChannelInApp:
		if !pref.InAppEnabled {
			return false, nil
		}
	case entity.ChannelWebhook:
		if !pref.WebhookEnabled {
			return false, nil
		}
	}

	// Check quiet hours (both boundaries must be set).
	if pref.QuietHoursStart != nil && pref.QuietHoursEnd != nil {
		now := time.Now()
		if pref.Timezone != "" {
			if loc, err := time.LoadLocation(pref.Timezone); err == nil {
				now = now.In(loc)
			}
		}
		hour := now.Hour()
		start, end := *pref.QuietHoursStart, *pref.QuietHoursEnd
		if start <= end {
			// Same-day window, e.g. 09:00–18:00.
			if hour >= start && hour < end {
				return false, nil
			}
		} else {
			// Overnight window, e.g. 22:00–06:00.
			if hour >= start || hour < end {
				return false, nil
			}
		}
	}

	return true, nil
}

// Get retrieves a user's notification preferences and returns as DTO.
func (s *PreferenceService) Get(ctx context.Context) (*dto.PreferenceResponse, error) {
	userID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}
	pref, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return mapper.ToPreferenceResponse(pref), nil
}

// Update updates a user's notification preferences from DTO.
func (s *PreferenceService) Update(ctx context.Context, req *dto.UpdatePreferenceRequest) (*dto.PreferenceResponse, error) {
	userID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}

	pref := mapper.ToPreferenceEntity(userID, req)
	if err := s.repo.Upsert(ctx, pref); err != nil {
		return nil, err
	}

	return mapper.ToPreferenceResponse(pref), nil
}
