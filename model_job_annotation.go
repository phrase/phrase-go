package phrase

import (
	"time"
)

// JobAnnotation struct for JobAnnotation
type JobAnnotation struct {
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
