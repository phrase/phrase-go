package phrase
// CommentUpdate struct for CommentUpdate
type CommentUpdate struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Comment message
	Message string `json:"message,omitempty"`
}
