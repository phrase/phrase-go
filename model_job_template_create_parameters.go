package phrase

// JobTemplateCreateParameters struct for JobTemplateCreateParameters
type JobTemplateCreateParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Job template name
	Name string `json:"name"`
	// Briefing for the translators
	Briefing string `json:"briefing,omitempty"`
	// Automatically translate the job using machine translation.
	Autotranslate *bool `json:"autotranslate,omitempty"`
	// The API id of the source language. This locale will be set as source locale for the job template. If not provided, the project default locale will be used.
	SourceLocaleId string `json:"source_locale_id,omitempty"`
	// Code of the account member to set as the job template owner. The referenced user must also be a member of the project; passing the code of an account member who is not a project member returns a 404. When omitted or blank, no owner is pre-set; the user who creates a job from this template is assigned as owner at job-creation time.
	OwnerId string `json:"owner_id,omitempty"`
}
