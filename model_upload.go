package phrase

import (
	"time"
)

// Upload struct for Upload
type Upload struct {
	Id       string `json:"id,omitempty"`
	Filename string `json:"filename,omitempty"`
	Format   string `json:"format,omitempty"`
	State    string `json:"state,omitempty"`
	Tag      string `json:"tag,omitempty"`
	// The URL to the upload in Phrase Strings app.
	Url       string        `json:"url,omitempty"`
	Summary   UploadSummary `json:"summary,omitempty"`
	CreatedAt time.Time     `json:"created_at,omitempty"`
	UpdatedAt time.Time     `json:"updated_at,omitempty"`
}
