package phrase
// JobsListParameters struct for JobsListParameters
type JobsListParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// filter by user owning job
	OwnedBy string `json:"owned_by,omitempty"`
	// filter by user assigned to job
	AssignedTo string `json:"assigned_to,omitempty"`
	// filter by state of job Valid states are <code>draft</code>, <code>in_progress</code>, <code>completed</code>
	State string `json:"state,omitempty"`
}
