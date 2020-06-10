package phrase

// TagWithStats1Statistics struct for TagWithStats1Statistics
type TagWithStats1Statistics struct {
	KeysTotalCount              int32 `json:"keys_total_count,omitempty"`
	TranslationsCompletedCount  int32 `json:"translations_completed_count,omitempty"`
	TranslationsUnverifiedCount int32 `json:"translations_unverified_count,omitempty"`
	KeysUntranslatedCount       int32 `json:"keys_untranslated_count,omitempty"`
}
