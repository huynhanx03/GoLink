package models

import (
	"go-link/common/pkg/database/widecolumn"
	"go-link/redirection/internal/core/entity"
)

const (
	TableName         = "links"
	OriginalURLColumn = "original_url"
)

type Link struct {
	*widecolumn.BaseModel[string]
	OriginalURL string `json:"original_url"`
}

func (Link) TableName() string {
	return TableName
}

func (Link) ColumnNames() []string {
	return []string{widecolumn.IDColumn, widecolumn.CreatedAtColumn, widecolumn.UpdatedAtColumn, OriginalURLColumn}
}

func (l Link) ColumnValues() []any {
	return []any{l.ID, l.CreatedAt, l.UpdatedAt, l.OriginalURL}
}

func (l *Link) ToEntity() *entity.Link {
	if l == nil {
		return nil
	}
	e := &entity.Link{
		OriginalURL: l.OriginalURL,
	}
	if l.BaseModel != nil {
		e.ID = l.ID
		e.CreatedAt = l.CreatedAt
		e.UpdatedAt = l.UpdatedAt
	}
	return e
}

func FromEntity(e *entity.Link) *Link {
	return &Link{
		BaseModel: &widecolumn.BaseModel[string]{
			ID:        e.ID,
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
		},
		OriginalURL: e.OriginalURL,
	}
}
