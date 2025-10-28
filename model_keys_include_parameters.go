package phrase

// KeysIncludeParameters struct for KeysIncludeParameters
type KeysIncludeParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Specify a query to do broad search for keys by name (including wildcards).  The following qualifiers are also supported in the search term:  * `ids:key_id,...` for queries on a comma-separated list of ids * `name:key_name` for text queries on exact key names - spaces, commas, and colons need to be escaped with double backslashes * `tags:tag_name` to filter for keys with certain tags * `translated:{true|false}` for translation status (also requires `locale_id` to be specified) * `updated_at:{>=|<=}2013-02-21T00:00:00Z` for date range queries * `unmentioned_in_upload:upload_id,...` to filter keys unmentioned within upload. When multiple upload IDs provided, matches only keys not mentioned in **all** uploads  Find more examples [here](/en/api/strings/usage-examples).
	Q string `json:"q,omitempty"`
	// Include translations in locale
	TargetLocaleId string `json:"target_locale_id,omitempty"`
	// Tag or comma-separated list of tags to add to the matching collection of keys
	Tags string `json:"tags,omitempty"`
}
