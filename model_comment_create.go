package phrase
// CommentCreate struct for CommentCreate
type CommentCreate struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Comment message
	Message string `json:"message,omitempty"`
}
