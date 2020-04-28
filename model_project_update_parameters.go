package phrase
import (
	"os"
)
// ProjectUpdateParameters struct for ProjectUpdateParameters
type ProjectUpdateParameters struct {
	// Name of the project
	Name string `json:"name,omitempty"`
	// Main file format specified by its API Extension name. Used for locale downloads if no format is specified. For API Extension names of available file formats see <a href=\"https://help.phrase.com/help/supported-platforms-and-formats\">Format Guide</a> or our <a href=\"#formats\">Formats API Endpoint</a>.
	MainFormat string `json:"main_format,omitempty"`
	// Indicates whether the project should share the account's translation memory
	SharesTranslationMemory bool `json:"shares_translation_memory,omitempty"`
	// Image to identify the project
	ProjectImage *os.File `json:"project_image,omitempty"`
	// Indicates whether the project image should be deleted.
	RemoveProjectImage bool `json:"remove_project_image,omitempty"`
	// Account ID to specify the actual account the project should be created in. Required if the requesting user is a member of multiple accounts.
	AccountId string `json:"account_id,omitempty"`
}
