package phrase

import (
	"time"
)

// NotificationGroupDetail struct for NotificationGroupDetail
type NotificationGroupDetail struct {
	Id                 string                 `json:"id,omitempty"`
	EventName          string                 `json:"event_name,omitempty"`
	CreatedAt          time.Time              `json:"created_at,omitempty"`
	UpdatedAt          time.Time              `json:"updated_at,omitempty"`
	NotificationsCount int32                  `json:"notifications_count,omitempty"`
	LatestNotification map[string]interface{} `json:"latest_notification,omitempty"`
}
