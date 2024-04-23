package phrase

// GlossaryCreateParameters struct for GlossaryCreateParameters
type GlossaryCreateParameters struct {
	// Name of the glossary
	Name string `json:"name"`
	// List of project ids the glossary should be assigned to.
	ProjectIds string `json:"project_ids,omitempty"`
	// List of space ids the glossary should be assigned to.
	SpaceIds []string `json:"space_ids,omitempty"`
}
