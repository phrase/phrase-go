package phrase
// JobLocale struct for JobLocale
type JobLocale struct {
	Id string `json:"id,omitempty"`
	Job JobPreview `json:"job,omitempty"`
	Locale map[string]interface{} `json:"locale,omitempty"`
	Users []map[string]interface{} `json:"users,omitempty"`
}
