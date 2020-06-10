package phrase

import (
	"time"
)

// DistributionPreview struct for DistributionPreview
type DistributionPreview struct {
	Id           string       `json:"id,omitempty"`
	Name         string       `json:"name,omitempty"`
	Project      ProjectShort `json:"project,omitempty"`
	Platforms    []string     `json:"platforms,omitempty"`
	ReleaseCount int32        `json:"release_count,omitempty"`
	CreatedAt    time.Time    `json:"created_at,omitempty"`
	DeletedAt    time.Time    `json:"deleted_at,omitempty"`
}
