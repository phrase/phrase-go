package phrase

// CommentsListParameters struct for CommentsListParameters
type CommentsListParameters struct {
	// Specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Search query for comment messages
	Query string `json:"query,omitempty"`
	// Search comments by their assigned locales
	LocaleIds []string `json:"locale_ids,omitempty"`
	// Specify filters to find comments by
	Filters []string `json:"filters,omitempty"`
	// Specify ordering of comments
	Order string `json:"order,omitempty"`
}
