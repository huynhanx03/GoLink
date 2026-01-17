package service

import (
	"context"
	"net/http"

	"go-link/common/pkg/common/apperr"
	"go-link/common/pkg/common/http/response"
	"go-link/generation/global"
	"go-link/generation/internal/core/dto"
	"go-link/generation/internal/core/mapper"
	"go-link/generation/internal/ports"

	"go.uber.org/zap"
)

type linkService struct {
	linkRepo  ports.LinkRepository
	linkCache ports.LinkCacheRepository
	codePool  ports.ShortCodePool
}

func NewLinkService(linkRepo ports.LinkRepository, codePool ports.ShortCodePool, cache ports.LinkCacheRepository) ports.LinkService {
	return &linkService{
		linkRepo:  linkRepo,
		linkCache: cache,
		codePool:  codePool,
	}
}

// Create creates a short link
func (s *linkService) Create(ctx context.Context, req *dto.CreateLinkRequest) (*dto.LinkResponse, error) {
	shortCode := s.codePool.GetOrGenerate()

	link := mapper.ToLinkEntityFromReq(req)
	link.ID = shortCode

	if err := s.linkRepo.Create(ctx, link); err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to create link", http.StatusInternalServerError)
	}

	if err := s.linkCache.Set(ctx, link); err != nil {
		global.Logger.Warn("Failed to cache link", zap.Error(err), zap.String("shortCode", shortCode))
	}

	global.Logger.Info("Link created successfully", zap.String("shortCode", shortCode), zap.Any("link", link))

	return mapper.ToLinkResponse(link), nil
}
