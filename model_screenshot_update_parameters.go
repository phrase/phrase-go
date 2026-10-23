package phrase

import (
	"os"
)

// ScreenshotUpdateParameters struct for ScreenshotUpdateParameters
type ScreenshotUpdateParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Name of the screenshot
	Name string `json:"name,omitempty"`
	// Description of the screenshot. Send an empty string or `null` to clear the value; omit the field to leave it unchanged. The Phrase CLI drops `null` values before sending, so use an empty string there.
	Description *NullableString `json:"description,omitempty"`
	// Screenshot file
	Filename *os.File `json:"filename,omitempty"`
}
