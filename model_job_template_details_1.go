package phrase

// JobTemplateDetails1 struct for JobTemplateDetails1
type JobTemplateDetails1 struct {
	Owner   UserPreview     `json:"owner,omitempty"`
	Creator UserPreview     `json:"creator,omitempty"`
	Locales []LocalePreview `json:"locales,omitempty"`
}
