package phrase

import (
	"time"
)

// RepoSync struct for RepoSync
type RepoSync struct {
	Id           string       `json:"id,omitempty"`
	Project      ProjectShort `json:"project,omitempty"`
	Provider     string       `json:"provider,omitempty"`
	Enabled      *bool        `json:"enabled,omitempty"`
	AutoImport   *bool        `json:"auto_import,omitempty"`
	RepoName     string       `json:"repo_name,omitempty"`
	CreatedAt    time.Time    `json:"created_at,omitempty"`
	LastImportAt time.Time    `json:"last_import_at,omitempty"`
	LastExportAt time.Time    `json:"last_export_at,omitempty"`
}
