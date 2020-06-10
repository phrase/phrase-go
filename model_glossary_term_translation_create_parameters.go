package phrase

// GlossaryTermTranslationCreateParameters struct for GlossaryTermTranslationCreateParameters
type GlossaryTermTranslationCreateParameters struct {
	// Identifies the language for this translation
	LocaleCode string `json:"locale_code,omitempty"`
	// The content of the translation
	Content string `json:"content,omitempty"`
}
