package phrase
import (
	"time"
)
// Invitation struct for Invitation
type Invitation struct {
	Id string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	Role string `json:"role,omitempty"`
	State string `json:"state,omitempty"`
	Projects []ProjectShort `json:"projects,omitempty"`
	Locales []LocalePreview `json:"locales,omitempty"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	AcceptedAt time.Time `json:"accepted_at,omitempty"`
}
