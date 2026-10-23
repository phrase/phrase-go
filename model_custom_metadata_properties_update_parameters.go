package phrase

// CustomMetadataPropertiesUpdateParameters struct for CustomMetadataPropertiesUpdateParameters
type CustomMetadataPropertiesUpdateParameters struct {
	// name of the property
	Name string `json:"name,omitempty"`
	// ids of projects that the property belongs to
	ProjectIds []string `json:"project_ids,omitempty"`
	// description of property. Send an empty string or `null` to clear the value; omit the field to leave it unchanged. The Phrase CLI drops `null` values before sending, so use an empty string there.
	Description *NullableString `json:"description,omitempty"`
	// value options of property (only applies to single or multi select properties)
	ValueOptions []string `json:"value_options,omitempty"`
}
