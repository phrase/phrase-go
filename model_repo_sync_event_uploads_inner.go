package phrase

// RepoSyncEventUploadsInner struct for RepoSyncEventUploadsInner
type RepoSyncEventUploadsInner struct {
	// The upload's public identifier (code).
	Id string `json:"id,omitempty"`
	// The upload's current public state. Multiple internal processing states are collapsed to `processing`; both internal error states are collapsed to `error`.
	State string `json:"state,omitempty"`
}
