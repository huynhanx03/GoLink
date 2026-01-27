package repository

import (
	"context"

	"go-link/common/pkg/database/widecolumn"
	"go-link/generation/global"
	"go-link/generation/internal/adapters/driven/db/models"
	"go-link/generation/internal/core/entity"
	"go-link/generation/internal/ports"
)

const (
	defaultTTL = 2592000 // 30 days
)

type LinkRepository struct {
	repo *widecolumn.BaseRepository[models.Link]
}

// NewLinkRepository creates a new instance of LinkRepository
func NewLinkRepository() ports.LinkRepository {
	return &LinkRepository{
		repo: widecolumn.NewBaseRepository(global.WideColumnClient.GetSession(), models.Link{}),
	}
}

// Create a link with TTL
func (l *LinkRepository) Create(ctx context.Context, link *entity.Link, ttl int) error {
	if ttl == 0 {
		ttl = defaultTTL
	}
	return l.repo.CreateWithTTL(ctx, models.FromEntity(link), ttl)
}

func (l *LinkRepository) Get(ctx context.Context, id string) (*entity.Link, error) {
	link, err := l.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return link.ToEntity(), nil
}

func (l *LinkRepository) Delete(ctx context.Context, id string) error {
	_, err := l.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	return l.repo.Delete(ctx, id)
}
