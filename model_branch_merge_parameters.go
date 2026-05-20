package phrase

// BranchMergeParameters struct for BranchMergeParameters
type BranchMergeParameters struct {
	// Conflict resolution strategy applied when the branch and its base have diverged. `use_main` keeps the values from the base branch; `use_branch` keeps the values from the branch being merged.
	Strategy string `json:"strategy,omitempty"`
}
