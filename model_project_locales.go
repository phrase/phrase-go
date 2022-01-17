package phrase

import (
	"time"
)

// ProjectLocales struct for ProjectLocales
type ProjectLocales struct {
	Id          string          `json:"id,omitempty"`
	Name        string          `json:"name,omitempty"`
	ProjectRole string          `json:"project_role,omitempty"`
	MainFormat  string          `json:"main_format,omitempty"`
	CreatedAt   time.Time       `json:"created_at,omitempty"`
	UpdatedAt   time.Time       `json:"updated_at,omitempty"`
	Locales     []LocalePreview `json:"locales,omitempty"`
}
