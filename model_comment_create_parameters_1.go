package phrase

// CommentCreateParameters1 struct for CommentCreateParameters1
type CommentCreateParameters1 struct {
	// Reply message body
	Message string `json:"message"`
	// Specify the branch to use
	Branch string `json:"branch,omitempty"`
}
