package phrase

import (
	"time"
)

// LocaleDetails struct for LocaleDetails
type LocaleDetails struct {
	Id           string           `json:"id,omitempty"`
	Name         string           `json:"name,omitempty"`
	Code         string           `json:"code,omitempty"`
	Default      *bool            `json:"default,omitempty"`
	Main         *bool            `json:"main,omitempty"`
	Rtl          *bool            `json:"rtl,omitempty"`
	PluralForms  []string         `json:"plural_forms,omitempty"`
	SourceLocale LocalePreview    `json:"source_locale,omitempty"`
	CreatedAt    time.Time        `json:"created_at,omitempty"`
	UpdatedAt    time.Time        `json:"updated_at,omitempty"`
	Statistics   LocaleStatistics `json:"statistics,omitempty"`
}
