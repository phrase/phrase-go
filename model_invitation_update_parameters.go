package phrase

// InvitationUpdateParameters struct for InvitationUpdateParameters
type InvitationUpdateParameters struct {
	// Invitiation role, can be any of Manager, Developer, Translator
	Role string `json:"role,omitempty"`
	// List of project ids the invited user has access to
	ProjectIds string `json:"project_ids,omitempty"`
	// List of locale ids the invited user has access to
	LocaleIds string `json:"locale_ids,omitempty"`
	// List of default locales for the user.
	DefaultLocaleCodes []string `json:"default_locale_codes,omitempty"`
	// Additional permissions depending on invitation role.
	Permissions map[string]string `json:"permissions,omitempty"`
}
