package phrase

// MachineTranslationLocaleProviderMappingsCreateParameters struct for MachineTranslationLocaleProviderMappingsCreateParameters
type MachineTranslationLocaleProviderMappingsCreateParameters struct {
	// The locale code of the source language (e.g. \"en\").
	SourceLocaleCode string `json:"source_locale_code"`
	// The locale code of the target language (e.g. \"de\"). Must differ from source_locale_code.
	TargetLocaleCode string `json:"target_locale_code"`
	// The machine translation service to use for this locale pair. Must be a service enabled for the account.
	Service string `json:"service"`
}
