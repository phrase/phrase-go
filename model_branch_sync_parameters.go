package phrase

// BranchSyncParameters struct for BranchSyncParameters
type BranchSyncParameters struct {
	// strategy used for conflicts, use_main or use_branch
	Strategy string `json:"strategy,omitempty"`
}
