package phrase
// GitlabSyncImport struct for GitlabSyncImport
type GitlabSyncImport struct {
	// Account ID to specify the actual account the GitLab Sync should be created in. Required if the requesting user is a member of multiple accounts.
	AccountId string `json:"account_id,omitempty"`
}
