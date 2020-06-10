package phrase

// TranslationsVerifyParameters struct for TranslationsVerifyParameters
type TranslationsVerifyParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// q_description_placeholder
	Q string `json:"q,omitempty"`
	// Sort criteria. Can be one of: key_name, created_at, updated_at.
	Sort string `json:"sort,omitempty"`
	// Order direction. Can be one of: asc, desc.
	Order string `json:"order,omitempty"`
}
