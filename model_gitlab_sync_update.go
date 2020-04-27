package phrase
// GitlabSyncUpdate struct for GitlabSyncUpdate
type GitlabSyncUpdate struct {
	// Account ID to specify the actual account the GitLab Sync should be created in. Required if the requesting user is a member of multiple accounts.
	AccountId string `json:"account_id,omitempty"`
	// Code of the related Phrase Project.
	PhraseProjectCode string `json:"phrase_project_code,omitempty"`
	// ID of the related GitLab Project.
	GitlabProjectId string `json:"gitlab_project_id,omitempty"`
	// Name of the GitLab Branch.
	GitlabBranchName string `json:"gitlab_branch_name,omitempty"`
}
