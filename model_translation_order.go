package phrase

import (
	"time"
)

// TranslationOrder struct for TranslationOrder
type TranslationOrder struct {
	Id              string          `json:"id,omitempty"`
	Name            string          `json:"name,omitempty"`
	Lsp             string          `json:"lsp,omitempty"`
	AmountInCents   int32           `json:"amount_in_cents,omitempty"`
	Currency        string          `json:"currency,omitempty"`
	Message         string          `json:"message,omitempty"`
	State           string          `json:"state,omitempty"`
	TranslationType string          `json:"translation_type,omitempty"`
	ProgressPercent int32           `json:"progress_percent,omitempty"`
	SourceLocale    LocalePreview   `json:"source_locale,omitempty"`
	TargetLocales   []LocalePreview `json:"target_locales,omitempty"`
	// Name of the tag whose keys are included in the order.
	TagName                          string            `json:"tag_name,omitempty"`
	Styleguide                       StyleguidePreview `json:"styleguide,omitempty"`
	UnverifyTranslationsUponDelivery *bool             `json:"unverify_translations_upon_delivery,omitempty"`
	Quality                          *bool             `json:"quality,omitempty"`
	Priority                         *bool             `json:"priority,omitempty"`
	CreatedAt                        time.Time         `json:"created_at,omitempty"`
	UpdatedAt                        time.Time         `json:"updated_at,omitempty"`
}
