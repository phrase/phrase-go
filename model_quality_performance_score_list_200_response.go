package phrase

// QualityPerformanceScoreList200Response struct for QualityPerformanceScoreList200Response
type QualityPerformanceScoreList200Response struct {
	Error ErrorError                                      `json:"error,omitempty"`
	Data  QualityPerformanceScoreList200ResponseAnyOfData `json:"data,omitempty"`
	// Array of errors for any failing translation IDs
	Errors []QualityPerformanceScoreList200ResponseAnyOfErrorsInner `json:"errors,omitempty"`
}
