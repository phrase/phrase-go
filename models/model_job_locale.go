package phrase
// JobLocale struct for JobLocale
type JobLocale struct {
	Id string `json:"id,omitempty"`
	Job JobPreview `json:"job,omitempty"`
	Locale LocalePreview `json:"locale,omitempty"`
	Users []UserPreview `json:"users,omitempty"`
}
