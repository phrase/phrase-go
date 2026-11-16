package phrase

// RepoSyncExportSchedule Schedule configuration for automatic exports. Only present when the automatic export feature is enabled for the account. `null` when the `schedule` trigger is not active.
type RepoSyncExportSchedule struct {
	// How often the export runs. `cron` means a fully custom cron expression is used (`cron_schedule` holds the raw expression).
	Frequency string `json:"frequency,omitempty"`
	// The underlying cron expression. Derived automatically from `frequency`, `time`, and day fields when using a named frequency; provided verbatim when `frequency` is `cron`.
	CronSchedule string `json:"cron_schedule,omitempty"`
	// IANA time zone name used to interpret the schedule.
	TimeZone string `json:"time_zone,omitempty"`
	// When the next scheduled export will run, in UTC.
	NextRunAt *NullableTime `json:"next_run_at,omitempty"`
}
