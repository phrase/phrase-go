package phrase

import (
	"time"
)

// JobDetails struct for JobDetails
type JobDetails struct {
	Id           string          `json:"id,omitempty"`
	Name         string          `json:"name,omitempty"`
	Briefing     string          `json:"briefing,omitempty"`
	DueDate      *NullableTime   `json:"due_date,omitempty"`
	State        string          `json:"state,omitempty"`
	TicketUrl    string          `json:"ticket_url,omitempty"`
	Project      ProjectShort    `json:"project,omitempty"`
	Branch       BranchName      `json:"branch,omitempty"`
	CreatedAt    time.Time       `json:"created_at,omitempty"`
	UpdatedAt    time.Time       `json:"updated_at,omitempty"`
	Owner        UserPreview     `json:"owner,omitempty"`
	JobTagName   string          `json:"job_tag_name,omitempty"`
	SourceLocale LocalePreview   `json:"source_locale,omitempty"`
	Locales      []LocalePreview `json:"locales,omitempty"`
	Keys         []KeyPreview    `json:"keys,omitempty"`
}
