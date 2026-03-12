package di

import (
	httpAdapter "go-link/notification/internal/adapters/driver/http"
	"go-link/notification/internal/core/service"
)

// SSEContainer holds all sse-domain dependencies.
type SSEContainer struct {
	Hub     *service.SSEHub
	Handler *httpAdapter.SSEHandler
}

// InitSSEDependencies initializes dependencies for the SSE component.
func InitSSEDependencies() SSEContainer {
	hub := service.NewSSEHub()
	handler := httpAdapter.NewSSEHandler(hub)

	return SSEContainer{
		Hub:     hub,
		Handler: handler,
	}
}
