package phrase
// JobLocalesCreate struct for JobLocalesCreate
type JobLocalesCreate struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// locale id
	LocaleId string `json:"locale_id,omitempty"`
	// Array of user ids to be assigned to the job locale
	UserIds []string `json:"user_ids,omitempty"`
}
