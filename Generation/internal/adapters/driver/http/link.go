package http

import (
	"context"

	"go-link/common/pkg/common/http/handler"
	"go-link/generation/internal/core/dto"
	"go-link/generation/internal/ports"
)

type LinkHandler interface {
	Create(ctx context.Context, req *dto.CreateLinkRequest) (*dto.LinkResponse, error)
	Delete(ctx context.Context, req *dto.DeleteLinkRequest) (*dto.LinkResponse, error)
}

type linkHandler struct {
	handler.BaseHandler
	linkService ports.LinkService
}

func NewLinkHandler(linkService ports.LinkService) LinkHandler {
	return &linkHandler{
		linkService: linkService,
	}
}

// Create creates a new short link
func (h *linkHandler) Create(ctx context.Context, req *dto.CreateLinkRequest) (*dto.LinkResponse, error) {
	return h.linkService.Create(ctx, req)
}

// Delete deletes a short link
func (h *linkHandler) Delete(ctx context.Context, req *dto.DeleteLinkRequest) (*dto.LinkResponse, error) {
	return nil, h.linkService.Delete(ctx, req)
}
