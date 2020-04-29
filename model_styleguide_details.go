package phrase
import (
	"time"
)
// StyleguideDetails struct for StyleguideDetails
type StyleguideDetails struct {
	Id string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	PublicUrl string `json:"public_url,omitempty"`
	Audience string `json:"audience,omitempty"`
	TargetAudience string `json:"target_audience,omitempty"`
	GrammaticalPerson string `json:"grammatical_person,omitempty"`
	VocabularyType string `json:"vocabulary_type,omitempty"`
	Business string `json:"business,omitempty"`
	CompanyBranding string `json:"company_branding,omitempty"`
	Formatting string `json:"formatting,omitempty"`
	GlossaryTerms string `json:"glossary_terms,omitempty"`
	GrammarConsistency string `json:"grammar_consistency,omitempty"`
	LiteralTranslation string `json:"literal_translation,omitempty"`
	OverallTone string `json:"overall_tone,omitempty"`
	Samples string `json:"samples,omitempty"`
}
