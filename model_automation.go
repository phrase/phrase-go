package phrase

import (
	"time"
)

// Automation struct for Automation
type Automation struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Status  string `json:"status,omitempty"`
	Trigger string `json:"trigger,omitempty"`
	// translation key statuses used to filter keys that are added to jobs
	StatusFilters []string `json:"status_filters,omitempty"`
	ProjectId     string   `json:"project_id,omitempty"`
	// All project IDs the automation applies to. Returned alongside the singular `project_id` for backwards compatibility.
	ProjectIds    []string `json:"project_ids,omitempty"`
	JobTemplateId string   `json:"job_template_id,omitempty"`
	// User ID of the job owner that newly created jobs are assigned to.
	JobOwnerId *NullableString `json:"job_owner_id,omitempty"`
	// When `true`, the automation only acts on locales that changed since its last run.
	IncludeOnlyUpdatedLocales *bool     `json:"include_only_updated_locales,omitempty"`
	Tags                      []string  `json:"tags,omitempty"`
	CronSchedule              string    `json:"cron_schedule,omitempty"`
	TimeZone                  string    `json:"time_zone,omitempty"`
	Account                   Account   `json:"account,omitempty"`
	CreatedAt                 time.Time `json:"created_at,omitempty"`
	UpdatedAt                 time.Time `json:"updated_at,omitempty"`
}
