package phrase

// ReleaseCreateParameters struct for ReleaseCreateParameters
type ReleaseCreateParameters struct {
	// Description of the release
	Description string `json:"description,omitempty"`
	// List of platforms the release should support.
	Platforms []string `json:"platforms,omitempty"`
	// List of locale ids that will be included in the release. If empty, distribution locales will be used
	LocaleIds []string `json:"locale_ids,omitempty"`
	// Branch used for release
	Branch string `json:"branch,omitempty"`
}
