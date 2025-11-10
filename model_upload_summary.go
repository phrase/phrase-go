package phrase

// UploadSummary struct for UploadSummary
type UploadSummary struct {
	LocalesCreated         int32 `json:"locales_created,omitempty"`
	TranslationKeysCreated int32 `json:"translation_keys_created,omitempty"`
	TranslationKeysUpdated int32 `json:"translation_keys_updated,omitempty"`
	// The number of translation keys in the project that were not mentioned in the upload.  Note: this field is not calculated (and shows 0) if the upload contains more than 10,000 keys.
	TranslationKeysUnmentioned int32 `json:"translation_keys_unmentioned,omitempty"`
	TranslationsCreated        int32 `json:"translations_created,omitempty"`
	TranslationsUpdated        int32 `json:"translations_updated,omitempty"`
	TagsCreated                int32 `json:"tags_created,omitempty"`
	TranslationKeysIgnored     int32 `json:"translation_keys_ignored,omitempty"`
	ProcessedTranslations      int32 `json:"processed_translations,omitempty"`
	UploadTotalTranslations    int32 `json:"upload_total_translations,omitempty"`
}
