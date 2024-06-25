package phrase

// LocaleDownloadParams The parameters of the download request.
type LocaleDownloadParams struct {
	FileFormat                    string `json:"file_format,omitempty"`
	LocaleId                      string `json:"locale_id,omitempty"`
	Tags                          string `json:"tags,omitempty"`
	Branch                        string `json:"branch,omitempty"`
	IncludeEmptyTranslations      *bool  `json:"include_empty_translations,omitempty"`
	IncludeTranslatedKeys         *bool  `json:"include_translated_keys,omitempty"`
	IncludeUnverifiedTranslations *bool  `json:"include_unverified_translations,omitempty"`
}
