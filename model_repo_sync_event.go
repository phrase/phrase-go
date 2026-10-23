package phrase

import (
	"time"
)

// RepoSyncEvent struct for RepoSyncEvent
type RepoSyncEvent struct {
	Id        string    `json:"id,omitempty"`
	Type      string    `json:"type,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Status    string    `json:"status,omitempty"`
	// URL of the pull request created on export
	PullRequestUrl string `json:"pull_request_url,omitempty"`
	// Whether the import was triggered by the repo push event
	AutoImport *bool `json:"auto_import,omitempty"`
	// List of error messages, in case of failure
	Errors []string `json:"errors,omitempty"`
	// The uploads associated with this import event, or `null` if no upload batch has been created yet. Each entry identifies a single file upload by its public code (`id`) and its current public state (`state`). Present on show responses only; absent from list responses. Possible `state` values: `initialized`, `waiting_for_preview`, `processing`, `success`, `error`.
	Uploads *[]RepoSyncEventUploadsInner `json:"uploads,omitempty"`
}
