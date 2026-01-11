package link

import (
	"encoding/json"
	"time"

	"go-link/redirection/internal/core/entity"
)

type CDCLink struct {
	ID          string    `json:"id"`
	OriginalURL CDCString `json:"original_url"`
	CreatedAt   CDCTime   `json:"created_at"`
	UpdatedAt   CDCTime   `json:"updated_at"`
}

type CDCString struct {
	Value string
}

func (s *CDCString) UnmarshalJSON(b []byte) error {
	// Try standard string
	var str string
	if err := json.Unmarshal(b, &str); err == nil {
		s.Value = str
		return nil
	}

	// Try wrapped object {"value": "..."}
	var obj struct {
		Value string `json:"value"`
	}
	if err := json.Unmarshal(b, &obj); err != nil {
		return err
	}
	s.Value = obj.Value
	return nil
}

type CDCTime struct {
	time.Time
}

func (t *CDCTime) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	switch val := v.(type) {
	case float64:
		t.Time = time.UnixMilli(int64(val))
		return nil
	case string:
		parsedTime, err := time.Parse(time.RFC3339, val)
		if err == nil {
			t.Time = parsedTime
			return nil
		}
	case map[string]interface{}:
		if numVal, ok := val["value"].(float64); ok {
			t.Time = time.UnixMilli(int64(numVal))
			return nil
		}
	}
	return nil
}

func (c *CDCLink) ToEntity() *entity.Link {
	return &entity.Link{
		ID:          c.ID,
		OriginalURL: c.OriginalURL.Value,
		CreatedAt:   c.CreatedAt.Time,
		UpdatedAt:   c.UpdatedAt.Time,
	}
}
