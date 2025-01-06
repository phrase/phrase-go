package phrase

// QualityPerformanceScoreList200ResponseAnyOf struct for QualityPerformanceScoreList200ResponseAnyOf
type QualityPerformanceScoreList200ResponseAnyOf struct {
	Data QualityPerformanceScoreList200ResponseAnyOfData `json:"data,omitempty"`
	// Array of errors for any failing translation IDs
	Errors []QualityPerformanceScoreList200ResponseAnyOfErrorsInner `json:"errors,omitempty"`
}
