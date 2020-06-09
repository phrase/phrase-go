package phrase
// Format struct for Format
type Format struct {
	Name string `json:"name,omitempty"`
	ApiName string `json:"api_name,omitempty"`
	Description string `json:"description,omitempty"`
	Extension string `json:"extension,omitempty"`
	DefaultEncoding string `json:"default_encoding,omitempty"`
	Importable *bool `json:"importable,omitempty"`
	Exportable *bool `json:"exportable,omitempty"`
	DefaultFile string `json:"default_file,omitempty"`
	RendersDefaultLocale *bool `json:"renders_default_locale,omitempty"`
	IncludesLocaleInformation *bool `json:"includes_locale_information,omitempty"`
}
