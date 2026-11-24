package phrase

import (
	"time"
)

// Note struct for Note
type Note struct {
	// Unique code identifying the note.
	Id string `json:"id,omitempty"`
	// Markdown-formatted content of the note (max 65,535 characters).
	Body      string    `json:"body,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
