package phrase

// QualityPerformanceScoreList200ResponseAnyOfDataTranslationsInner struct for QualityPerformanceScoreList200ResponseAnyOfDataTranslationsInner
type QualityPerformanceScoreList200ResponseAnyOfDataTranslationsInner struct {
	// Engine used for the translation scoring
	Engine string `json:"engine,omitempty"`
	// Quality score for the translation
	Score float32 `json:"score,omitempty"`
	// One of the translation ids passed in arguments
	Id string `json:"id,omitempty"`
}
