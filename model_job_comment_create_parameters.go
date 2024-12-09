package phrase

// JobCommentCreateParameters struct for JobCommentCreateParameters
type JobCommentCreateParameters struct {
	// Job comment message
	Message string `json:"message,omitempty"`
	// Branch name for the job
	Branch string `json:"branch,omitempty"`
}
