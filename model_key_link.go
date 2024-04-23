package phrase

import (
	"time"
)

// KeyLink struct for KeyLink
type KeyLink struct {
	// The timestamp when the link was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// The timestamp when the link was last updated.
	UpdatedAt time.Time   `json:"updated_at,omitempty"`
	CreatedBy UserPreview `json:"created_by,omitempty"`
	UpdatedBy UserPreview `json:"updated_by,omitempty"`
	Account   Account     `json:"account,omitempty"`
	Parent    KeyPreview  `json:"parent,omitempty"`
	// The child translation keys linked to the parent.
	Children []KeyPreview `json:"children,omitempty"`
}
