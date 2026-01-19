package ports

import (
	"context"

	"go-link/common/pkg/cdc"

	"go-link/redirection/internal/core/entity"
)

type LinkRepository interface {
	GetOriginalURL(ctx context.Context, shortCode string) (*entity.Link, error)
	CreateBulk(ctx context.Context, links []*entity.Link) error
	DeleteBulk(ctx context.Context, ids []string) error
}

type LinkCacheRepository interface {
	Set(ctx context.Context, link *entity.Link) error
	Get(ctx context.Context, id string) (*entity.Link, error)
	DeleteBulk(ctx context.Context, ids []string) error
}

type LinkService interface {
	GetOriginalURL(ctx context.Context, shortCode string) (string, error)
	HandleLinkBatchChange(ctx context.Context, batch []*cdc.DebeziumPayload[entity.Link]) error
}

type LinkConsumer interface {
	Start(ctx context.Context) error
	Stop() error
}
