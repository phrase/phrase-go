package phrase

// TranslationsUnverifyParameters struct for TranslationsUnverifyParameters
type TranslationsUnverifyParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Specify a query to find translations by content (including wildcards).  *Note: Search is limited to 10000 results and may not include recently updated data (depending on the project size).*  The following qualifiers are supported in the query: * `id:translation_id,...` for queries on a comma-separated list of ids * `tags:XYZ` for tags on the translation * `unverified:{true|false}` for verification status * `excluded:{true|false}` for exclusion status * `updated_at:{>=|<=}2013-02-21T00:00:00Z` for date range queries  Find more examples [here](/en/api/strings/usage-examples).
	Q string `json:"q,omitempty"`
	// Sort criteria. Can be one of: key_name, created_at, updated_at.
	Sort string `json:"sort,omitempty"`
	// Order direction. Can be one of: asc, desc.
	Order string `json:"order,omitempty"`
}
