package phrase

// ProjectReport struct for ProjectReport
type ProjectReport struct {
	LocalesCount                int32        `json:"locales_count,omitempty"`
	KeysCount                   int32        `json:"keys_count,omitempty"`
	TranslationsCount           int32        `json:"translations_count,omitempty"`
	UntranslatedKeysCount       int32        `json:"untranslated_keys_count,omitempty"`
	UnverifiedTranslationsCount int32        `json:"unverified_translations_count,omitempty"`
	ReviewedTranslationsCount   int32        `json:"reviewed_translations_count,omitempty"`
	ManagedWordsCount           int32        `json:"managed_words_count,omitempty"`
	Project                     ProjectShort `json:"project,omitempty"`
}
