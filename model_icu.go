package phrase

// Icu struct for Icu
type Icu struct {
	// Object keys are dynamic and based on requested locale codes, see example.
	LocaleCode string `json:"locale_code,omitempty"`
}
