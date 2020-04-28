package phrase
import (
	"time"
)
// BitbucketSync struct for BitbucketSync
type BitbucketSync struct {
	Id string `json:"id,omitempty"`
	RepositoryName string `json:"repository_name,omitempty"`
	LastExportToBitbucketAt time.Time `json:"last_export_to_bitbucket_at,omitempty"`
	LastImportFromBitbucketAt time.Time `json:"last_import_from_bitbucket_at,omitempty"`
	ValidPhraseappYaml bool `json:"valid_phraseapp_yaml,omitempty"`
	PhraseappProjects []ProjectShort `json:"phraseapp_projects,omitempty"`
}
