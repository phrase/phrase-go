package phrase
// TagCreate struct for TagCreate
type TagCreate struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Name of the tag
	Name string `json:"name,omitempty"`
}
