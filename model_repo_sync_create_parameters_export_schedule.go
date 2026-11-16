package phrase

// RepoSyncCreateParametersExportSchedule Schedule settings for automatic exports. Required when `export_triggers` includes `schedule`. Requires the automatic export feature to be enabled.
type RepoSyncCreateParametersExportSchedule struct {
	Frequency string `json:"frequency,omitempty"`
	// Time of day for the export (HH:MM format). Used with `day`, `week`, and `month` frequencies.
	Time string `json:"time,omitempty"`
	// IANA time zone name.
	TimeZone string `json:"time_zone,omitempty"`
	// Comma-separated day numbers (0=Sunday … 6=Saturday). Used with `week` frequency.
	DaysOfWeek string `json:"days_of_week,omitempty"`
	// Day of the month (1–31). Used with `month` frequency.
	DayOfMonth int32 `json:"day_of_month,omitempty"`
	// Raw cron expression. Used when `frequency` is `cron`.
	CronSchedule string `json:"cron_schedule,omitempty"`
}
