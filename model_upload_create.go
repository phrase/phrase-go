package phrase
// UploadCreate struct for UploadCreate
type UploadCreate struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// File to be imported
	File string `json:"file,omitempty"`
	// File format. Auto-detected when possible and not specified.
	FileFormat string `json:"file_format,omitempty"`
	// Locale of the file's content. Can be the name or public id of the locale. Preferred is the public id.
	LocaleId string `json:"locale_id,omitempty"`
	// List of tags separated by comma to be associated with the new keys contained in the upload.
	Tags string `json:"tags,omitempty"`
	// Indicates whether existing translations should be updated with the file content.
	UpdateTranslations string `json:"update_translations,omitempty"`
	// Existing key descriptions will be updated with the file content. Empty descriptions overwrite existing descriptions.
	UpdateDescriptions string `json:"update_descriptions,omitempty"`
	// This option is obsolete. Providing the option will cause a bad request error.
	ConvertEmoji string `json:"convert_emoji,omitempty"`
	// Indicates whether the upload should not create upload tags.
	SkipUploadTags string `json:"skip_upload_tags,omitempty"`
	// Indicates whether the upload should unverify updated translations.
	SkipUnverification string `json:"skip_unverification,omitempty"`
	// Enforces a specific encoding on the file contents. Valid options are \"UTF-8\", \"UTF-16\" and \"ISO-8859-1\".
	FileEncoding string `json:"file_encoding,omitempty"`
	// Optional, format specific mapping between locale names and the columns the translations to those locales are contained in.
	LocaleMapping string `json:"locale_mapping,omitempty"`
	// Additional options available for specific formats. See our format guide for complete list.
	FormatOptions string `json:"format_options,omitempty"`
	// If set, translations for the uploaded language will be fetched automatically.
	Autotranslate string `json:"autotranslate,omitempty"`
	// Indicated whether the imported translations should be marked as reviewed. This setting is available if the review workflow (currently beta) is enabled for the project.
	MarkReviewed string `json:"mark_reviewed,omitempty"`
}
