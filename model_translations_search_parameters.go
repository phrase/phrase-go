package phrase

// TranslationsSearchParameters struct for TranslationsSearchParameters
type TranslationsSearchParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Sort criteria. Can be one of: key_name, created_at, updated_at.
	Sort string `json:"sort,omitempty"`
	// Order direction. Can be one of: asc, desc.
	Order string `json:"order,omitempty"`
	// Specify a query to find translations by content (including wildcards).<br><br> <i>Note: Search is limited to 10000 results and may not include recently updated data (depending on the project size).</i><br> The following qualifiers are supported in the query:<br> <ul>   <li><code>id:translation_id,...</code> for queries on a comma-separated list of ids</li>   <li><code>tags:XYZ</code> for tags on the translation</li>   <li><code>unverified:{true|false}</code> for verification status</li>   <li><code>excluded:{true|false}</code> for exclusion status</li>   <li><code>updated_at:{>=|<=}2013-02-21T00:00:00Z</code> for date range queries</li> </ul> Find more examples <a href=\"#overview--usage-examples\">here</a>.
	Q string `json:"q,omitempty"`
}
