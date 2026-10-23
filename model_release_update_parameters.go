package phrase

// ReleaseUpdateParameters struct for ReleaseUpdateParameters
type ReleaseUpdateParameters struct {
	// Description of the release. Send an empty string or `null` to clear the value; omit the field to leave it unchanged. The Phrase CLI drops `null` values before sending, so use an empty string there.
	Description *NullableString `json:"description,omitempty"`
	// List of platforms the release should support.
	Platforms []string `json:"platforms,omitempty"`
	// Minimum version of the app that the release supports in semver format. Send an empty string or `null` to clear the value; omit the field to leave it unchanged. The Phrase CLI drops `null` values before sending, so use an empty string there.
	AppMinVersion *NullableString `json:"app_min_version,omitempty"`
	// Maximum version of the app that the release supports in semver format. Send an empty string or `null` to clear the value; omit the field to leave it unchanged. The Phrase CLI drops `null` values before sending, so use an empty string there.
	AppMaxVersion *NullableString `json:"app_max_version,omitempty"`
	// Branch used for release
	Branch string `json:"branch,omitempty"`
}
