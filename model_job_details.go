package phrase

import (
	"time"
)

// JobDetails struct for JobDetails
type JobDetails struct {
	Id        string        `json:"id,omitempty"`
	Name      string        `json:"name,omitempty"`
	Briefing  string        `json:"briefing,omitempty"`
	DueDate   *NullableTime `json:"due_date,omitempty"`
	State     string        `json:"state,omitempty"`
	TicketUrl string        `json:"ticket_url,omitempty"`
	Project   ProjectShort  `json:"project,omitempty"`
	Branch    BranchName    `json:"branch,omitempty"`
	CreatedAt time.Time     `json:"created_at,omitempty"`
	UpdatedAt time.Time     `json:"updated_at,omitempty"`
	// The ID of the automation that created this job, or null if the job was created manually.
	AutomationId *NullableString `json:"automation_id,omitempty"`
	// The ID of the job template this job was created from, or null if no template was used.
	JobTemplateId               *NullableString `json:"job_template_id,omitempty"`
	Owner                       UserPreview     `json:"owner,omitempty"`
	JobTagName                  string          `json:"job_tag_name,omitempty"`
	SourceTranslationsUpdatedAt time.Time       `json:"source_translations_updated_at,omitempty"`
	SourceLocale                LocalePreview   `json:"source_locale,omitempty"`
	Locales                     []LocalePreview `json:"locales,omitempty"`
	Keys                        []KeyPreview    `json:"keys,omitempty"`
	// Returned only when `include_annotations=true` is supplied on the request.
	Annotations []JobAnnotationShort `json:"annotations,omitempty"`
	// `true` if the job has been locked by the project's job-locking workflow (translations attached to the job are read-only until the job advances).
	Locked *bool `json:"locked,omitempty"`
}
