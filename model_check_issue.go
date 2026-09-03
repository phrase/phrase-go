package phrase

import (
	"time"
)

// CheckIssue struct for CheckIssue
type CheckIssue struct {
	Id string `json:"id,omitempty"`
	// Identifier of the check that reported this issue. One of: `translation_content_length`, `translation_placeholder_usage`, `translation_glossary_usage`.
	CheckName string `json:"check_name,omitempty"`
	// Current state of the check issue. One of: `active`, `solved`, `dismissed`.
	State string `json:"state,omitempty"`
	// Human-readable description of the reported issue, always in English. This message is intended for display only. Its wording may change at any time and it should not be parsed or relied upon programmatically.
	Description string        `json:"description,omitempty"`
	DismissedAt *NullableTime `json:"dismissed_at,omitempty"`
	SolvedAt    *NullableTime `json:"solved_at,omitempty"`
	CreatedAt   time.Time     `json:"created_at,omitempty"`
	UpdatedAt   time.Time     `json:"updated_at,omitempty"`
	Translation Translation   `json:"translation,omitempty"`
}
