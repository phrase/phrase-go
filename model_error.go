package phrase

// ModelError Error field for when the request completely fails
type ModelError struct {
	Error ErrorError `json:"error,omitempty"`
}
