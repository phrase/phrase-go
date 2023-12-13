package phrase

import (
	"time"
)

// CustomMetadataProperty struct for CustomMetadataProperty
type CustomMetadataProperty struct {
	Id           string                 `json:"id,omitempty"`
	Name         string                 `json:"name,omitempty"`
	Description  string                 `json:"description,omitempty"`
	DataType     CustomMetadataDataType `json:"data_type,omitempty"`
	User         UserPreview            `json:"user,omitempty"`
	Projects     []ProjectShort         `json:"projects,omitempty"`
	ValueOptions []string               `json:"value_options,omitempty"`
	CreatedAt    time.Time              `json:"created_at,omitempty"`
	UpdatedAt    time.Time              `json:"updated_at,omitempty"`
}
