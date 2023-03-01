package phrase

// OrganizationJobTemplateLocalesCreateParameters struct for OrganizationJobTemplateLocalesCreateParameters
type OrganizationJobTemplateLocalesCreateParameters struct {
	// locale name
	LocaleName string `json:"locale_name"`
	// locale code
	LocaleCode string `json:"locale_code"`
	// Array of user ids to be assigned to the job template locale
	UserIds []string `json:"user_ids,omitempty"`
	// Array of reviewer ids to be assigned to the job template locale
	ReviewerIds []string `json:"reviewer_ids,omitempty"`
	// Array of team ids to be assigned to the job locale as translators
	TranslatorTeamIds []string `json:"translator_team_ids,omitempty"`
	// Array of team ids to be assigned to the job locale as reviewers
	ReviewerTeamIds []string `json:"reviewer_team_ids,omitempty"`
}
