package phrase
import (
	"time"
)
// ReleasePreview struct for ReleasePreview
type ReleasePreview struct {
	Id string `json:"id,omitempty"`
	Version int32 `json:"version,omitempty"`
	AppMinVersion string `json:"app_min_version,omitempty"`
	AppMaxVersion string `json:"app_max_version,omitempty"`
	Description string `json:"description,omitempty"`
	Platforms []string `json:"platforms,omitempty"`
	Environments []string `json:"environments,omitempty"`
	LocaleCodes []string `json:"locale_codes,omitempty"`
	Project map[string]interface{} `json:"project,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
