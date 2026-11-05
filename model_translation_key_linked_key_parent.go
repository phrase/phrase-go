package phrase

// TranslationKeyLinkedKeyParent The parent key's public identifier and name. Set for child keys.
type TranslationKeyLinkedKeyParent struct {
	// The parent key's public identifier (code).
	Id string `json:"id,omitempty"`
	// The parent key's name.
	Name string `json:"name,omitempty"`
}
