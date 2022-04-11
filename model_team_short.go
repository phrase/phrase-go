package phrase

import (
	"time"
)

// TeamShort struct for TeamShort
type TeamShort struct {
	Id        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
