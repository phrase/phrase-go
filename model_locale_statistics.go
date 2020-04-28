package phrase
// LocaleStatistics struct for LocaleStatistics
type LocaleStatistics struct {
	KeysTotalCount int32 `json:"keys_total_count,omitempty"`
	KeysUntranslatedCount int32 `json:"keys_untranslated_count,omitempty"`
	WordsTotalCount int32 `json:"words_total_count,omitempty"`
	TranslationsCompletedCount int32 `json:"translations_completed_count,omitempty"`
	TranslationsUnverifiedCount int32 `json:"translations_unverified_count,omitempty"`
	UnverifiedWordsCount int32 `json:"unverified_words_count,omitempty"`
	MissingWordsCount int32 `json:"missing_words_count,omitempty"`
}
