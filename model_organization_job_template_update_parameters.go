package phrase

// OrganizationJobTemplateUpdateParameters struct for OrganizationJobTemplateUpdateParameters
type OrganizationJobTemplateUpdateParameters struct {
	// Job template name
	Name string `json:"name"`
	// Briefing for the translators
	Briefing string `json:"briefing,omitempty"`
	// Automatically translate the job using machine translation.
	Autotranslate *bool `json:"autotranslate,omitempty"`
}
