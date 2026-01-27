package entity

import (
	"time"
)

type Link struct {
	ID          string    `json:"id"`
	OriginalURL string    `json:"original_url"`
	UserID      int       `json:"user_id"`
	TenantID    int       `json:"tenant_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
