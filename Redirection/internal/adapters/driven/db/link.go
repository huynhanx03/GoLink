package repository

import (
	"context"

	"go-link/common/pkg/database/widecolumn"

	"go-link/redirection/global"
	"go-link/redirection/internal/adapters/driven/db/models"
	"go-link/redirection/internal/core/entity"
	"go-link/redirection/internal/ports"
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

// GetOriginalURL retrieves the original URL for a given short code
func (l *LinkRepository) GetOriginalURL(ctx context.Context, shortCode string) (*entity.Link, error) {
	model, err := l.repo.Get(ctx, shortCode)
	if err != nil {
		return nil, err
	}
	return model.ToEntity(), nil
}

// CreateBatch creates multiple links in a batch
func (l *LinkRepository) CreateBatch(ctx context.Context, links []*entity.Link) error {
	m := make([]*models.Link, len(links))
	for i, link := range links {
		m[i] = models.FromEntity(link)
	}
	return l.repo.CreateBatch(ctx, m)
}

func (l *LinkRepository) DeleteBatch(ctx context.Context, ids []string) error {
	args := make([]any, len(ids))
	for i, v := range ids {
		args[i] = v
	}
	return l.repo.DeleteBatch(ctx, args)
}
