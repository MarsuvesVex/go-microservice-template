package events

import (
	"encoding/json"
	"time"
)

type Event struct {
	ID          string          `json:"id,omitempty"`
	Topic       string          `json:"topic"`
	Source      string          `json:"source"`
	OccurredAt  time.Time       `json:"occurred_at,omitempty"`
	PublishedAt time.Time       `json:"published_at,omitempty"`
	Data        json.RawMessage `json:"data"`
	Metadata    map[string]any  `json:"metadata,omitempty"`
}
