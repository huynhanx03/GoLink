package models

import (
	"go-link/common/pkg/database/widecolumn"
	"go-link/generation/internal/core/entity"
)

const (
	TableName         = "links"
	OriginalURLColumn = "original_url"
	UserIDColumn      = "user_id"
	TenantIDColumn    = "tenant_id"
)

type Link struct {
	*widecolumn.BaseModel[string]
	OriginalURL string `json:"original_url"`
	UserID      int    `json:"user_id"`
	TenantID    int    `json:"tenant_id"`
}

func (Link) TableName() string {
	return TableName
}

func (Link) ColumnNames() []string {
	return []string{widecolumn.IDColumn, widecolumn.CreatedAtColumn, widecolumn.UpdatedAtColumn, OriginalURLColumn, UserIDColumn, TenantIDColumn}
}

func (l Link) ColumnValues() []any {
	return []any{l.ID, l.CreatedAt, l.UpdatedAt, l.OriginalURL, l.UserID, l.TenantID}
}

func FromEntity(e *entity.Link) *Link {
	return &Link{
		BaseModel: &widecolumn.BaseModel[string]{
			ID:        e.ID,
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
		},
		OriginalURL: e.OriginalURL,
		UserID:      e.UserID,
		TenantID:    e.TenantID,
	}
}

func (l *Link) ToEntity() *entity.Link {
	return &entity.Link{
		ID:          l.ID,
		OriginalURL: l.OriginalURL,
		UserID:      l.UserID,
		TenantID:    l.TenantID,
		CreatedAt:   l.CreatedAt,
		UpdatedAt:   l.UpdatedAt,
	}
}
