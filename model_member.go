package phrase

// Member struct for Member
type Member struct {
	Id                 string           `json:"id,omitempty"`
	Email              string           `json:"email,omitempty"`
	Username           string           `json:"username,omitempty"`
	Role               string           `json:"role,omitempty"`
	DefaultLocaleCodes []string         `json:"default_locale_codes,omitempty"`
	Projects           []ProjectLocales `json:"projects,omitempty"`
}
