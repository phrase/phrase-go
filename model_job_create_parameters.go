package phrase

import (
	"time"
)

// JobCreateParameters struct for JobCreateParameters
type JobCreateParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Job name
	Name string `json:"name,omitempty"`
	// Briefing for the translators
	Briefing string `json:"briefing,omitempty"`
	// Date the job should be finished
	DueDate time.Time `json:"due_date,omitempty"`
	// tags of keys that should be included within the job
	Tags []string `json:"tags,omitempty"`
	// ids of keys that should be included within the job
	TranslationKeyIds []string `json:"translation_key_ids,omitempty"`
}
