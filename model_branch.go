package phrase

import (
	"time"
)

// Branch struct for Branch
type Branch struct {
	BaseProjectId   string      `json:"base_project_id,omitempty"`
	BranchProjectId string      `json:"branch_project_id,omitempty"`
	Name            string      `json:"name,omitempty"`
	CreatedAt       time.Time   `json:"created_at,omitempty"`
	UpdatedAt       time.Time   `json:"updated_at,omitempty"`
	MergedAt        time.Time   `json:"merged_at,omitempty"`
	MergedBy        UserPreview `json:"merged_by,omitempty"`
	CreatedBy       UserPreview `json:"created_by,omitempty"`
	State           string      `json:"state,omitempty"`
}
