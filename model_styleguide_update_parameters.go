package phrase

// StyleguideUpdateParameters struct for StyleguideUpdateParameters
type StyleguideUpdateParameters struct {
	// Style guide title
	Title string `json:"title,omitempty"`
	// Audience description. Send an empty string or `null` to clear the value; omit the field to leave it unchanged. The Phrase CLI drops `null` values before sending, so use an empty string there.
	Audience *NullableString `json:"audience,omitempty"`
	// Target audience for the translations.
	TargetAudience string `json:"target_audience,omitempty"`
	// Preferred grammatical person.
	GrammaticalPerson string `json:"grammatical_person,omitempty"`
	// Vocabulary register the translations should use.
	VocabularyType string `json:"vocabulary_type,omitempty"`
	// Description of the business. Send an empty string or `null` to clear the value; omit the field to leave it unchanged. The Phrase CLI drops `null` values before sending, so use an empty string there.
	Business *NullableString `json:"business,omitempty"`
	// Company branding to remain consistent. Send an empty string or `null` to clear the value; omit the field to leave it unchanged. The Phrase CLI drops `null` values before sending, so use an empty string there.
	CompanyBranding *NullableString `json:"company_branding,omitempty"`
	// Formatting requirements and character limitations. Send an empty string or `null` to clear the value; omit the field to leave it unchanged. The Phrase CLI drops `null` values before sending, so use an empty string there.
	Formatting *NullableString `json:"formatting,omitempty"`
	// List of terms and/or phrases that need to be translated consistently. Send an empty string or `null` to clear the value; omit the field to leave it unchanged. The Phrase CLI drops `null` values before sending, so use an empty string there.
	GlossaryTerms *NullableString `json:"glossary_terms,omitempty"`
	// Formal or informal pronouns, consistent conjugation, grammatical gender. Send an empty string or `null` to clear the value; omit the field to leave it unchanged. The Phrase CLI drops `null` values before sending, so use an empty string there.
	GrammarConsistency *NullableString `json:"grammar_consistency,omitempty"`
	// Can be one of: Cultural/Conversational, Literal, Neutral. Send an empty string or `null` to clear the value; omit the field to leave it unchanged. The Phrase CLI drops `null` values before sending, so use an empty string there.
	LiteralTranslation *NullableString `json:"literal_translation,omitempty"`
	// Tone requirement descriptions. Send an empty string or `null` to clear the value; omit the field to leave it unchanged. The Phrase CLI drops `null` values before sending, so use an empty string there.
	OverallTone *NullableString `json:"overall_tone,omitempty"`
	// Provide links to sample product pages, FAQ pages, etc. to give the translator a point of reference. You can also provide past translations. Even snippets or short paragraphs are helpful for maintaining consistency. Send an empty string or `null` to clear the value; omit the field to leave it unchanged. The Phrase CLI drops `null` values before sending, so use an empty string there.
	Samples *NullableString `json:"samples,omitempty"`
}
