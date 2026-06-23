package phrase

// BranchComparisonDiff struct for BranchComparisonDiff
type BranchComparisonDiff struct {
	// Changes made to this resource type in the base branch since the branch was created.
	BaseChanges []BranchComparisonChange `json:"base_changes,omitempty"`
	// Changes made to this resource type in the feature branch.
	HeadChanges []BranchComparisonChange `json:"head_changes,omitempty"`
	// Conflicting changes present in both branches, keyed by conflict type. Possible conflict type keys: `changed_in_head_changed_in_base`, `added_in_head_added_in_base`, `changed_in_head_deleted_in_base`, `deleted_in_head_changed_in_base`. Each value contains `base` and `head` arrays of changed attribute objects.
	Conflicts map[string]BranchComparisonDiffConflictsValue `json:"conflicts,omitempty"`
}
