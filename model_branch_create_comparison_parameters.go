package phrase

// BranchCreateComparisonParameters struct for BranchCreateComparisonParameters
type BranchCreateComparisonParameters struct {
	// direction of comparison, possible values are sync or merge (only for v2 branches)
	Direction string `json:"direction,omitempty"`
}
