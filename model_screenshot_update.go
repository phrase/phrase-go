package phrase
// ScreenshotUpdate struct for ScreenshotUpdate
type ScreenshotUpdate struct {
	// Name of the screenshot
	Name string `json:"name,omitempty"`
	// Description of the screenshot
	Description string `json:"description,omitempty"`
	// Screenshot file
	Filename string `json:"filename,omitempty"`
}
