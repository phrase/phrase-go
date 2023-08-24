package phrase

import (
	"time"
)

// Comment struct for Comment
type Comment struct {
	Id             string          `json:"id,omitempty"`
	Message        string          `json:"message,omitempty"`
	HasReplies     *bool           `json:"has_replies,omitempty"`
	User           UserPreview     `json:"user,omitempty"`
	CreatedAt      time.Time       `json:"created_at,omitempty"`
	UpdatedAt      time.Time       `json:"updated_at,omitempty"`
	MentionedUsers []UserPreview   `json:"mentioned_users,omitempty"`
	Locales        []LocalePreview `json:"locales,omitempty"`
}
