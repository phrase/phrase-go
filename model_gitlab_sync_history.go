package phrase

import (
	"time"
)

// GitlabSyncHistory struct for GitlabSyncHistory
type GitlabSyncHistory struct {
	Status  string                    `json:"status,omitempty"`
	Action  string                    `json:"action,omitempty"`
	Errors  []GitlabSyncHistoryErrors `json:"errors,omitempty"`
	Date    time.Time                 `json:"date,omitempty"`
	Details map[string]interface{}    `json:"details,omitempty"`
}
