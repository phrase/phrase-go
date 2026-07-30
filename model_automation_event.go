package phrase

import (
	"time"
)

// AutomationEvent struct for AutomationEvent
type AutomationEvent struct {
	// Unique identifier of the automation event.
	Id string `json:"id,omitempty"`
	// Identifier of the automation that produced this event.
	AutomationId string `json:"automation_id,omitempty"`
	// Outcome of the automation run.
	State string `json:"state,omitempty"`
	// What caused the automation to run.
	TriggeredBy string `json:"triggered_by,omitempty"`
	// Timestamp when the event was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Number of jobs created during this automation run.
	JobsCreated int32 `json:"jobs_created,omitempty"`
	// Identifiers of the jobs created during this automation run.
	JobIds  []string               `json:"job_ids,omitempty"`
	Project AutomationEventProject `json:"project,omitempty"`
	// Error message describing the failure when state is `failure`; null otherwise.
	Details *NullableString `json:"details,omitempty"`
}
