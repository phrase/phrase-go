package phrase

// AutomationEventProject The project associated with this automation event. Null when no project is set.
type AutomationEventProject struct {
	// Project identifier.
	Id string `json:"id,omitempty"`
	// Project name.
	Name string `json:"name,omitempty"`
}
