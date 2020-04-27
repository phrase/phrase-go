package phrase
import (
	"time"
)
// Translation struct for Translation
type Translation struct {
	Id string `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
	Unverified bool `json:"unverified,omitempty"`
	Excluded bool `json:"excluded,omitempty"`
	PluralSuffix string `json:"plural_suffix,omitempty"`
	Key KeyPreview `json:"key,omitempty"`
	Locale map[string]interface{} `json:"locale,omitempty"`
	Placeholders []string `json:"placeholders,omitempty"`
	State string `json:"state,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
