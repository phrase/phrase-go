package phrase

// FigmaAttachmentCreateParameters struct for FigmaAttachmentCreateParameters
type FigmaAttachmentCreateParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Figma file url
	Url string `json:"url,omitempty"`
}
