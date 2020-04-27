package phrase
import (
	"time"
)
// ProjectShort struct for ProjectShort
type ProjectShort struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	MainFormat string `json:"main_format,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
