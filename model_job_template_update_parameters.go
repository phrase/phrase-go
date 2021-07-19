package phrase

// JobTemplateUpdateParameters struct for JobTemplateUpdateParameters
type JobTemplateUpdateParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Job template name
	Name string `json:"name"`
	// Briefing for the translators
	Briefing string `json:"briefing,omitempty"`
}
