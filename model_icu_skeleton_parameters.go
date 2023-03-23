package phrase

// IcuSkeletonParameters struct for IcuSkeletonParameters
type IcuSkeletonParameters struct {
	// Source content
	Content string `json:"content,omitempty"`
	// Locale codes
	LocaleCodes []string `json:"locale_codes,omitempty"`
	// Keep the content and add missing plural forms for each locale
	KeepContent *bool `json:"keep_content,omitempty"`
	// Indicates whether the zero form should be included or excluded in the returned skeletons
	ZeroFormEnabled *bool `json:"zero_form_enabled,omitempty"`
	// Strings supports two CLDR variants, when it comes to pluralization rules. \\ You can choose which one you want to use when constructing the skeletons. Possible values \\ are `legacy` and `cldr_41`. Default value is `legacy`.
	CldrVersion string `json:"cldr_version,omitempty"`
}
