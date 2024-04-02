package phrase

// KeyLinksCreateParameters struct for KeyLinksCreateParameters
type KeyLinksCreateParameters struct {
	// The IDs of the child keys to link to the parent key. Can be left empty, to only mark the given translation-key as parent
	ChildKeyIds []string `json:"child_key_ids"`
}
