package phrase

// ReleaseCreateParameters1 struct for ReleaseCreateParameters1
type ReleaseCreateParameters1 struct {
	// Cron schedule for the scheduler. Read more about the format of this field at https://en.wikipedia.org/wiki/Cron
	CronSchedule string `json:"cron_schedule,omitempty"`
	// Time zone for the scheduler
	TimeZone string `json:"time_zone,omitempty"`
	// List of locale ids that will be included in the release.
	LocaleIds []string `json:"locale_ids,omitempty"`
	// Only include tagged keys in the release. For a key to be included it must be tagged with all tags provided
	Tags []string `json:"tags,omitempty"`
	// Branch used for release
	Branch string `json:"branch,omitempty"`
	// The created releases will be available only for apps with version greater or equal to this value
	AppMinVersion string `json:"app_min_version,omitempty"`
	// The created releases will be available only for apps with version less or equal to this value
	AppMaxVersion string `json:"app_max_version,omitempty"`
}
