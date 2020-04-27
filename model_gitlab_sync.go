package phrase
import (
	"time"
)
// GitlabSync struct for GitlabSync
type GitlabSync struct {
	Id string `json:"id,omitempty"`
	ProjectId string `json:"project_id,omitempty"`
	GitlabProjectId int32 `json:"gitlab_project_id,omitempty"`
	GitlabBranchName string `json:"gitlab_branch_name,omitempty"`
	AutoImport bool `json:"auto_import,omitempty"`
	AutoImportSecret string `json:"auto_import_secret,omitempty"`
	AutoImportUrl string `json:"auto_import_url,omitempty"`
	SelfHostedApiUrl string `json:"self_hosted_api_url,omitempty"`
	LastExportedAt time.Time `json:"last_exported_at,omitempty"`
	LastImportedAt time.Time `json:"last_imported_at,omitempty"`
	LastStatus string `json:"last_status,omitempty"`
}
