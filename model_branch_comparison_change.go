package phrase

// BranchComparisonChange A single resource change. `from` holds the state before the change (null when the resource was added). `to` holds the state after (null when the resource was deleted). Fields depend on the resource type.
type BranchComparisonChange struct {
	From *map[string]interface{} `json:"from,omitempty"`
	To   *map[string]interface{} `json:"to,omitempty"`
}
