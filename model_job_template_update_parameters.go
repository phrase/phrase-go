package phrase

// JobTemplateUpdateParameters struct for JobTemplateUpdateParameters
type JobTemplateUpdateParameters struct {
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
	// Code of the account member to set as the job template owner. The referenced user must also be a member of the project; passing the code of an account member who is not a project member returns a 404. Pass an empty string to clear a previously set owner; when blank, jobs created from this template will default to assigning the job creator as owner.
	OwnerId string `json:"owner_id,omitempty"`
}
