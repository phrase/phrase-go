package phrase
import (
	"time"
)
// TranslationKey struct for TranslationKey
type TranslationKey struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	NameHash string `json:"name_hash,omitempty"`
	Plural bool `json:"plural,omitempty"`
	Tags []string `json:"tags,omitempty"`
	DataType string `json:"data_type,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
