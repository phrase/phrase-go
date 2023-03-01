package phrase

import (
	"time"
)

// OrganizationJobTemplateDetails struct for OrganizationJobTemplateDetails
type OrganizationJobTemplateDetails struct {
	Id        string          `json:"id,omitempty"`
	Name      string          `json:"name,omitempty"`
	Briefing  string          `json:"briefing,omitempty"`
	CreatedAt time.Time       `json:"created_at,omitempty"`
	UpdatedAt time.Time       `json:"updated_at,omitempty"`
	Owner     UserPreview     `json:"owner,omitempty"`
	Creator   UserPreview     `json:"creator,omitempty"`
	Locales   []LocalePreview `json:"locales,omitempty"`
}
