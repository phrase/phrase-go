package phrase

// MemberSpaces struct for MemberSpaces
type MemberSpaces struct {
	Id            string      `json:"id,omitempty"`
	Name          string      `json:"name,omitempty"`
	CreatedAt     interface{} `json:"created_at,omitempty"`
	UpdatedAt     interface{} `json:"updated_at,omitempty"`
	ProjectsCount int32       `json:"projects_count,omitempty"`
}
