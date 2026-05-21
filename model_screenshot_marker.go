package phrase

import (
	"time"
)

// ScreenshotMarker struct for ScreenshotMarker
type ScreenshotMarker struct {
	Id           string                       `json:"id,omitempty"`
	Presentation ScreenshotMarkerPresentation `json:"presentation,omitempty"`
	// Marker presentation style. The default value is `default`.
	PresentationType string         `json:"presentation_type,omitempty"`
	CreatedAt        time.Time      `json:"created_at,omitempty"`
	UpdatedAt        time.Time      `json:"updated_at,omitempty"`
	TranslationKey   TranslationKey `json:"translation_key,omitempty"`
}
