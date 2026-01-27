package ports

import (
	"context"
	"go-link/generation/internal/core/dto"
	"go-link/generation/internal/core/entity"
)

type LinkRepository interface {
	Create(ctx context.Context, link *entity.Link, ttl int) error
	Get(ctx context.Context, id string) (*entity.Link, error)
	Delete(ctx context.Context, id string) error
}

type LinkCacheRepository interface {
	Set(ctx context.Context, link *entity.Link) error
	IncrementQuota(ctx context.Context, tenantID int) (int64, error)
	DecrementQuota(ctx context.Context, tenantID int) (int64, error)
	GetUserLevel(ctx context.Context, userID int) (int, error)
	SetUserLevel(ctx context.Context, userID int, level int) error
}

type ShortCodePool interface {
	GetOrGenerate() string
}

type LinkService interface {
	Create(ctx context.Context, req *dto.CreateLinkRequest) (*dto.LinkResponse, error)
	Delete(ctx context.Context, req *dto.DeleteLinkRequest) error
}
