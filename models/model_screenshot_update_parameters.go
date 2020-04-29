package phrase
import (
	"os"
)
// ScreenshotUpdateParameters struct for ScreenshotUpdateParameters
type ScreenshotUpdateParameters struct {
	// Name of the screenshot
	Name string `json:"name,omitempty"`
	// Description of the screenshot
	Description string `json:"description,omitempty"`
	// Screenshot file
	Filename *os.File `json:"filename,omitempty"`
}
