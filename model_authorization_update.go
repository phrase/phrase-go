package phrase
// AuthorizationUpdate struct for AuthorizationUpdate
type AuthorizationUpdate struct {
	// A note to help you remember what the access is used for.
	Note string `json:"note,omitempty"`
	// A list of scopes that the access can be used for.
	Scopes string `json:"scopes,omitempty"`
	// Expiration date for the authorization token. Null means no expiration date (default).
	ExpiresAt string `json:"expires_at,omitempty"`
}
