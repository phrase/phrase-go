package phrase
// BitbucketSyncsList struct for BitbucketSyncsList
type BitbucketSyncsList struct {
	// Account ID to specify the actual account the project should be created in. Required if the requesting user is a member of multiple accounts.
	AccountId string `json:"account_id,omitempty"`
}
