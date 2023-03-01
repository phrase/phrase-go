package phrase

import (
	"time"
)

// OrganizationJobTemplate struct for OrganizationJobTemplate
type OrganizationJobTemplate struct {
	Id        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Briefing  string    `json:"briefing,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
