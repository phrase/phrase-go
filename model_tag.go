package phrase

import (
	"time"
)

// Tag struct for Tag
type Tag struct {
	Name      string    `json:"name,omitempty"`
	KeysCount int32     `json:"keys_count,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
