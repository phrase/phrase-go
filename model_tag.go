package phrase

import (
	"time"
)

// Tag struct for Tag
type Tag struct {
	Name      string `json:"name,omitempty"`
	KeysCount int32  `json:"keys_count,omitempty"`
	// `true` when the tag was created automatically by the system (e.g. for jobs, uploads, or Figma attachments) rather than by a user.
	SystemTag *bool     `json:"system_tag,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
