package phrase
// CommentCreateParameters struct for CommentCreateParameters
type CommentCreateParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Comment message
	Message string `json:"message,omitempty"`
}
