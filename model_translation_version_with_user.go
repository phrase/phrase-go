package phrase

import (
	"time"
)

// TranslationVersionWithUser struct for TranslationVersionWithUser
type TranslationVersionWithUser struct {
	Id           string        `json:"id,omitempty"`
	Content      string        `json:"content,omitempty"`
	PluralSuffix string        `json:"plural_suffix,omitempty"`
	Key          KeyPreview    `json:"key,omitempty"`
	Locale       LocalePreview `json:"locale,omitempty"`
	CreatedAt    time.Time     `json:"created_at,omitempty"`
	UpdatedAt    time.Time     `json:"updated_at,omitempty"`
	ChangedAt    time.Time     `json:"changed_at,omitempty"`
	User         UserPreview   `json:"user,omitempty"`
}
