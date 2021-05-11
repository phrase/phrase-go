package phrase

import (
	"time"
)

// MemberSpaces struct for MemberSpaces
type MemberSpaces struct {
	Id            string    `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
	ProjectsCount int32     `json:"projects_count,omitempty"`
}
