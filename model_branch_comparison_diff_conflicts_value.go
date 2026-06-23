package phrase

// BranchComparisonDiffConflictsValue struct for BranchComparisonDiffConflictsValue
type BranchComparisonDiffConflictsValue struct {
	Base []map[string]interface{} `json:"base,omitempty"`
	Head []map[string]interface{} `json:"head,omitempty"`
}
