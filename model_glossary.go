package phrase

import (
	"time"
)

// Glossary struct for Glossary
type Glossary struct {
	Id        string         `json:"id,omitempty"`
	Name      string         `json:"name,omitempty"`
	Projects  []ProjectShort `json:"projects,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
}
