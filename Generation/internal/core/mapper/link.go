package mapper

import (
	"time"

	"go-link/generation/internal/constant"
	"go-link/generation/internal/core/dto"
	"go-link/generation/internal/core/entity"
)

func ToLinkEntityFromReq(req *dto.CreateLinkRequest) *entity.Link {
	return &entity.Link{
		OriginalURL: req.OriginalURL,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func ToLinkResponse(l *entity.Link) *dto.LinkResponse {
	return &dto.LinkResponse{
		ShortLink: constant.URL + "/" + l.ID,
	}
}
