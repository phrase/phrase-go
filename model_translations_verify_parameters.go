package phrase

// TranslationsVerifyParameters struct for TranslationsVerifyParameters
type TranslationsVerifyParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// specify the locale of the translations to be verified
	LocaleId string `json:"locale_id,omitempty"`
	// Specify a query to find translations by content (including wildcards).  *Note: Search is limited to 10000 results and may not include recently updated data (depending on the project size).*  The following qualifiers are supported in the query: * `id:translation_id,...` for queries on a comma-separated list of ids * `tags:XYZ` for tags on the translation * `unverified:{true|false}` for verification status * `excluded:{true|false}` for exclusion status * `updated_at:{>=|<=}2013-02-21T00:00:00Z` for date range queries  Find more examples [here](/en/api/strings/usage-examples).
	Q string `json:"q,omitempty"`
}
