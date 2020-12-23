package phrase

// InvitationUpdateSettingsParameters struct for InvitationUpdateSettingsParameters
type InvitationUpdateSettingsParameters struct {
	// Member role, can be any of of Manager, Developer, Translator
	ProjectRole string `json:"project_role,omitempty"`
	// List of locale ids the user has access to.
	LocaleIds []string `json:"locale_ids,omitempty"`
}
