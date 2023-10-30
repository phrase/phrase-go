package phrase

// InvitationCreate422Response struct for InvitationCreate422Response
type InvitationCreate422Response struct {
	Message string                                   `json:"message,omitempty"`
	Errors  []InvitationCreate422ResponseErrorsInner `json:"errors,omitempty"`
}
