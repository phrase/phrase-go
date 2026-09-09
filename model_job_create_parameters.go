package phrase

// JobCreateParameters struct for JobCreateParameters
type JobCreateParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Job name
	Name string `json:"name"`
	// The API id of the source language
	SourceLocaleId string `json:"source_locale_id,omitempty"`
	// Briefing for the translators
	Briefing string `json:"briefing,omitempty"`
	// Date the job should be finished
	DueDate *NullableTime `json:"due_date,omitempty"`
	// URL to a ticket for this job (e.g. Jira, Trello)
	TicketUrl string `json:"ticket_url,omitempty"`
	// tags of keys that should be included within the job.  *Note: a tag matches every key currently carrying that tag, not just the ones you just tagged. For example, if hundreds of pre-existing keys already share the tag `myUploadTag`, adding it here pulls in every one of them, not only the key you just tagged. Use `translation_key_ids` to scope the job to specific keys instead.*
	Tags []string `json:"tags,omitempty"`
	// ids of keys that should be included within the job
	TranslationKeyIds []string `json:"translation_key_ids,omitempty"`
	// List of target locales for the job. Mutually exclusive with `job_template_id`.
	TargetLocaleIds []string `json:"target_locale_ids,omitempty"`
	// id of a job template you would like to model the created job after. Any manually added parameters will take preference over template attributes. Mutually exclusive with `target_locale_ids`.
	JobTemplateId string `json:"job_template_id,omitempty"`
	// Automatically translate the job using machine translation.
	Autotranslate *bool `json:"autotranslate,omitempty"`
}
