package phrase

import (
	"time"
)

// ProjectMember struct for ProjectMember
type ProjectMember struct {
	Id        string    `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Name      string    `json:"name,omitempty"`
	Position  string    `json:"position,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
