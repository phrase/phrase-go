package phrase

import (
	"time"
)

// Variable struct for Variable
type Variable struct {
	Name      string    `json:"name,omitempty"`
	Value     string    `json:"value,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
