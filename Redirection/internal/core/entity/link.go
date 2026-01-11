package entity

import (
	"time"
)

type Link struct {
	ID          string    `json:"id"`
	OriginalURL string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
