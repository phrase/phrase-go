package phrase

import (
	"time"
)

// KeyLink struct for KeyLink
type KeyLink struct {
	// The timestamp when the link was created.
	CreatedAt time.Time `json:"created_at"`
	// The timestamp when the link was last updated.
	UpdatedAt time.Time   `json:"updated_at"`
	CreatedBy UserPreview `json:"created_by"`
	UpdatedBy UserPreview `json:"updated_by"`
	Account   Account     `json:"account"`
	Parent    KeyPreview  `json:"parent"`
	// The child translation keys linked to the parent.
	Children []KeyPreview `json:"children"`
}
