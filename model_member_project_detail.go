package phrase

// MemberProjectDetail struct for MemberProjectDetail
type MemberProjectDetail struct {
	Id                 string                                 `json:"id,omitempty"`
	Email              string                                 `json:"email,omitempty"`
	Username           string                                 `json:"username,omitempty"`
	Role               string                                 `json:"role,omitempty"`
	Projects           []ProjectLocales                       `json:"projects,omitempty"`
	Permissions        map[string]interface{}                 `json:"permissions,omitempty"`
	LocaleIds          []string                               `json:"locale_ids,omitempty"`
	DefaultLocaleCodes []string                               `json:"default_locale_codes,omitempty"`
	Spaces             []MemberSpacesInner                    `json:"spaces,omitempty"`
	ProjectRoles       []MemberProjectDetailProjectRolesInner `json:"project_roles,omitempty"`
}
