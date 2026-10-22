package phrase

import (
	"time"
)

// Glossary struct for Glossary
type Glossary struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	// Number of terms in the glossary.
	TermsCount int32          `json:"terms_count,omitempty"`
	Projects   []ProjectShort `json:"projects,omitempty"`
	Spaces     []Space        `json:"spaces,omitempty"`
	CreatedAt  time.Time      `json:"created_at,omitempty"`
	UpdatedAt  time.Time      `json:"updated_at,omitempty"`
}
