package phrase

// DistributionUpdateParameters struct for DistributionUpdateParameters
type DistributionUpdateParameters struct {
	// Name of the distribution
	Name string `json:"name,omitempty"`
	// Project id the distribution should be assigned to.
	ProjectId string `json:"project_id,omitempty"`
	// List of platforms the distribution should support.
	Platforms []string `json:"platforms,omitempty"`
	// List of locale ids that will be part of distribution releases
	LocaleIds []string `json:"locale_ids,omitempty"`
	// Additional formatting and render options. Only <code>enclose_in_cdata</code> is available for platform <code>android</code>.
	FormatOptions map[string]string `json:"format_options,omitempty"`
	// Indicates whether to fallback to non regional locale when locale can not be found
	FallbackToNonRegionalLocale *bool `json:"fallback_to_non_regional_locale,omitempty"`
	// Indicates whether to fallback to projects default locale when locale can not be found
	FallbackToDefaultLocale *bool `json:"fallback_to_default_locale,omitempty"`
	// Use last reviewed instead of latest translation in a project
	UseLastReviewedVersion *bool `json:"use_last_reviewed_version,omitempty"`
}
