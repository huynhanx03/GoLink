package entity

import (
	"go-link/common/pkg/database/widecolumn"
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
	return []interface{}{l.ID, l.CreatedAt, l.UpdatedAt, l.OriginalURL}
}
