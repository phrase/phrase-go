package phrase

// VariableCreateParameters struct for VariableCreateParameters
type VariableCreateParameters struct {
	// Name of the variable
	Name string `json:"name"`
	// Value of the variable
	Value string `json:"value,omitempty"`
}
