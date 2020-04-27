package phrase
import (
	"time"
)
// TranslationVersion struct for TranslationVersion
type TranslationVersion struct {
	Id string `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
	PluralSuffix string `json:"plural_suffix,omitempty"`
	Key map[string]interface{} `json:"key,omitempty"`
	Locale map[string]interface{} `json:"locale,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	ChangedAt time.Time `json:"changed_at,omitempty"`
}
