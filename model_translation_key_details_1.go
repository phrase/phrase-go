package phrase
// TranslationKeyDetails1 struct for TranslationKeyDetails1
type TranslationKeyDetails1 struct {
	NamePlural string `json:"name_plural,omitempty"`
	CommentsCount int32 `json:"comments_count,omitempty"`
	MaxCharactersAllowed int32 `json:"max_characters_allowed,omitempty"`
	ScreenshotUrl string `json:"screenshot_url,omitempty"`
	Unformatted *bool `json:"unformatted,omitempty"`
	XmlSpacePreserve *bool `json:"xml_space_preserve,omitempty"`
	OriginalFile string `json:"original_file,omitempty"`
	FormatValueType string `json:"format_value_type,omitempty"`
	Creator UserPreview `json:"creator,omitempty"`
}
