package phrase

import (
	"time"
)

// ReleaseTrigger struct for ReleaseTrigger
type ReleaseTrigger struct {
	Id     string `json:"id,omitempty"`
	Branch string `json:"branch,omitempty"`
	// Cron schedule for the scheduler. Read more about the format of this field at https://en.wikipedia.org/wiki/Cron
	CronSchedule string `json:"cron_schedule,omitempty"`
	// Time zone for the scheduler
	TimeZone string `json:"time_zone,omitempty"`
	// The next time a release will be created for this trigger
	NextRunAt     time.Time       `json:"next_run_at,omitempty"`
	AppMinVersion string          `json:"app_min_version,omitempty"`
	AppMaxVersion string          `json:"app_max_version,omitempty"`
	Locales       []LocalePreview `json:"locales,omitempty"`
	Tags          []string        `json:"tags,omitempty"`
	CreatedAt     time.Time       `json:"created_at,omitempty"`
	UpdatedAt     time.Time       `json:"updated_at,omitempty"`
}
