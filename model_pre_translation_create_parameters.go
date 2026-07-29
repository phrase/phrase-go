package phrase

// PreTranslationCreateParameters struct for PreTranslationCreateParameters
type PreTranslationCreateParameters struct {
	// Resource type to pre-translate.
	TranslatableType string `json:"translatable_type,omitempty"`
	// ID of the targeted resource: locale ID for `locale`, job ID for `job`, key ID for `translation_key`, upload ID for `upload`.
	TranslatableId string `json:"translatable_id,omitempty"`
}
