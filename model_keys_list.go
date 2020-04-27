package phrase
// KeysList struct for KeysList
type KeysList struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Sort by field. Can be one of: name, created_at, updated_at.
	Sort string `json:"sort,omitempty"`
	// Order direction. Can be one of: asc, desc.
	Order string `json:"order,omitempty"`
	// q_description_placeholder
	Q string `json:"q,omitempty"`
	// Locale used to determine the translation state of a key when filtering for untranslated or translated keys.
	LocaleId string `json:"locale_id,omitempty"`
}
