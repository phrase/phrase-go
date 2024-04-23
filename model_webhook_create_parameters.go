package phrase

// WebhookCreateParameters struct for WebhookCreateParameters
type WebhookCreateParameters struct {
	// Callback URL to send requests to
	CallbackUrl string `json:"callback_url"`
	// Webhook secret used to calculate signature. If empty, the default project secret will be used.
	Secret string `json:"secret,omitempty"`
	// Webhook description
	Description string `json:"description,omitempty"`
	// List of event names to trigger the webhook (separated by comma)
	Events string `json:"events,omitempty"`
	// Whether webhook is active or inactive
	Active *bool `json:"active,omitempty"`
	// If enabled, webhook will also be triggered for events from branches of the project specified.
	IncludeBranches *bool `json:"include_branches,omitempty"`
}
