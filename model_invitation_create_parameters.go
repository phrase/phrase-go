package phrase

// InvitationCreateParameters struct for InvitationCreateParameters
type InvitationCreateParameters struct {
	// The email of the invited user. The `email` can not be updated once created. Create a new invitation for each unique email.
	Email string `json:"email"`
	// Invitiation role, can be any of Manager, Developer, Translator.
	Role string `json:"role"`
	// List of project ids the invited user has access to.
	ProjectIds string `json:"project_ids,omitempty"`
	// List of locale ids the invited user has access to.
	LocaleIds string `json:"locale_ids,omitempty"`
	// List of spaces the user is assigned to.
	SpaceIds []string `json:"space_ids,omitempty"`
	// List of teams the user is assigned to.
	TeamIds []string `json:"team_ids,omitempty"`
	// List of default locales for the user.
	DefaultLocaleCodes []string `json:"default_locale_codes,omitempty"`
	// Additional permissions depending on invitation role. Available permissions are `create_upload` and `review_translations`
	Permissions map[string]string `json:"permissions,omitempty"`
}
