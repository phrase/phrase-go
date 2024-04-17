package phrase

import (
	"time"
)

// RepoSyncEvent struct for RepoSyncEvent
type RepoSyncEvent struct {
	EventType string    `json:"event_type,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Status    string    `json:"status,omitempty"`
	// URL of the pull request created on export
	PullRequestUrl string `json:"pull_request_url,omitempty"`
	// Whether the import was triggered by the repo push event
	AutoImport *bool `json:"auto_import,omitempty"`
	// List of error messages, in case of failure
	Errors []RepoSyncEventErrorsInner `json:"errors,omitempty"`
}
