package phrase

import (
	"time"
)

// TranslationDetails struct for TranslationDetails
type TranslationDetails struct {
	Id           string        `json:"id,omitempty"`
	Content      string        `json:"content"`
	Unverified   *bool         `json:"unverified,omitempty"`
	Excluded     *bool         `json:"excluded,omitempty"`
	PluralSuffix string        `json:"plural_suffix,omitempty"`
	Key          KeyPreview    `json:"key,omitempty"`
	Locale       LocalePreview `json:"locale,omitempty"`
	Placeholders []string      `json:"placeholders,omitempty"`
	State        string        `json:"state,omitempty"`
	CreatedAt    time.Time     `json:"created_at,omitempty"`
	UpdatedAt    time.Time     `json:"updated_at,omitempty"`
	User         UserPreview   `json:"user,omitempty"`
	WordCount    int32         `json:"word_count,omitempty"`
}
