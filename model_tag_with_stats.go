package phrase

import (
	"time"
)

// TagWithStats struct for TagWithStats
type TagWithStats struct {
	Name       string                     `json:"name,omitempty"`
	KeysCount  int32                      `json:"keys_count,omitempty"`
	CreatedAt  time.Time                  `json:"created_at,omitempty"`
	UpdatedAt  time.Time                  `json:"updated_at,omitempty"`
	Statistics []TagWithStats1Statistics1 `json:"statistics,omitempty"`
}
