package phrase

import (
	"time"
)

// Project struct for Project
type Project struct {
	Id              string    `json:"id,omitempty"`
	Name            string    `json:"name,omitempty"`
	MainFormat      string    `json:"main_format,omitempty"`
	ProjectImageUrl string    `json:"project_image_url,omitempty"`
	Account         Account   `json:"account,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
}
