package phrase
// Member struct for Member
type Member struct {
	Id string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Role string `json:"role,omitempty"`
	Projects []map[string]interface{} `json:"projects,omitempty"`
}
