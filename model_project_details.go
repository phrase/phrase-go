package phrase

import (
	"time"
)

// ProjectDetails struct for ProjectDetails
type ProjectDetails struct {
	Id                      string      `json:"id,omitempty"`
	Name                    string      `json:"name,omitempty"`
	Slug                    string      `json:"slug,omitempty"`
	MainFormat              string      `json:"main_format,omitempty"`
	ProjectImageUrl         string      `json:"project_image_url,omitempty"`
	Media                   string      `json:"media,omitempty"`
	Account                 Account     `json:"account,omitempty"`
	Space                   Space1      `json:"space,omitempty"`
	PointOfContact          UserPreview `json:"point_of_contact,omitempty"`
	CreatedAt               time.Time   `json:"created_at,omitempty"`
	UpdatedAt               time.Time   `json:"updated_at,omitempty"`
	SharesTranslationMemory *bool       `json:"shares_translation_memory,omitempty"`
}
