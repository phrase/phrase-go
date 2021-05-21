package phrase

import (
	"time"
)

// Account struct for Account
type Account struct {
	Id             string       `json:"id,omitempty"`
	Name           string       `json:"name,omitempty"`
	Slug           string       `json:"slug,omitempty"`
	Company        string       `json:"company,omitempty"`
	CreatedAt      time.Time    `json:"created_at,omitempty"`
	UpdatedAt      time.Time    `json:"updated_at,omitempty"`
	CompanyLogoUrl string       `json:"company_logo_url,omitempty"`
	Subscription   Subscription `json:"subscription,omitempty"`
}
