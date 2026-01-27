package http

import (
	"context"

	"go-link/orchestrator/internal/core/dto"
	"go-link/orchestrator/internal/ports"

	"go-link/common/pkg/common/http/handler"
)

// OrchestratorHandler defines the Orchestrator HTTP handler interface.
type OrchestratorHandler interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error)
}

type orchestratorHandler struct {
	handler.BaseHandler
	service ports.OrchestratorService
}

// NewOrchestratorHandler creates a new OrchestratorHandler instance.
func NewOrchestratorHandler(service ports.OrchestratorService) OrchestratorHandler {
	return &orchestratorHandler{
		service: service,
	}
}

// Register handles user registration.
func (h *orchestratorHandler) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	return h.service.RegisterUser(ctx, req)
}
