package phrase
// TagCreateParameters struct for TagCreateParameters
type TagCreateParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Name of the tag
	Name string `json:"name,omitempty"`
}
