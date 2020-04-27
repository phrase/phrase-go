package phrase
// GlossaryUpdate struct for GlossaryUpdate
type GlossaryUpdate struct {
	// Name of the glossary
	Name string `json:"name,omitempty"`
	// List of project ids the glossary should be assigned to.
	ProjectIds string `json:"project_ids,omitempty"`
	// List of space ids the glossary should be assigned to.
	SpaceIds string `json:"space_ids,omitempty"`
}
