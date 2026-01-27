package phrase

import (
	"time"
)

// UploadBatch struct for UploadBatch
type UploadBatch struct {
	// Processing state of the upload batch
	State string `json:"state,omitempty"`
	// Indicates whether unmentioned keys will be deleted after processing all uploads in the batch
	DeleteUnmentionedKeys *bool        `json:"delete_unmentioned_keys,omitempty"`
	CreatedAt             time.Time    `json:"created_at,omitempty"`
	UpdatedAt             time.Time    `json:"updated_at,omitempty"`
	Project               ProjectShort `json:"project,omitempty"`
	User                  UserPreview  `json:"user,omitempty"`
	Uploads               []Upload     `json:"uploads,omitempty"`
}
