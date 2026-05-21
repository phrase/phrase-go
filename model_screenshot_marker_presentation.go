package phrase

// ScreenshotMarkerPresentation Coordinates and size of the marker on the screenshot, in pixels.
type ScreenshotMarkerPresentation struct {
	X int32 `json:"x,omitempty"`
	Y int32 `json:"y,omitempty"`
	W int32 `json:"w,omitempty"`
	H int32 `json:"h,omitempty"`
}
