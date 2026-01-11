package ports

import (
	"context"
	"go-link/generation/internal/core/dto"
	"go-link/generation/internal/core/entity"
)

type LinkRepository interface {
	Create(ctx context.Context, link *entity.Link) error
}

type LinkCacheRepository interface {
	Set(ctx context.Context, link *entity.Link) error
}

type LinkService interface {
	Create(ctx context.Context, req *dto.CreateLinkRequest) (*dto.LinkResponse, error)
}
