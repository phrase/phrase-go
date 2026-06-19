package phrase

import (
	"time"
)

// Job struct for Job
type Job struct {
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
	JobTemplateId *NullableString `json:"job_template_id,omitempty"`
	// The review due date for this job. Returns `null` when the project does not have review workflow enabled.
	ReviewDueDate *NullableTime `json:"review_due_date,omitempty"`
}
