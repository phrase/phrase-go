package phrase

// InvitationCreateParameters struct for InvitationCreateParameters
type InvitationCreateParameters struct {
	// The email of the invited user. The <code>email</code> can not be updated once created. Create a new invitation for each unique email.
	Email string `json:"email,omitempty"`
	// Invitiation role, can be any of Manager, Developer, Translator.
	Role string `json:"role,omitempty"`
	// List of project ids the invited user has access to.
	ProjectIds string `json:"project_ids,omitempty"`
	// List of locale ids the invited user has access to.
	LocaleIds string `json:"locale_ids,omitempty"`
	// List of spaces the user is assigned to.
	SpaceIds []string `json:"space_ids,omitempty"`
	// List of default locales for the user.
	DefaultLocaleCodes []string `json:"default_locale_codes,omitempty"`
	// Additional permissions depending on invitation role. Available permissions are <code>create_upload</code> and <code>review_translations</code>
	Permissions map[string]string `json:"permissions,omitempty"`
}
