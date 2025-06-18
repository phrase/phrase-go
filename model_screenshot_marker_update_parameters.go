package phrase

// ScreenshotMarkerUpdateParameters struct for ScreenshotMarkerUpdateParameters
type ScreenshotMarkerUpdateParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Specify the Key ID which should be highlighted on the specified screenshot. The Key must belong to the project.
	KeyId string `json:"key_id,omitempty"`
	// Presentation details of the screenshot marker in JSON format.  Each Screenshot Marker is represented as a rectangular shaped highlight box with the name of the specified Key attached. You can specify the marker position on the screenshot (`x`-axis and `y`-axis in pixels) from the top left corner of the screenshot and the dimensions of the marker itself (`w` and `h` in pixels).
	Presentation string `json:"presentation,omitempty"`
}
