package phrase

// BranchMergeParameters struct for BranchMergeParameters
type BranchMergeParameters struct {
	// strategy used for merge blocking, use_main or use_branch
	Strategy string `json:"strategy,omitempty"`
}
