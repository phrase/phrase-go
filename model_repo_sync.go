package phrase

import (
	"time"
)

// RepoSync struct for RepoSync
type RepoSync struct {
	Id string `json:"id,omitempty"`
	// Optional custom display name for this repo sync. When null or blank, the sync is displayed using the associated project name.
	Name       *NullableString `json:"name,omitempty"`
	Project    ProjectShort    `json:"project,omitempty"`
	Provider   string          `json:"provider,omitempty"`
	Enabled    *bool           `json:"enabled,omitempty"`
	AutoImport *bool           `json:"auto_import,omitempty"`
	RepoName   string          `json:"repo_name,omitempty"`
	// Branch used as the source of exports/PRs. May be `null` when the sync is configured to push directly to `base_branch`.
	PrBranch     *NullableString `json:"pr_branch,omitempty"`
	CreatedAt    time.Time       `json:"created_at,omitempty"`
	LastImportAt *NullableTime   `json:"last_import_at,omitempty"`
	LastExportAt *NullableTime   `json:"last_export_at,omitempty"`
	Owner        UserShort       `json:"owner,omitempty"`
	// Events that trigger an automatic export. Only present when the automatic export feature is enabled for the account. Possible values: `schedule` (runs on the configured schedule), `branch_merged` (runs when a branch is merged in the connected repository), `job_completed` (runs when a translation job completes in the project).
	ExportTriggers []string                `json:"export_triggers,omitempty"`
	ExportSchedule *RepoSyncExportSchedule `json:"export_schedule,omitempty"`
}
