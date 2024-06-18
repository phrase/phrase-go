package phrase

// ReleaseCreateParameters struct for ReleaseCreateParameters
type ReleaseCreateParameters struct {
	// Description of the release
	Description string `json:"description,omitempty"`
	// List of platforms the release should support.
	Platforms []string `json:"platforms,omitempty"`
	// List of locale ids that will be included in the release. If empty, distribution locales will be used
	LocaleIds []string `json:"locale_ids,omitempty"`
	// Only include tagged keys in the release. For a key to be included it must be tagged with all tags provided
	Tags []string `json:"tags,omitempty"`
	// Minimum version of the app that the release supports in semver format
	AppMinVersion string `json:"app_min_version,omitempty"`
	// Maximum version of the app that the release supports in semver format
	AppMaxVersion string `json:"app_max_version,omitempty"`
	// Branch used for release
	Branch string `json:"branch,omitempty"`
}
