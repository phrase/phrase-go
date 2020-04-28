package phrase
import (
	"time"
)
// Release struct for Release
type Release struct {
	Id string `json:"id,omitempty"`
	Version int32 `json:"version,omitempty"`
	AppMinVersion string `json:"app_min_version,omitempty"`
	AppMaxVersion string `json:"app_max_version,omitempty"`
	Description string `json:"description,omitempty"`
	Platforms []string `json:"platforms,omitempty"`
	Environments []string `json:"environments,omitempty"`
	Locales []LocalePreview `json:"locales,omitempty"`
	Project ProjectShort `json:"project,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
