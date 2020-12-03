package phrase

import (
	"time"
)

// Webhook struct for Webhook
type Webhook struct {
	Id              string    `json:"id,omitempty"`
	CallbackUrl     string    `json:"callback_url,omitempty"`
	Description     string    `json:"description,omitempty"`
	Events          []string  `json:"events,omitempty"`
	Active          *bool     `json:"active,omitempty"`
	IncludeBranches *bool     `json:"include_branches,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
}
