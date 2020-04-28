package phrase
// MemberUpdateParameters struct for MemberUpdateParameters
type MemberUpdateParameters struct {
	// Member role, can be any of of Manager, Developer, Translator
	Role string `json:"role,omitempty"`
	// List of project ids the user has access to. 
	ProjectIds string `json:"project_ids,omitempty"`
	// List of locale ids the user has access to.
	LocaleIds string `json:"locale_ids,omitempty"`
	// Additional permissions depending on member role. Available permissions are <code>create_upload</code> and <code>review_translations</code>
	Permissions map[string]interface{} `json:"permissions,omitempty"`
}
