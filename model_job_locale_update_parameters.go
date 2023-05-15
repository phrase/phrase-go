package phrase

// JobLocaleUpdateParameters struct for JobLocaleUpdateParameters
type JobLocaleUpdateParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// ID of a target locale to update
	LocaleId string `json:"locale_id,omitempty"`
	// Array of user ids to be assigned to the job locale
	UserIds []string `json:"user_ids,omitempty"`
	// Array of reviewer ids to be assigned to the job locale as reviewers
	ReviewerIds []string `json:"reviewer_ids,omitempty"`
	// Array of team ids to be assigned to the job locale as translators
	TranslatorTeamIds []string `json:"translator_team_ids,omitempty"`
	// Array of team ids to be assigned to the job locale as reviewers
	ReviewerTeamIds []string `json:"reviewer_team_ids,omitempty"`
}
