package phrase

import (
	"time"
)

// JobTemplate struct for JobTemplate
type JobTemplate struct {
	Id        string       `json:"id,omitempty"`
	Name      string       `json:"name,omitempty"`
	Briefing  string       `json:"briefing,omitempty"`
	Project   ProjectShort `json:"project,omitempty"`
	Branch    Branch       `json:"branch,omitempty"`
	CreatedAt time.Time    `json:"created_at,omitempty"`
	UpdatedAt time.Time    `json:"updated_at,omitempty"`
}
