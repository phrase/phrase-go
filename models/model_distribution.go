package phrase
import (
	"time"
)
// Distribution struct for Distribution
type Distribution struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Project ProjectShort `json:"project,omitempty"`
	Platforms []string `json:"platforms,omitempty"`
	Releases []ReleasePreview `json:"releases,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
