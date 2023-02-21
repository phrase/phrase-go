package phrase

// JobTemplateLocales struct for JobTemplateLocales
type JobTemplateLocales struct {
	Id          string             `json:"id,omitempty"`
	JobTemplate JobTemplatePreview `json:"job_template,omitempty"`
	Locale      LocalePreview      `json:"locale,omitempty"`
	Users       []Items            `json:"users,omitempty"`
	Teams       []Items            `json:"teams,omitempty"`
}
