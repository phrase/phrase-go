package phrase

import (
	"time"
)

// TeamDetail struct for TeamDetail
type TeamDetail struct {
	Id        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Projects  []Project `json:"projects,omitempty"`
	Spaces    []Space   `json:"spaces,omitempty"`
	Users     []User    `json:"users,omitempty"`
}
