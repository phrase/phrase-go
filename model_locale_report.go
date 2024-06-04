package phrase

// LocaleReport struct for LocaleReport
type LocaleReport struct {
	KeysCount                        int32         `json:"keys_count,omitempty"`
	TranslatedTranslationsPercentage float32       `json:"translated_translations_percentage,omitempty"`
	UnverifiedTranslationsPercentage float32       `json:"unverified_translations_percentage,omitempty"`
	ReviewedTranslationsPercentage   float32       `json:"reviewed_translations_percentage,omitempty"`
	UntranslatedKeysPercentage       float32       `json:"untranslated_keys_percentage,omitempty"`
	CompletedTranslationsCount       int32         `json:"completed_translations_count,omitempty"`
	UntranslatedKeysCount            int32         `json:"untranslated_keys_count,omitempty"`
	UnverifiedTranslationsCount      int32         `json:"unverified_translations_count,omitempty"`
	ReviewedTranslationsCount        int32         `json:"reviewed_translations_count,omitempty"`
	SourceWordCount                  int32         `json:"source_word_count,omitempty"`
	WordCount                        int32         `json:"word_count,omitempty"`
	WordCountUnverified              int32         `json:"word_count_unverified,omitempty"`
	WordCountMissing                 int32         `json:"word_count_missing,omitempty"`
	Locale                           LocalePreview `json:"locale,omitempty"`
}
