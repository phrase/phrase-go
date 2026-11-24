package phrase

// ProjectNoteUpdateParameters struct for ProjectNoteUpdateParameters
type ProjectNoteUpdateParameters struct {
	// Markdown content for the project note (max 65,535 characters).
	Body string `json:"body"`
}
