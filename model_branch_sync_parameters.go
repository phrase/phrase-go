package phrase

// BranchSyncParameters struct for BranchSyncParameters
type BranchSyncParameters struct {
	// Conflict resolution strategy applied when the branch and its base have diverged. `use_main` keeps the values from the base branch; `use_branch` keeps the values from this branch.
	Strategy string `json:"strategy,omitempty"`
}
