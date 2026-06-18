package phrase

// RepoSyncCreateParameters struct for RepoSyncCreateParameters
type RepoSyncCreateParameters struct {
	// ID of the project to connect the Repo Sync to.
	ProjectId string `json:"project_id"`
	// Optional custom display name for this repo sync. Defaults to null; when null the project name is used as the display name.
	Name *NullableString `json:"name,omitempty"`
	// The Git provider to use.
	GitProvider string `json:"git_provider,omitempty"`
	// The authentication method used to connect to the Git provider. Defaults to `token` if not specified.  Valid values: - `token` — Personal access token stored on the Repo Sync. Supported by all providers. - `github_app` — Authenticate via the Phrase GitHub App installation on your account. GitHub only. The account must already have the GitHub App installed; if not, the response will include a `github_app_installation_url`. - `self_hosted` — Token-based auth for self-hosted Git instances. Requires `custom_api_endpoint`.
	ConnectionType string `json:"connection_type"`
	// Full repository name including the owner, e.g. `my-org/my-repo`.
	RepoName string `json:"repo_name"`
	// The default branch to use for imports and exports.
	BaseBranch string `json:"base_branch,omitempty"`
	// Branch that translations are exported to before opening a pull request. If omitted, exports go directly to `base_branch`.
	PrBranch string `json:"pr_branch,omitempty"`
	// Enable automatic import of translations triggered by pushes to the repository.
	AutoImport *bool `json:"auto_import,omitempty"`
	// Personal access token for the Git provider. Required when `connection_type` is `token` or `self_hosted`. Not used for `github_app`.
	AccessToken string `json:"access_token,omitempty"`
	// Custom API endpoint URL for self-hosted Git instances. Required when `connection_type` is `self_hosted`.
	CustomApiEndpoint string `json:"custom_api_endpoint,omitempty"`
}
