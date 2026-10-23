package phrase

// GlossaryTermUpdateParameters struct for GlossaryTermUpdateParameters
type GlossaryTermUpdateParameters struct {
	// Glossary term
	Term string `json:"term,omitempty"`
	// Description of term. Send an empty string or `null` to clear the value; omit the field to leave it unchanged. The Phrase CLI drops `null` values before sending, so use an empty string there.
	Description *NullableString `json:"description,omitempty"`
	// Indicates whether the term should be used for all languages or can be translated
	Translatable *bool `json:"translatable,omitempty"`
	// Indicates whether the term is case sensitive
	CaseSensitive *bool `json:"case_sensitive,omitempty"`
}
