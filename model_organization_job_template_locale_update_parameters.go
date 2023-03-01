package phrase

// OrganizationJobTemplateLocaleUpdateParameters struct for OrganizationJobTemplateLocaleUpdateParameters
type OrganizationJobTemplateLocaleUpdateParameters struct {
	// locale name
	LocaleName string `json:"locale_name,omitempty"`
	// locale code
	LocaleCode string `json:"locale_code,omitempty"`
	// Array of user ids to be assigned to the job template locale
	UserIds []string `json:"user_ids,omitempty"`
	// Array of reviewer ids to be assigned to the job template locale
	ReviewerIds []string `json:"reviewer_ids,omitempty"`
	// Array of team ids to be assigned to the job locale as translators
	TranslatorTeamIds []string `json:"translator_team_ids,omitempty"`
	// Array of team ids to be assigned to the job locale as reviewers
	ReviewerTeamIds []string `json:"reviewer_team_ids,omitempty"`
}
