package phrase
// CommentUpdateParameters struct for CommentUpdateParameters
type CommentUpdateParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Comment message
	Message string `json:"message,omitempty"`
}
