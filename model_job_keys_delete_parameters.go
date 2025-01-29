package phrase

// JobKeysDeleteParameters struct for JobKeysDeleteParameters
type JobKeysDeleteParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// ids of keys that should be deleted from the job
	TranslationKeyIds []string `json:"translation_key_ids"`
}
