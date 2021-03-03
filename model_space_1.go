package phrase

import (
	"time"
)

// Space1 struct for Space1
type Space1 struct {
	Id            string    `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
	ProjectsCount int32     `json:"projects_count,omitempty"`
}
