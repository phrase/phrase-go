package phrase

// Format struct for Format
type Format struct {
	// Human-readable display name of the format.
	Name string `json:"name"`
	// Identifier used to reference this format in API requests, such as the file_format parameter on the uploads and downloads endpoints.
	ApiName string `json:"api_name"`
	// Human-readable summary of the format and its typical use case.
	Description string `json:"description"`
	// Default file extension associated with this format.
	Extension string `json:"extension"`
	// Default character encoding used when reading or writing files in this format.
	DefaultEncoding string `json:"default_encoding"`
	// Whether locale files can be imported using this format.
	Importable *bool `json:"importable"`
	// Whether locale files can be exported using this format.
	Exportable *bool `json:"exportable"`
	// Conventional file path pattern for this format. Contains locale_name as a placeholder for the locale identifier.
	DefaultFile string `json:"default_file"`
	// When true, exported files contain the default locale's content for any key that has no translation in the target locale.
	RendersDefaultLocale *bool `json:"renders_default_locale"`
	// When true, files in this format embed locale information so Phrase can detect the locale automatically on import.
	IncludesLocaleInformation *bool `json:"includes_locale_information"`
}
