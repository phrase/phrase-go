package phrase

// JobDetails1 struct for JobDetails1
type JobDetails1 struct {
	Owner      UserPreview     `json:"owner,omitempty"`
	JobTagName string          `json:"job_tag_name,omitempty"`
	Locales    []LocalePreview `json:"locales,omitempty"`
	Keys       []KeyPreview    `json:"keys,omitempty"`
}
