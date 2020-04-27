package phrase
// Errors struct for Errors
type Errors struct {
	Message string `json:"message,omitempty"`
	Errors []ErrorsErrors `json:"errors,omitempty"`
}
