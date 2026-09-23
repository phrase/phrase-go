package phrase

// OrganizationJobTemplateUpdateParameters struct for OrganizationJobTemplateUpdateParameters
type OrganizationJobTemplateUpdateParameters struct {
	// Job template name
	Name string `json:"name"`
	// Briefing for the translators
	Briefing string `json:"briefing,omitempty"`
	// Automatically translate the job using machine translation.
	Autotranslate *bool `json:"autotranslate,omitempty"`
	// Code of the account member to set as the job template owner. Pass an empty string to clear a previously set owner; when blank, the owner is set to null and jobs created from this template will default to assigning the job creator as owner.
	OwnerId string `json:"owner_id,omitempty"`
}
