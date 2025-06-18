package phrase

// TranslationCreateParameters struct for TranslationCreateParameters
type TranslationCreateParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Locale. Can be the name or id of the locale. Preferred is id
	LocaleId string `json:"locale_id,omitempty"`
	// Key
	KeyId string `json:"key_id,omitempty"`
	// Translation content
	Content string `json:"content,omitempty"`
	// Plural suffix. Can be one of: zero, one, two, few, many, other. Must be specified if the key associated to the translation is pluralized.
	PluralSuffix string `json:"plural_suffix,omitempty"`
	// Indicates whether translation is unverified. Part of the [Advanced Workflows](https://support.phrase.com/hc/en-us/articles/5784094755484) feature.
	Unverified *bool `json:"unverified,omitempty"`
	// Indicates whether translation is excluded.
	Excluded *bool `json:"excluded,omitempty"`
	// Indicates whether the translation should be auto-translated. Responses with status 422 if provided for translation within a non-default locale or the project does not have the Autopilot feature enabled.
	Autotranslate *bool `json:"autotranslate,omitempty"`
}
