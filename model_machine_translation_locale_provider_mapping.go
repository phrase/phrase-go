package phrase

// MachineTranslationLocaleProviderMapping struct for MachineTranslationLocaleProviderMapping
type MachineTranslationLocaleProviderMapping struct {
	// The locale code of the source language for this mapping.
	SourceLocaleCode string `json:"source_locale_code,omitempty"`
	// The locale code of the target language for this mapping.
	TargetLocaleCode string `json:"target_locale_code,omitempty"`
	// The translation service applied when translating from source to target locale.
	Service string `json:"service,omitempty"`
}
