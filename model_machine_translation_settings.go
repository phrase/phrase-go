package phrase

// MachineTranslationSettings struct for MachineTranslationSettings
type MachineTranslationSettings struct {
	// The default machine translation engine configured for the account. Returns \"microsoft_translate\" when no service has been explicitly configured. Supported values: language_ai_translate, aita_translate, microsoft_translate, google_translate, amazon_translate, intento_translate, gpt_translate.
	DefaultService *NullableString `json:"default_service,omitempty"`
	// Number of machine translation characters consumed in the current billing period.
	MachineTranslationUnitsUsed int32 `json:"machine_translation_units_used,omitempty"`
	// Total machine translation character quota granted for the current billing period.
	MachineTranslationUnitsTotal int32 `json:"machine_translation_units_total,omitempty"`
	// Per-locale-pair provider overrides. When a matching mapping exists for a source/target locale pair, that provider takes precedence over the account default.
	LocaleProviderMappings []MachineTranslationLocaleProviderMapping `json:"locale_provider_mappings,omitempty"`
}
