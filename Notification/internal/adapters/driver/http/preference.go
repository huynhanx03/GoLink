package http

import (
	"context"

	"go-link/common/pkg/common/http/handler"
	"go-link/notification/internal/core/dto"
	"go-link/notification/internal/ports"
)

// PreferenceHandler defines the HTTP handler interface for user preferences.
type PreferenceHandler interface {
	Get(ctx context.Context, _ *struct{}) (*dto.PreferenceResponse, error)
	Update(ctx context.Context, req *dto.UpdatePreferenceRequest) (*dto.PreferenceResponse, error)
}

type preferenceHandler struct {
	handler.BaseHandler
	service ports.PreferenceService
}

// NewPreferenceHandler creates a new PreferenceHandler.
func NewPreferenceHandler(service ports.PreferenceService) PreferenceHandler {
	return &preferenceHandler{service: service}
}

// Get retrieves the user's notification preferences.
func (h *preferenceHandler) Get(ctx context.Context, _ *struct{}) (*dto.PreferenceResponse, error) {
	return h.service.Get(ctx)
}

// Update modifies the user's notification preferences.
func (h *preferenceHandler) Update(ctx context.Context, req *dto.UpdatePreferenceRequest) (*dto.PreferenceResponse, error) {
	return h.service.Update(ctx, req)
}
