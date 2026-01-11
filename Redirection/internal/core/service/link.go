package service

import (
	"context"
	"net/http"

	"go-link/common/pkg/cdc"
	"go-link/common/pkg/common/apperr"
	"go-link/common/pkg/common/http/response"

	"go-link/redirection/global"
	"go-link/redirection/internal/core/entity"
	"go-link/redirection/internal/ports"

	"go.uber.org/zap"
)

type linkService struct {
	linkRepo  ports.LinkRepository
	linkCache ports.LinkCacheRepository
}

func NewLinkService(linkRepo ports.LinkRepository, linkCache ports.LinkCacheRepository) ports.LinkService {
	return &linkService{
		linkRepo:  linkRepo,
		linkCache: linkCache,
	}
}

// GetOriginalURL retrieves the original URL
func (s *linkService) GetOriginalURL(ctx context.Context, shortCode string) (string, error) {
	entity, err := s.linkCache.Get(ctx, shortCode)

	if entity != nil {
		global.Logger.Info("Link found in cache", zap.String("shortCode", shortCode), zap.String("originalURL", entity.OriginalURL))
		return entity.OriginalURL, nil
	}

	entity, err = s.linkRepo.GetOriginalURL(ctx, shortCode)
	if err != nil {
		return "", err
	}

	err = s.linkCache.Set(ctx, entity)
	if err != nil {
		global.Logger.Error("Failed to set link in cache", zap.Error(err))
	}

	global.Logger.Info("Link found in database", zap.String("shortCode", shortCode), zap.String("originalURL", entity.OriginalURL))
	return entity.OriginalURL, nil
}

func (s *linkService) HandleLinkBatchChange(ctx context.Context, batch []*cdc.DebeziumPayload[entity.Link]) error {
	var (
		linksToCreate []*entity.Link
		idsToDelete []string
	)

	for _, payload := range batch {
		switch payload.Op {
		case cdc.OpCreate:
			if payload.After != nil {
				linksToCreate = append(linksToCreate, payload.After)
			}
		case cdc.OpDelete:
			if payload.Before != nil {
				idsToDelete = append(idsToDelete, payload.Before.ID)
			}
		}
	}

	if len(linksToCreate) > 0 {
		if err := s.linkRepo.CreateBatch(ctx, linksToCreate); err != nil {
			return apperr.Wrap(err, response.CodeInternalServer, "failed to batch save link", http.StatusInternalServerError)
		}
	}

	if len(idsToDelete) > 0 {
		if err := s.linkRepo.DeleteBatch(ctx, idsToDelete); err != nil {
			return apperr.Wrap(err, response.CodeInternalServer, "failed to batch remove link", http.StatusInternalServerError)
		}
		if err := s.linkCache.DeleteBatch(ctx, idsToDelete); err != nil {
			return apperr.Wrap(err, response.CodeInternalServer, "failed to batch remove link", http.StatusInternalServerError)
		}
	}

	return nil
}
