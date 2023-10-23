package phrase

// RepliesListParameters struct for RepliesListParameters
type RepliesListParameters struct {
	// Specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Search query for comment messages
	Query string `json:"query,omitempty"`
	// Specify filters to find comments by
	Filters []string `json:"filters,omitempty"`
	// Specify ordering of comments
	Order string `json:"order,omitempty"`
}
