package repository

import (
	"context"

	"go-link/common/pkg/database/widecolumn"
	"go-link/generation/global"
	"go-link/generation/internal/core/entity"
	"go-link/generation/internal/ports"
)

const (
	defaultTTL = 2592000 // 30 days
)

type LinkRepository struct {
	repo *widecolumn.BaseRepository[entity.Link]
}

// NewLinkRepository creates a new instance of LinkRepository
func NewLinkRepository() ports.LinkRepository {
	return &LinkRepository{
		repo: widecolumn.NewBaseRepository(global.WideColumnClient.GetSession(), entity.Link{}),
	}
}

// Create a link with TTL
func (r *LinkRepository) Create(ctx context.Context, link *entity.Link) error {
	return r.repo.CreateWithTTL(ctx, link, defaultTTL)
}
