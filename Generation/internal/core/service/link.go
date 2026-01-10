package service

import (
	"context"
	"net/http"

	"go-link/common/pkg/common/apperr"
	"go-link/common/pkg/common/http/response"
	"go-link/generation/internal/core/dto"
	"go-link/generation/internal/core/mapper"
	"go-link/generation/internal/ports"
	"go-link/generation/pkg/utils"
)

type linkService struct {
	linkRepo      ports.LinkRepository
	snowflakeNode *utils.Node
}

func NewLinkService(linkRepo ports.LinkRepository) ports.LinkService {
	node, err := utils.NewNode(1)
	if err != nil {
		return nil
	}

	return &linkService{
		linkRepo:      linkRepo,
		snowflakeNode: node,
	}
}

// Generate creates a short link
func (s *linkService) Create(ctx context.Context, req *dto.CreateLinkRequest) (*dto.LinkResponse, error) {
	id := s.snowflakeNode.Generate()
	shortCode := utils.Base62Encode(id)

	link := mapper.ToLinkEntityFromReq(req)
	link.ID = shortCode

	if err := s.linkRepo.Create(ctx, link); err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to create link", http.StatusInternalServerError)
	}

	return mapper.ToLinkResponse(link), nil
}
