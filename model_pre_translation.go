package phrase

import (
	"time"
)

// PreTranslation struct for PreTranslation
type PreTranslation struct {
	Id string `json:"id,omitempty"`
	// Current execution state of the pre-translation job. Jobs start as `pending` while queued, transition to `running` while executing, and settle to `success` or `error`.
	Status string `json:"status,omitempty"`
	// Resource type that was pre-translated.
	TranslatableType string `json:"translatable_type,omitempty"`
	// ID of the targeted resource (locale ID, job ID, key ID, or upload ID).
	TranslatableId string `json:"translatable_id,omitempty"`
	// Error message. `null` unless the job's status is `error`.
	Error     *NullableString `json:"error,omitempty"`
	CreatedAt time.Time       `json:"created_at,omitempty"`
	UpdatedAt time.Time       `json:"updated_at,omitempty"`
}
