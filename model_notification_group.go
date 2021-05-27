package phrase

import (
	"time"
)

// NotificationGroup struct for NotificationGroup
type NotificationGroup struct {
	Id        string    `json:"id,omitempty"`
	EventName string    `json:"event_name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
