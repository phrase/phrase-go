package phrase
// GlossaryTermTranslationUpdateParameters struct for GlossaryTermTranslationUpdateParameters
type GlossaryTermTranslationUpdateParameters struct {
	// Identifies the language for this translation
	LocaleCode string `json:"locale_code,omitempty"`
	// The content of the translation
	Content string `json:"content,omitempty"`
}
