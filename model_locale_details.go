package phrase

import (
	"time"
)

// LocaleDetails struct for LocaleDetails
type LocaleDetails struct {
	Id                 string        `json:"id,omitempty"`
	Name               string        `json:"name,omitempty"`
	Code               string        `json:"code,omitempty"`
	Default            *bool         `json:"default,omitempty"`
	Main               *bool         `json:"main,omitempty"`
	Rtl                *bool         `json:"rtl,omitempty"`
	PluralForms        []string      `json:"plural_forms,omitempty"`
	OrdinalPluralForms []string      `json:"ordinal_plural_forms,omitempty"`
	SourceLocale       LocalePreview `json:"source_locale,omitempty"`
	FallbackLocale     LocalePreview `json:"fallback_locale,omitempty"`
	LanguageAiProfile  string        `json:"language_ai_profile,omitempty"`
	// Indicates that new translations for this locale are marked as unverified. Only applies to locales using the basic verification workflow. Part of the [Advanced Workflows](https://support.phrase.com/hc/en-us/articles/5784094755484) feature.
	UnverifyNewTranslations *bool `json:"unverify_new_translations,omitempty"`
	// Indicates that updated translations for this locale are marked as unverified. Only applies to locales using the basic verification workflow. Part of the [Advanced Workflows](https://support.phrase.com/hc/en-us/articles/5784094755484) feature.
	UnverifyUpdatedTranslations *bool `json:"unverify_updated_translations,omitempty"`
	// Indicates that translations for this locale are marked as unverified when the source language has been changed.
	UnverifyOnSourceChanges *bool            `json:"unverify_on_source_changes,omitempty"`
	CreatedAt               time.Time        `json:"created_at,omitempty"`
	UpdatedAt               time.Time        `json:"updated_at,omitempty"`
	Statistics              LocaleStatistics `json:"statistics,omitempty"`
}
