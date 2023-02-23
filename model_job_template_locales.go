package phrase

// JobTemplateLocales struct for JobTemplateLocales
type JobTemplateLocales struct {
	Id          string              `json:"id,omitempty"`
	JobTemplate JobTemplatePreview  `json:"job_template,omitempty"`
	Locale      LocalePreview       `json:"locale,omitempty"`
	Users       []LocaleUserPreview `json:"users,omitempty"`
	Teams       []LocaleTeamPreview `json:"teams,omitempty"`
}
