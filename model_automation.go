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
	StatusFilters []string  `json:"status_filters,omitempty"`
	ProjectId     string    `json:"project_id,omitempty"`
	JobTemplateId string    `json:"job_template_id,omitempty"`
	Tags          []string  `json:"tags,omitempty"`
	CronSchedule  string    `json:"cron_schedule,omitempty"`
	TimeZone      string    `json:"time_zone,omitempty"`
	Account       Account   `json:"account,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}
