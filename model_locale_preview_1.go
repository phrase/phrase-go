package phrase

// LocalePreview1 struct for LocalePreview1
type LocalePreview1 struct {
	Id      string       `json:"id,omitempty"`
	Name    string       `json:"name,omitempty"`
	Code    string       `json:"code,omitempty"`
	Project ProjectShort `json:"project,omitempty"`
}
