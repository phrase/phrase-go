package phrase

// CommentCreateParameters struct for CommentCreateParameters
type CommentCreateParameters struct {
	// Comment message
	Message string `json:"message"`
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// specify the locales for the comment
	LocaleIds []string `json:"locale_ids,omitempty"`
}
