package phrase
import (
	"time"
)
// BlacklistedKey struct for BlacklistedKey
type BlacklistedKey struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
