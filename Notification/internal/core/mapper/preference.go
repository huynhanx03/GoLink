package mapper

import (
	"go-link/notification/internal/core/dto"
	"go-link/notification/internal/core/entity"
)

// ToPreferenceEntity converts an update request DTO to a domain entity.
// Note: UserID implies the identifier is already resolved.
func ToPreferenceEntity(userID string, req *dto.UpdatePreferenceRequest) *entity.UserPreference {
	if req == nil {
		return nil
	}

	return &entity.UserPreference{
		UserID:          userID,
		EmailEnabled:    req.EmailEnabled,
		InAppEnabled:    req.InAppEnabled,
		WebhookEnabled:  req.WebhookEnabled,
		QuietHoursStart: req.QuietHoursStart,
		QuietHoursEnd:   req.QuietHoursEnd,
		Timezone:        req.Timezone,
	}
}

// ToPreferenceResponse converts a domain entity to a response DTO.
func ToPreferenceResponse(e *entity.UserPreference) *dto.PreferenceResponse {
	if e == nil {
		return nil
	}

	return &dto.PreferenceResponse{
		UserID:          e.UserID,
		EmailEnabled:    e.EmailEnabled,
		InAppEnabled:    e.InAppEnabled,
		WebhookEnabled:  e.WebhookEnabled,
		QuietHoursStart: e.QuietHoursStart,
		QuietHoursEnd:   e.QuietHoursEnd,
		Timezone:        e.Timezone,
	}
}
