package phrase

// UserPreview struct for UserPreview
type UserPreview struct {
	Id          string `json:"id,omitempty"`
	Username    string `json:"username,omitempty"`
	Name        string `json:"name,omitempty"`
	GravatarUid string `json:"gravatar_uid,omitempty"`
}
