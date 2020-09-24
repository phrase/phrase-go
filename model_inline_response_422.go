package phrase

// InlineResponse422 struct for InlineResponse422
type InlineResponse422 struct {
	Message string                    `json:"message,omitempty"`
	Errors  []InlineResponse422Errors `json:"errors,omitempty"`
}
