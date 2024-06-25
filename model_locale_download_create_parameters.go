package phrase

// LocaleDownloadCreateParameters struct for LocaleDownloadCreateParameters
type LocaleDownloadCreateParameters struct {
	// File format name. See the <a href=\"https://support.phrase.com/hc/en-us/sections/6111343326364\">format guide</a> for all supported file formats.
	FileFormat string `json:"file_format"`
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Limit results to keys tagged with a list of comma separated tag names.
	Tags string `json:"tags,omitempty"`
	// Indicates whether keys without translations should be included in the output as well.
	IncludeEmptyTranslations *bool `json:"include_empty_translations,omitempty"`
	// Indicates whether zero forms should be included when empty in pluralized keys.
	ExcludeEmptyZeroForms *bool `json:"exclude_empty_zero_forms,omitempty"`
	// Include translated keys in the locale file. Use in combination with include_empty_translations to obtain only untranslated keys.
	IncludeTranslatedKeys *bool `json:"include_translated_keys,omitempty"`
	// Indicates whether [NOTRANSLATE] tags should be kept.
	KeepNotranslateTags *bool `json:"keep_notranslate_tags,omitempty"`
	// Additional formatting and render options. See the <a href=\"https://support.phrase.com/hc/en-us/sections/6111343326364\">format guide</a> for a list of options available for each format. Specify format options like this: <code>...&format_options[foo]=bar</code>
	FormatOptions map[string]interface{} `json:"format_options,omitempty"`
	// Enforces a specific encoding on the file contents. Valid options are \"UTF-8\", \"UTF-16\" and \"ISO-8859-1\".
	Encoding string `json:"encoding,omitempty"`
	// if set to false unverified translations are excluded
	IncludeUnverifiedTranslations *bool `json:"include_unverified_translations,omitempty"`
	// If set to true the last reviewed version of a translation is used. This is only available if the review workflow is enabled for the project.
	UseLastReviewedVersion *bool `json:"use_last_reviewed_version,omitempty"`
	// If a key has no translation in the locale being downloaded the translation in the fallback locale will be used. Provide the ID of the locale that should be used as the fallback. Requires include_empty_translations to be set to <code>true</code>.
	FallbackLocaleId string `json:"fallback_locale_id,omitempty"`
	// Provides the source language of a corresponding job as the source language of the generated locale file. This parameter will be ignored unless used in combination with a <code>tag</code> parameter indicating a specific job.
	SourceLocaleId string `json:"source_locale_id,omitempty"`
	// Custom metadata filters. Provide the name of the metadata field and the value to filter by. Only keys with matching metadata will be included in the download.
	CustomMetadataFilters map[string]interface{} `json:"custom_metadata_filters,omitempty"`
}
