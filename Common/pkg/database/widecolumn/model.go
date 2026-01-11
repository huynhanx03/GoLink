package widecolumn

import (
	"go-link/common/pkg/constraints"
	"time"
)

const (
	IDColumn        = "id"
	CreatedAtColumn = "created_at"
	UpdatedAtColumn = "updated_at"
)

type BaseModel[ID constraints.ID] struct {
	ID        ID        `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Model interface {
	TableName() string
	ColumnNames() []string
	ColumnValues() []any
}
