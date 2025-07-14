package phrase

// KeyPreview struct for KeyPreview
type KeyPreview struct {
	Id              string `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Plural          *bool  `json:"plural,omitempty"`
	UseOrdinalRules *bool  `json:"use_ordinal_rules,omitempty"`
}
