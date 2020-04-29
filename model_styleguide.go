package phrase
import (
	"time"
)
// Styleguide struct for Styleguide
type Styleguide struct {
	Id string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
