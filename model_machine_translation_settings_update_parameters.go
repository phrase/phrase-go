package phrase

// MachineTranslationSettingsUpdateParameters struct for MachineTranslationSettingsUpdateParameters
type MachineTranslationSettingsUpdateParameters struct {
	// The machine translation engine to use as the account default. Supported values: language_ai_translate, aita_translate, microsoft_translate, google_translate, amazon_translate, intento_translate, gpt_translate. Pass null or an empty string to reset to the plan default.
	DefaultService *NullableString `json:"default_service,omitempty"`
}
