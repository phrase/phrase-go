package phrase

import (
	"time"
)

// LocaleDownload struct for LocaleDownload
type LocaleDownload struct {
	Id string `json:"id,omitempty"`
	// The status of the download request. Possible values are `processing`, `completed`, and `error`.
	Status      string               `json:"status,omitempty"`
	Result      LocaleDownloadResult `json:"result,omitempty"`
	Params      LocaleDownloadParams `json:"params,omitempty"`
	Error       string               `json:"error,omitempty"`
	CreatedAt   time.Time            `json:"created_at,omitempty"`
	CompletedAt time.Time            `json:"completed_at,omitempty"`
}
