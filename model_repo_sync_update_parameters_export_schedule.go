package phrase

// RepoSyncUpdateParametersExportSchedule Updates the schedule for automatic exports. Only applied when `export_triggers` includes `schedule`. Requires `manage_github_sync?` permission and the automatic export feature to be enabled.
type RepoSyncUpdateParametersExportSchedule struct {
	Frequency    string `json:"frequency,omitempty"`
	Time         string `json:"time,omitempty"`
	TimeZone     string `json:"time_zone,omitempty"`
	DaysOfWeek   string `json:"days_of_week,omitempty"`
	DayOfMonth   int32  `json:"day_of_month,omitempty"`
	CronSchedule string `json:"cron_schedule,omitempty"`
}
