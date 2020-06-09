package phrase
// WebhookCreateParameters struct for WebhookCreateParameters
type WebhookCreateParameters struct {
	// Callback URL to send requests to
	CallbackUrl string `json:"callback_url,omitempty"`
	// Webhook description
	Description string `json:"description,omitempty"`
	// List of event names to trigger the webhook (separated by comma)
	Events string `json:"events,omitempty"`
	// Whether webhook is active or inactive
	Active *bool `json:"active,omitempty"`
}
