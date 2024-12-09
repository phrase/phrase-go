package phrase

// JobCommentUpdateParameters struct for JobCommentUpdateParameters
type JobCommentUpdateParameters struct {
	// Comment message
	Message string `json:"message,omitempty"`
	// Branch name for the job
	Branch string `json:"branch,omitempty"`
}
