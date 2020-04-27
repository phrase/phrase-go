package phrase
// GlossaryTermCreate struct for GlossaryTermCreate
type GlossaryTermCreate struct {
	// Glossary term
	Term string `json:"term,omitempty"`
	// Description of term
	Description string `json:"description,omitempty"`
	// Indicates whether the term should be used for all languages or can be translated
	Translatable string `json:"translatable,omitempty"`
	// Indicates whether the term is case sensitive
	CaseSensitive string `json:"case_sensitive,omitempty"`
}
