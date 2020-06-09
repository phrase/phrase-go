package phrase
import (
	"time"
)
// GlossaryTerm struct for GlossaryTerm
type GlossaryTerm struct {
	Id string `json:"id,omitempty"`
	Term string `json:"term,omitempty"`
	Description string `json:"description,omitempty"`
	Translatable *bool `json:"translatable,omitempty"`
	CaseSensitive *bool `json:"case_sensitive,omitempty"`
	Translations []GlossaryTermTranslation `json:"translations,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
