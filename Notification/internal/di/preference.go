package di

import (
	httpAdapter "go-link/notification/internal/adapters/driver/http"
	"go-link/notification/internal/core/service"
	"go-link/notification/internal/ports"
)

// PreferenceContainer holds all preference-domain dependencies.
type PreferenceContainer struct {
	Service ports.PreferenceService
	Handler httpAdapter.PreferenceHandler
}

// InitPreferenceDependencies initializes dependencies for the Preference domain.
func InitPreferenceDependencies(
	repo ports.UserPreferenceRepository,
) PreferenceContainer {
	svc := service.NewPreferenceService(repo)
	handler := httpAdapter.NewPreferenceHandler(svc)

	return PreferenceContainer{
		Service: svc,
		Handler: handler,
	}
}
