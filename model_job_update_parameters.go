package phrase

// JobUpdateParameters struct for JobUpdateParameters
type JobUpdateParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Job name
	Name string `json:"name,omitempty"`
	// Briefing for the translators
	Briefing string `json:"briefing,omitempty"`
	// Date the job should be finished
	DueDate *NullableTime `json:"due_date,omitempty"`
	// URL to a ticket for this job (e.g. Jira, Trello)
	TicketUrl string `json:"ticket_url,omitempty"`
}
