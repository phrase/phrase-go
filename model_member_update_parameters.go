package phrase

// MemberUpdateParameters struct for MemberUpdateParameters
type MemberUpdateParameters struct {
	// Update strategy, can be any of set, add, remove. If provided, it will set, add or remove given spaces, projects and locale ids from users access list.
	Strategy string `json:"strategy,omitempty"`
	// Member role, can be any of of Admin, ProjectManager, Developer, Designer, Translator
	Role string `json:"role,omitempty"`
	// List of project ids the user has access to.
	ProjectIds string `json:"project_ids,omitempty"`
	// List of locale ids the user has access to. This is project-scoped and must be paired with `project_ids` (and `strategy`) to take effect; it does not grant account-wide language access.
	LocaleIds string `json:"locale_ids,omitempty"`
	// List of default locales for the user, applied across all of the user's projects. This is the account-level field corresponding to \"Language access\" in the translator profile UI.
	DefaultLocaleCodes []string `json:"default_locale_codes,omitempty"`
	// List of spaces the user is assigned to.
	SpaceIds []string `json:"space_ids,omitempty"`
	// Additional permissions depending on member role. Available permissions are `create_upload` and `review_translations`
	Permissions map[string]string `json:"permissions,omitempty"`
}
