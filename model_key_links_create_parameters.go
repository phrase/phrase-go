package phrase

// KeyLinksCreateParameters struct for KeyLinksCreateParameters
type KeyLinksCreateParameters struct {
	// Codes of the keys to link as children. Pass an empty array to mark the parent key without linking any children.
	ChildKeyIds []string `json:"child_key_ids"`
}
