package phrase

// ReleaseUpdateParameters struct for ReleaseUpdateParameters
type ReleaseUpdateParameters struct {
	// Description of the release
	Description string `json:"description,omitempty"`
	// List of platforms the release should support.
	Platforms []string `json:"platforms,omitempty"`
	// Branch used for release
	Branch string `json:"branch,omitempty"`
}
