package phrase

// RepoSyncUpdateParameters struct for RepoSyncUpdateParameters
type RepoSyncUpdateParameters struct {
	// Display name for the repo sync.
	Name string `json:"name,omitempty"`
	// Code of the user to assign as owner. Requires `manage_github_sync?` permission on the project; returns `403` if not satisfied.
	OwnerId string `json:"owner_id,omitempty"`
	// Replaces the current set of automatic export triggers. Requires `manage_github_sync?` permission on the project and the automatic export feature to be enabled for the account. Accepted values: `schedule`, `branch_merged`, `job_completed`.
	ExportTriggers []string                               `json:"export_triggers,omitempty"`
	ExportSchedule RepoSyncUpdateParametersExportSchedule `json:"export_schedule,omitempty"`
}
