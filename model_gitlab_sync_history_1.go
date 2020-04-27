package phrase
import (
	"time"
)
// GitlabSyncHistory1 struct for GitlabSyncHistory1
type GitlabSyncHistory1 struct {
	Status int32 `json:"status,omitempty"`
	Action string `json:"action,omitempty"`
	Errors []string `json:"errors,omitempty"`
	Date time.Time `json:"date,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
}
