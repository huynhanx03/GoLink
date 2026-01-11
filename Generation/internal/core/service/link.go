package service

import (
	"context"
	"net/http"

	"go-link/common/pkg/common/apperr"
	"go-link/common/pkg/common/http/response"
	"go-link/common/pkg/unique"
	"go-link/generation/global"
	"go-link/generation/internal/core/dto"
	"go-link/generation/internal/core/mapper"
	"go-link/generation/internal/ports"

	"go.uber.org/zap"
)

type linkService struct {
	linkRepo      ports.LinkRepository
	linkCache     ports.LinkCacheRepository
	snowflakeNode *unique.SnowflakeNode
}

func NewLinkService(linkRepo ports.LinkRepository, node *unique.SnowflakeNode, cache ports.LinkCacheRepository) ports.LinkService {
	return &linkService{
		linkRepo:      linkRepo,
		linkCache:     cache,
		snowflakeNode: node,
	}
}

// Generate creates a short link
func (s *linkService) Create(ctx context.Context, req *dto.CreateLinkRequest) (*dto.LinkResponse, error) {
	id := s.snowflakeNode.Generate()
	shortCode := unique.Base62Encode(id)

	link := mapper.ToLinkEntityFromReq(req)
	link.ID = shortCode

	if err := s.linkRepo.Create(ctx, link); err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to create link", http.StatusInternalServerError)
	}

	if err := s.linkCache.Set(ctx, link); err != nil {

	}

	global.Logger.Info("Link created successfully", zap.String("shortCode", shortCode), zap.Any("link", link))

	return mapper.ToLinkResponse(link), nil
}
