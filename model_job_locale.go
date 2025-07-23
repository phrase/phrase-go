package phrase

import (
	"time"
)

// JobLocale struct for JobLocale
type JobLocale struct {
	Id                     string               `json:"id,omitempty"`
	Job                    JobPreview           `json:"job,omitempty"`
	Locale                 LocalePreview        `json:"locale,omitempty"`
	Users                  []LocaleUserPreview  `json:"users,omitempty"`
	Teams                  []LocaleTeamPreview  `json:"teams,omitempty"`
	Completed              *bool                `json:"completed,omitempty"`
	TranslationCompletedAt time.Time            `json:"translation_completed_at,omitempty"`
	ReviewCompletedAt      time.Time            `json:"review_completed_at,omitempty"`
	Annotations            []JobAnnotationShort `json:"annotations,omitempty"`
}
