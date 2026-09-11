package phrase

import (
	"time"
)

// Upload struct for Upload
type Upload struct {
	Id       string `json:"id,omitempty"`
	Filename string `json:"filename,omitempty"`
	Format   string `json:"format,omitempty"`
	// Processing state of the upload: `initialized`, `processing`, `success`, or `error`. `error` means processing failed — for example the file could not be parsed, or a provided `file_format` didn't match the file's actual content. Poll this field until it leaves `initialized`/`processing` to get the final outcome.
	State string `json:"state,omitempty"`
	// A user-facing message explaining why the upload failed, or `null` if the upload did not fail.  This message is intended for display only. Its wording may change at any time and it should not be parsed or relied upon programmatically.
	ErrorMessage *NullableString `json:"error_message,omitempty"`
	// Unique tag of the upload
	Tag string `json:"tag,omitempty"`
	// List of tags that were assigned to the uploaded keys
	Tags []string `json:"tags,omitempty"`
	// The URL to the upload in Phrase Strings app.
	Url       string        `json:"url,omitempty"`
	User      UserPreview   `json:"user,omitempty"`
	Summary   UploadSummary `json:"summary,omitempty"`
	CreatedAt time.Time     `json:"created_at,omitempty"`
	UpdatedAt time.Time     `json:"updated_at,omitempty"`
}
