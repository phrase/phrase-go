package phrase
// JobKeysCreate struct for JobKeysCreate
type JobKeysCreate struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// ids of keys that should added to the job
	TranslationKeyIds string `json:"translation_key_ids,omitempty"`
}
