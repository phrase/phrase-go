package phrase

// KeyLinksBatchDestroyParameters struct for KeyLinksBatchDestroyParameters
type KeyLinksBatchDestroyParameters struct {
	// The IDs of the child keys to unlink from the parent key.
	ChildKeyIds []string `json:"child_key_ids"`
}
