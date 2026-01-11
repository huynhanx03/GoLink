package mapper

import (
	"go-link/redirection/internal/constant"
	"go-link/redirection/internal/core/dto"
	"go-link/redirection/internal/core/entity"
)

func ToLinkResponse(l *entity.Link) *dto.LinkResponse {
	return &dto.LinkResponse{
		ShortLink: constant.URL + "/" + l.ID,
	}
}
