package phrase

import (
	"time"
)

// JobUpdateParameters struct for JobUpdateParameters
type JobUpdateParameters struct {
	// Job name
	Name string `json:"name,omitempty"`
	// Briefing for the translators. Send an empty string or `null` to clear the value; omit the field to leave it unchanged. The Phrase CLI drops `null` values before sending, so use an empty string there.
	Briefing *NullableString `json:"briefing,omitempty"`
	// Date the job should be finished
	DueDate *NullableTime `json:"due_date,omitempty"`
	// Date the review of the job should be finished. Must not be before `due_date`.
	ReviewDueDate time.Time `json:"review_due_date,omitempty"`
	// URL to a ticket for this job (e.g. Jira, Trello) Send an empty string or `null` to clear the value; omit the field to leave it unchanged. The Phrase CLI drops `null` values before sending, so use an empty string there.
	TicketUrl *NullableString `json:"ticket_url,omitempty"`
	// List of target locales for the job.
	TargetLocaleIds []string `json:"target_locale_ids,omitempty"`
	// Automatically translate the job using machine translation
	Autotranslate *bool `json:"autotranslate,omitempty"`
}
