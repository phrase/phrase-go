package phrase

import (
	"time"
)

// Authorization struct for Authorization
type Authorization struct {
	Id             string    `json:"id,omitempty"`
	Note           string    `json:"note,omitempty"`
	TokenLastEight string    `json:"token_last_eight,omitempty"`
	HashedToken    string    `json:"hashed_token,omitempty"`
	Scopes         []string  `json:"scopes,omitempty"`
	ExpiresAt      time.Time `json:"expires_at,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}
