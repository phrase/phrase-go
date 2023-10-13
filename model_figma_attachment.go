package phrase

import (
	"time"
)

// FigmaAttachment struct for FigmaAttachment
type FigmaAttachment struct {
	Id        string    `json:"id,omitempty"`
	Url       string    `json:"url,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
