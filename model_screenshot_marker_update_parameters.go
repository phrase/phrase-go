package phrase

// ScreenshotMarkerUpdateParameters struct for ScreenshotMarkerUpdateParameters
type ScreenshotMarkerUpdateParameters struct {
	// Specify the Key ID which should be highlighted on the specified screenshot. The Key must belong to the project.
	KeyId string `json:"key_id,omitempty"`
	// Presentation details of the screenshot marker in JSON format.<br/><br/>Each Screenshot Marker is represented as a rectangular shaped highlight box with the name of the specified Key attached. You can specify the marker position on the screenshot (<code>x</code>-axis and <code>y</code>-axis in pixels) from the top left corner of the screenshot and the dimensions of the marker itself (<code>w</code> and <code>h</code> in pixels).
	Presentation string `json:"presentation,omitempty"`
}
