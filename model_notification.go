package phrase

import (
	"time"
)

// Notification struct for Notification
type Notification struct {
	Id          string                 `json:"id,omitempty"`
	EventName   string                 `json:"event_name,omitempty"`
	CreatedAt   time.Time              `json:"created_at,omitempty"`
	UpdatedAt   time.Time              `json:"updated_at,omitempty"`
	DeliveredAt time.Time              `json:"delivered_at,omitempty"`
	SeenAt      time.Time              `json:"seen_at,omitempty"`
	Data        map[string]interface{} `json:"data,omitempty"`
	Resource    map[string]interface{} `json:"resource,omitempty"`
	Locale      Locale                 `json:"locale,omitempty"`
	User        UserPreview            `json:"user,omitempty"`
	Project     Project                `json:"project,omitempty"`
	Account     Account                `json:"account,omitempty"`
	Group       NotificationGroup      `json:"group,omitempty"`
}
