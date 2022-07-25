package phrase

import (
	"time"
)

// WebhookDelivery struct for WebhookDelivery
type WebhookDelivery struct {
	Id                 string    `json:"id,omitempty"`
	WebhookId          string    `json:"webhook_id,omitempty"`
	ResponseStatusCode int32     `json:"response_status_code,omitempty"`
	DeliveredAt        time.Time `json:"delivered_at,omitempty"`
	DurationMs         int32     `json:"duration_ms,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
}
