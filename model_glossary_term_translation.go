package phrase

import (
	"time"
)

// GlossaryTermTranslation struct for GlossaryTermTranslation
type GlossaryTermTranslation struct {
	Id         string    `json:"id,omitempty"`
	LocaleCode string    `json:"locale_code,omitempty"`
	Content    string    `json:"content,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
