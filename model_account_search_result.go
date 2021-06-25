package phrase

// AccountSearchResult struct for AccountSearchResult
type AccountSearchResult struct {
	Query             string        `json:"query,omitempty"`
	Excerpt           string        `json:"excerpt,omitempty"`
	Key               KeyPreview    `json:"key,omitempty"`
	Locale            LocalePreview `json:"locale,omitempty"`
	Project           Project       `json:"project,omitempty"`
	Translation       Translation   `json:"translation,omitempty"`
	OtherTranslations []Translation `json:"other_translations,omitempty"`
}
