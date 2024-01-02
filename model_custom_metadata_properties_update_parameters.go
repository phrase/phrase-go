package phrase

// CustomMetadataPropertiesUpdateParameters struct for CustomMetadataPropertiesUpdateParameters
type CustomMetadataPropertiesUpdateParameters struct {
	// name of the property
	Name string `json:"name,omitempty"`
	// ids of projects that the property belongs to
	ProjectIds []string `json:"project_ids,omitempty"`
	// description of property
	Description string `json:"description,omitempty"`
	// value options of property (only applies to single or multi select properties)
	ValueOptions []string `json:"value_options,omitempty"`
}
