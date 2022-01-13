package phrase

import (
	"time"
)

// Team1 struct for Team1
type Team1 struct {
	Id        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
