package phrase

// UploadSummary struct for UploadSummary
type UploadSummary struct {
	LocalesCreated             int32 `json:"locales_created,omitempty"`
	TranslationKeysCreated     int32 `json:"translation_keys_created,omitempty"`
	TranslationKeysUpdated     int32 `json:"translation_keys_updated,omitempty"`
	TranslationKeysUnmentioned int32 `json:"translation_keys_unmentioned,omitempty"`
	TranslationsCreated        int32 `json:"translations_created,omitempty"`
	TranslationsUpdated        int32 `json:"translations_updated,omitempty"`
	TagsCreated                int32 `json:"tags_created,omitempty"`
	TranslationKeysIgnored     int32 `json:"translation_keys_ignored,omitempty"`
}
