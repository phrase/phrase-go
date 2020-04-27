package phrase
import (
	"os"
)
// ScreenshotCreate struct for ScreenshotCreate
type ScreenshotCreate struct {
	// Name of the screenshot
	Name string `json:"name,omitempty"`
	// Description of the screenshot
	Description string `json:"description,omitempty"`
	// Screenshot file
	Filename *os.File `json:"filename,omitempty"`
}
