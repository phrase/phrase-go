package phrase

// Member struct for Member
type Member struct {
	Id                 string                 `json:"id,omitempty"`
	Email              string                 `json:"email,omitempty"`
	Username           string                 `json:"username,omitempty"`
	Role               string                 `json:"role,omitempty"`
	Projects           []ProjectLocales       `json:"projects,omitempty"`
	Permissions        map[string]interface{} `json:"permissions,omitempty"`
	DefaultLocaleCodes []string               `json:"default_locale_codes,omitempty"`
	Spaces             []MemberSpaces         `json:"spaces,omitempty"`
}
