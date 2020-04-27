package phrase
import (
	"time"
)
// ScreenshotMarker struct for ScreenshotMarker
type ScreenshotMarker struct {
	Id string `json:"id,omitempty"`
	Presentation string `json:"presentation,omitempty"`
	PresentationType string `json:"presentation_type,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	TranslationKey map[string]interface{} `json:"translation_key,omitempty"`
}
