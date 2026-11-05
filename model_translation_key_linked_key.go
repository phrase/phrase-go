package phrase

// TranslationKeyLinkedKey Describes a key's role in a parent–child link relationship.
type TranslationKeyLinkedKey struct {
	// Whether this key is the parent or the child in the link relationship.
	Role   string                         `json:"role,omitempty"`
	Parent *TranslationKeyLinkedKeyParent `json:"parent,omitempty"`
	// The number of linked child keys. Set for parent keys.
	ChildrenCount *NullableInt32 `json:"children_count,omitempty"`
}
