package phrase
// KeysDelete struct for KeysDelete
type KeysDelete struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// q_description_placeholder
	Q string `json:"q,omitempty"`
	// Locale used to determine the translation state of a key when filtering for untranslated or translated keys.
	LocaleId string `json:"locale_id,omitempty"`
}
