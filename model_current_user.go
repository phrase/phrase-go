package phrase

import (
	"time"
)

// CurrentUser struct for CurrentUser
type CurrentUser struct {
	Id        string    `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Position  string    `json:"position,omitempty"`
	Language  string    `json:"language,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
