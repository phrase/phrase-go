package phrase

// ErrorError struct for ErrorError
type ErrorError struct {
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}
