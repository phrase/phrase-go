package phrase

// KeysSearchParameters struct for KeysSearchParameters
type KeysSearchParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Sort by field. Can be one of: name, created_at, updated_at.
	Sort string `json:"sort,omitempty"`
	// Order direction. Can be one of: asc, desc.
	Order string `json:"order,omitempty"`
	// Specify a query to do broad search for keys by name (including wildcards).  The following qualifiers are also supported in the search term:  * `ids:key_id,...` for queries on a comma-separated list of ids * `name:key_name,...` for text queries on a comma-seperated list of exact key names - spaces, commas, and colons need to be escaped with double backslashes * `tags:tag_name,...` to filter for keys with certain comma-seperated list of tags * `uploads:upload_id,...` to filter for keys with certain comma-seperated list of uploads * `job:{true|false}` to filter for keys mentioned in an active job * `translated:{true|false}` for translation status (also requires `locale_id` to be specified) * `updated_at:{>=|<=}2013-02-21T00:00:00Z` for date range queries * `unmentioned_in_upload:upload_id,...` to filter keys unmentioned within upload. When multiple upload IDs provided, matches only keys not mentioned in **all** uploads  Find more examples [here](/en/api/strings/usage-examples). Please note: If `tags` are added to filter the search, the search will be limited to a maximum of 65,536 tagged keys.
	Q string `json:"q,omitempty"`
	// Locale used to determine the translation state of a key when filtering for untranslated or translated keys.
	LocaleId string `json:"locale_id,omitempty"`
}
