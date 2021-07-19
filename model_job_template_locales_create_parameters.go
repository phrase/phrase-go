package phrase

// JobTemplateLocalesCreateParameters struct for JobTemplateLocalesCreateParameters
type JobTemplateLocalesCreateParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// locale id
	LocaleId string `json:"locale_id"`
	// Array of user ids to be assigned to the job template locale
	UserIds []string `json:"user_ids,omitempty"`
	// Array of reviewer ids to be assigned to the job template locale
	ReviewerIds []string `json:"reviewer_ids,omitempty"`
}
