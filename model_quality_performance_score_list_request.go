package phrase

// QualityPerformanceScoreListRequest struct for QualityPerformanceScoreListRequest
type QualityPerformanceScoreListRequest struct {
	// Translation ids you want to get the quality performance score for
	TranslationIds []string `json:"translation_ids,omitempty"`
}
