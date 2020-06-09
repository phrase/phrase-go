package phrase
// GlossaryTermCreateParameters struct for GlossaryTermCreateParameters
type GlossaryTermCreateParameters struct {
	// Glossary term
	Term string `json:"term,omitempty"`
	// Description of term
	Description string `json:"description,omitempty"`
	// Indicates whether the term should be used for all languages or can be translated
	Translatable *bool `json:"translatable,omitempty"`
	// Indicates whether the term is case sensitive
	CaseSensitive *bool `json:"case_sensitive,omitempty"`
}
