package phrase

// SearchInAccountParameters struct for SearchInAccountParameters
type SearchInAccountParameters struct {
	// Search query
	Query string `json:"query,omitempty"`
	// Locale code
	LocaleCode string `json:"locale_code,omitempty"`
	// Page
	Page int32 `json:"page,omitempty"`
	// Number of results per page
	PerPage int32 `json:"per_page,omitempty"`
}
