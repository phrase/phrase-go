package phrase

// OrganizationJobTemplateCreateParameters struct for OrganizationJobTemplateCreateParameters
type OrganizationJobTemplateCreateParameters struct {
	// Job template name
	Name string `json:"name"`
	// Briefing for the translators
	Briefing string `json:"briefing,omitempty"`
	// Automatically translate the job using machine translation.
	Autotranslate *bool `json:"autotranslate,omitempty"`
	// Exclude linked child keys from jobs created from this template. Defaults to `false`.
	ExcludeLinkedChildKeys *bool `json:"exclude_linked_child_keys,omitempty"`
	// Code of the account member to set as the job template owner. When omitted or blank, no owner is pre-set; the user who creates a job from this template is assigned as its owner at job-creation time.
	OwnerId string `json:"owner_id,omitempty"`
}
