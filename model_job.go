package phrase

import (
	"time"
)

// Job struct for Job
type Job struct {
	Id        string        `json:"id,omitempty"`
	Name      string        `json:"name,omitempty"`
	Briefing  string        `json:"briefing,omitempty"`
	DueDate   *NullableTime `json:"due_date,omitempty"`
	State     string        `json:"state,omitempty"`
	TicketUrl string        `json:"ticket_url,omitempty"`
	Project   ProjectShort  `json:"project,omitempty"`
	Branch    BranchName    `json:"branch,omitempty"`
	CreatedAt time.Time     `json:"created_at,omitempty"`
	UpdatedAt time.Time     `json:"updated_at,omitempty"`
}
