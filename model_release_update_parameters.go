package phrase

// ReleaseUpdateParameters struct for ReleaseUpdateParameters
type ReleaseUpdateParameters struct {
	// Description of the release
	Description string `json:"description,omitempty"`
	// List of platforms the release should support.
	Platforms []string `json:"platforms,omitempty"`
	// Minimum version of the app that the release supports in semver format
	AppMinVersion string `json:"app_min_version,omitempty"`
	// Maximum version of the app that the release supports in semver format
	AppMaxVersion string `json:"app_max_version,omitempty"`
	// Branch used for release
	Branch string `json:"branch,omitempty"`
}
