package phrase

import (
	"time"
)

// OrganizationJobTemplateDetails struct for OrganizationJobTemplateDetails
type OrganizationJobTemplateDetails struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Briefing string `json:"briefing,omitempty"`
	// When `true`, jobs created from this template are auto-translated on creation. Maps to the `autotranslate` field on the request body.
	AutotranslateEnabled *bool `json:"autotranslate_enabled,omitempty"`
	// Optional. ID of the source locale used by jobs created from this template. When omitted, the project's default source locale is used.
	SourceLocaleId *NullableString `json:"source_locale_id,omitempty"`
	CreatedAt      time.Time       `json:"created_at,omitempty"`
	UpdatedAt      time.Time       `json:"updated_at,omitempty"`
	Owner          UserPreview     `json:"owner,omitempty"`
	Creator        UserPreview     `json:"creator,omitempty"`
	Locales        []LocalePreview `json:"locales,omitempty"`
}
