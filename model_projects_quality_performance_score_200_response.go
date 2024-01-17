package phrase

// ProjectsQualityPerformanceScore200Response struct for ProjectsQualityPerformanceScore200Response
type ProjectsQualityPerformanceScore200Response struct {
	Error ErrorError                                          `json:"error,omitempty"`
	Data  ProjectsQualityPerformanceScore200ResponseAnyOfData `json:"data,omitempty"`
	// Array of errors for any failing translation ids
	Errors []ProjectsQualityPerformanceScore200ResponseAnyOfErrorsInner `json:"errors,omitempty"`
}
