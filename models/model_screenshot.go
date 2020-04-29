package phrase
import (
	"time"
)
// Screenshot struct for Screenshot
type Screenshot struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ScreenshotUrl string `json:"screenshot_url,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	MarkersCount int32 `json:"markers_count,omitempty"`
}
