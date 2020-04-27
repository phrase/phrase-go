package phrase
// KeysTag struct for KeysTag
type KeysTag struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// q_description_placeholder
	Q string `json:"q,omitempty"`
	// Locale used to determine the translation state of a key when filtering for untranslated or translated keys.
	LocaleId string `json:"locale_id,omitempty"`
	// Tag or comma-separated list of tags to add to the matching collection of keys
	Tags string `json:"tags,omitempty"`
}
