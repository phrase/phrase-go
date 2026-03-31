package phrase

// RepoSyncImportParameters struct for RepoSyncImportParameters
type RepoSyncImportParameters struct {
	// Branch to import from
	RepositoryBranch string `json:"repository_branch,omitempty"`
	// Strings branch to import to
	Branch string `json:"branch,omitempty"`
}
