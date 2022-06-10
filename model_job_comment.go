package phrase

import (
	"time"
)

// JobComment struct for JobComment
type JobComment struct {
	Id             string        `json:"id,omitempty"`
	Message        string        `json:"message,omitempty"`
	JobId          string        `json:"job_id,omitempty"`
	User           UserPreview   `json:"user,omitempty"`
	CreatedAt      time.Time     `json:"created_at,omitempty"`
	UpdatedAt      time.Time     `json:"updated_at,omitempty"`
	MentionedUsers []UserPreview `json:"mentioned_users,omitempty"`
}
