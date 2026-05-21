package phrase

// StyleguideCreateParameters struct for StyleguideCreateParameters
type StyleguideCreateParameters struct {
	// Style guide title
	Title string `json:"title"`
	// Audience description
	Audience string `json:"audience,omitempty"`
	// Target audience for the translations.
	TargetAudience string `json:"target_audience,omitempty"`
	// Preferred grammatical person.
	GrammaticalPerson string `json:"grammatical_person,omitempty"`
	// Vocabulary register the translations should use.
	VocabularyType string `json:"vocabulary_type,omitempty"`
	// Description of the business
	Business string `json:"business,omitempty"`
	// Company branding to remain consistent.
	CompanyBranding string `json:"company_branding,omitempty"`
	// Formatting requirements and character limitations.
	Formatting string `json:"formatting,omitempty"`
	// List of terms and/or phrases that need to be translated consistently.
	GlossaryTerms string `json:"glossary_terms,omitempty"`
	// Formal or informal pronouns, consistent conjugation, grammatical gender
	GrammarConsistency string `json:"grammar_consistency,omitempty"`
	// Can be one of: Cultural/Conversational, Literal, Neutral.
	LiteralTranslation string `json:"literal_translation,omitempty"`
	// Tone requirement descriptions
	OverallTone string `json:"overall_tone,omitempty"`
	// Provide links to sample product pages, FAQ pages, etc. to give the translator a point of reference. You can also provide past translations. Even snippets or short paragraphs are helpful for maintaining consistency.
	Samples string `json:"samples,omitempty"`
}
