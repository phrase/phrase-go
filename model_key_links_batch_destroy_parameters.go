package phrase

// KeyLinksBatchDestroyParameters struct for KeyLinksBatchDestroyParameters
type KeyLinksBatchDestroyParameters struct {
	// Codes of the child keys to unlink. Required when unlink_parent is false or omitted. Ignored when unlink_parent is true.
	ChildKeyIds []string `json:"child_key_ids"`
	// When true, dissolves the entire linked-key group by unlinking all children and removing the group. The child_key_ids field is ignored when this is set to true.
	UnlinkParent *bool `json:"unlink_parent,omitempty"`
	// Controls what happens to child key translation content after unlinking. keep_content (default) copies the parent translation into each child; remove_content clears each child translation.
	Strategy string `json:"strategy,omitempty"`
}
