package phrase

import (
	"time"
)

// Member struct for Member
type Member struct {
	Id                 string                 `json:"id,omitempty"`
	Email              string                 `json:"email,omitempty"`
	Username           string                 `json:"username,omitempty"`
	CreatedAt          time.Time              `json:"created_at,omitempty"`
	LastActivityAt     time.Time              `json:"last_activity_at,omitempty"`
	Role               string                 `json:"role,omitempty"`
	Projects           []ProjectLocales       `json:"projects,omitempty"`
	Permissions        map[string]interface{} `json:"permissions,omitempty"`
	DefaultLocaleCodes []string               `json:"default_locale_codes,omitempty"`
	Teams              []Team1                `json:"teams,omitempty"`
	Spaces             []MemberSpaces         `json:"spaces,omitempty"`
}
