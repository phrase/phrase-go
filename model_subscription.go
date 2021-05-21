package phrase

// Subscription struct for Subscription
type Subscription struct {
	IsCurrent    *bool `json:"is_current,omitempty"`
	TrialExpired *bool `json:"trial_expired,omitempty"`
}
