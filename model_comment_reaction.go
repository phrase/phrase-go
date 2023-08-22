package phrase

import (
	"time"
)

// CommentReaction struct for CommentReaction
type CommentReaction struct {
	Id        string      `json:"id,omitempty"`
	Emoji     string      `json:"emoji,omitempty"`
	CreatedAt time.Time   `json:"created_at,omitempty"`
	UpdatedAt time.Time   `json:"updated_at,omitempty"`
	Comment   Comment     `json:"comment,omitempty"`
	User      UserPreview `json:"user,omitempty"`
}
