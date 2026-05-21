package phrase

import (
	"time"
)

// Distribution struct for Distribution
type Distribution struct {
	Id           string        `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Project      ProjectShort  `json:"project,omitempty"`
	Platforms    []string      `json:"platforms,omitempty"`
	ReleaseCount int32         `json:"release_count,omitempty"`
	CreatedAt    time.Time     `json:"created_at,omitempty"`
	UpdatedAt    time.Time     `json:"updated_at,omitempty"`
	DeletedAt    *NullableTime `json:"deleted_at,omitempty"`
}
