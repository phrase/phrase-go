package phrase

// BranchCreateParameters struct for BranchCreateParameters
type BranchCreateParameters struct {
	// Name of the branch
	Name string `json:"name"`
	// Name of an existing branch to use as the base for the new branch.
	Base string `json:"base,omitempty"`
}
