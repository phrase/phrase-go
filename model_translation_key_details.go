package phrase
import (
	"time"
)
// TranslationKeyDetails struct for TranslationKeyDetails
type TranslationKeyDetails struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	NameHash string `json:"name_hash,omitempty"`
	Plural *bool `json:"plural,omitempty"`
	Tags []string `json:"tags,omitempty"`
	DataType string `json:"data_type,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	NamePlural string `json:"name_plural,omitempty"`
	CommentsCount int32 `json:"comments_count,omitempty"`
	MaxCharactersAllowed int32 `json:"max_characters_allowed,omitempty"`
	ScreenshotUrl string `json:"screenshot_url,omitempty"`
	Unformatted *bool `json:"unformatted,omitempty"`
	XmlSpacePreserve *bool `json:"xml_space_preserve,omitempty"`
	OriginalFile string `json:"original_file,omitempty"`
	FormatValueType string `json:"format_value_type,omitempty"`
	Creator UserPreview `json:"creator,omitempty"`
}
