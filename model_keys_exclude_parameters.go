package phrase

// KeysExcludeParameters struct for KeysExcludeParameters
type KeysExcludeParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Specify a query to do broad search for keys by name (including wildcards).<br><br> The following qualifiers are also supported in the search term:<br> <ul>   <li><code>ids:key_id,...</code> for queries on a comma-separated list of ids</li>   <li><code>name:key_name</code> for text queries on exact key names - spaces, commas, and colons  need to be escaped with double backslashes</li>   <li><code>tags:tag_name</code> to filter for keys with certain tags</li>   <li><code>translated:{true|false}</code> for translation status (also requires <code>locale_id</code> to be specified)</li>   <li><code>updated_at:{>=|<=}2013-02-21T00:00:00Z</code> for date range queries</li>   <li><code>unmentioned_in_upload:upload_id</code> to filter keys unmentioned within upload</li> </ul> Find more examples <a href=\"#overview--usage-examples\">here</a>.
	Q string `json:"q,omitempty"`
	// Locale used to exlcude or include keys.
	TargetLocaleId string `json:"target_locale_id,omitempty"`
	// Tag or comma-separated list of tags to add to the matching collection of keys
	Tags string `json:"tags,omitempty"`
}
