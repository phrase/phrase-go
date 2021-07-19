package phrase

// JobTemplateLocale struct for JobTemplateLocale
type JobTemplateLocale struct {
	Id          string                   `json:"id,omitempty"`
	JobTemplate JobTemplatePreview       `json:"job_template,omitempty"`
	Locale      LocalePreview            `json:"locale,omitempty"`
	Users       []JobTemplateUserPreview `json:"users,omitempty"`
}
