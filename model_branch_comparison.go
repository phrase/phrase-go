package phrase

// BranchComparison Comparison result of a branch against its base branch. Top-level properties correspond to resource types. Each resource type lists changes made in the base branch (`base_changes`), changes made in the feature branch (`head_changes`), and conflicting entries (`conflicts`).
type BranchComparison struct {
	TranslationKeys BranchComparisonDiff `json:"translation_keys,omitempty"`
	Translations    BranchComparisonDiff `json:"translations,omitempty"`
	Locales         BranchComparisonDiff `json:"locales,omitempty"`
	Tags            BranchComparisonDiff `json:"tags,omitempty"`
}
