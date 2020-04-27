package phrase
// LocaleDownload struct for LocaleDownload
type LocaleDownload struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// File format name. See the format guide for all supported file formats.
	FileFormat string `json:"file_format,omitempty"`
	// Limit results to keys tagged with a list of comma separated tag names.
	Tags string `json:"tags,omitempty"`
	// Limit download to tagged keys. This parameter is deprecated. Please use the \"tags\" parameter instead
	Tag string `json:"tag,omitempty"`
	// Indicates whether keys without translations should be included in the output as well.
	IncludeEmptyTranslations string `json:"include_empty_translations,omitempty"`
	// Include translated keys in the locale file. Use in combination with include_empty_translations to obtain only untranslated keys.
	IncludeTranslatedKeys string `json:"include_translated_keys,omitempty"`
	// Indicates whether [NOTRANSLATE] tags should be kept.
	KeepNotranslateTags string `json:"keep_notranslate_tags,omitempty"`
	// This option is obsolete. Projects that were created on or after Nov 29th 2019 or that did not contain emoji by then will not require this flag any longer since emoji are now supported natively.
	ConvertEmoji string `json:"convert_emoji,omitempty"`
	// Additional formatting and render options. See the <a href=\"https://help.phrase.com/help/supported-platforms-and-formats\">format guide</a> for a list of options available for each format. Specify format options like this: <code>...&format_options[foo]=bar</code>
	FormatOptions string `json:"format_options,omitempty"`
	// Enforces a specific encoding on the file contents. Valid options are \"UTF-8\", \"UTF-16\" and \"ISO-8859-1\".
	Encoding string `json:"encoding,omitempty"`
	// Indicates whether the locale file should skip all unverified translations. This parameter is deprecated and should be replaced with <code>include_unverified_translations</code>.
	SkipUnverifiedTranslations string `json:"skip_unverified_translations,omitempty"`
	// if set to false unverified translations are excluded
	IncludeUnverifiedTranslations string `json:"include_unverified_translations,omitempty"`
	// If set to true the last reviewed version of a translation is used. This is only available if the review workflow (currently in beta) is enabled for the project.
	UseLastReviewedVersion string `json:"use_last_reviewed_version,omitempty"`
	// If a key has no translation in the locale being downloaded the translation in the fallback locale will be used. Provide the public ID of the locale that should be used as the fallback. Requires include_empty_translations to be set to <code>true</code>.
	FallbackLocaleId string `json:"fallback_locale_id,omitempty"`
}
