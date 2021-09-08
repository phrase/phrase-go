package phrase

// TranslationsVerifyParameters struct for TranslationsVerifyParameters
type TranslationsVerifyParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// specify the locale of the translations to be verified
	LocaleId string `json:"locale_id,omitempty"`
	// Specify a query to find translations by content (including wildcards).<br><br> The following qualifiers are supported in the query:<br> <ul>   <li><code>id:translation_id,...</code> for queries on a comma-separated list of ids</li>   <li><code>tags:XYZ</code> for tags on the translation</li>   <li><code>unverified:{true|false}</code> for verification status</li>   <li><code>excluded:{true|false}</code> for exclusion status</li>   <li><code>updated_at:{>=|<=}2013-02-21T00:00:00Z</code> for date range queries</li> </ul> Find more examples <a href=\"#overview--usage-examples\">here</a>.
	Q string `json:"q,omitempty"`
}
