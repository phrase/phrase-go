package phrase

import (
	"os"
)

// ProjectCreateParameters struct for ProjectCreateParameters
type ProjectCreateParameters struct {
	// Name of the project
	Name string `json:"name"`
	// Main file format specified by its API Extension name. Used for locale downloads if no format is specified. For API Extension names of available file formats see [Format Guide](https://support.phrase.com/hc/en-us/sections/6111343326364) or our [Formats API Endpoint](/en/api/strings/formats/list-formats).
	MainFormat string `json:"main_format,omitempty"`
	// (Optional) Main technology stack used in the project. It affects for example the suggested placeholder style. Predefined values include: `Ruby`, `JavaScript`, `AngularJS`, `React`, `iOS`, `Android`, `Python`, `PHP`, `Java`, `Go`, `Windows Phone`, `Rails`, `Node.js`, `.NET`, `Django`, `Symfony`, `Yii Framework`, `Zend Framework`, `Apple App Store Description`, `Google Play Description`, but it can also take any other value.
	Media string `json:"media,omitempty"`
	// Indicates whether the project should share the account's translation memory
	SharesTranslationMemory *bool `json:"shares_translation_memory,omitempty"`
	// Image to identify the project
	ProjectImage *os.File `json:"project_image,omitempty"`
	// Indicates whether the project image should be deleted.
	RemoveProjectImage *bool `json:"remove_project_image,omitempty"`
	// Account ID to specify the actual account the project should be created in. Required if the requesting user is a member of multiple accounts.
	AccountId string `json:"account_id,omitempty"`
	// (Optional) User ID of the point of contact for the project.
	PointOfContact string `json:"point_of_contact,omitempty"`
	// When a source project ID is given, a clone of that project will be created, including all locales, keys and translations as well as the main project settings if they are not defined otherwise through the params.
	SourceProjectId string `json:"source_project_id,omitempty"`
	// (Optional) Review Workflow. \"simple\" / \"review\". [Read more](https://support.phrase.com/hc/en-us/articles/5784094755484)
	Workflow string `json:"workflow,omitempty"`
	// (Optional) Enable machine translation support in the project. Required for Pre-Translation
	MachineTranslationEnabled *bool `json:"machine_translation_enabled,omitempty"`
	// (Optional) Enable branching in the project
	EnableBranching *bool `json:"enable_branching,omitempty"`
	// (Optional) Protect the master branch in project where branching is enabled
	ProtectMasterBranch *bool `json:"protect_master_branch,omitempty"`
	// (Optional) Otherwise, translators are not allowed to edit translations other than strings
	EnableAllDataTypeTranslationKeysForTranslators *bool `json:"enable_all_data_type_translation_keys_for_translators,omitempty"`
	// (Optional) We can validate and highlight your ICU messages. [Read more](https://support.phrase.com/hc/en-us/articles/5822319545116)
	EnableIcuMessageFormat *bool `json:"enable_icu_message_format,omitempty"`
	// (Optional) Displays the input fields for the 'ZERO' plural form for every key as well although only some languages require the 'ZERO' explicitly.
	ZeroPluralFormEnabled *bool `json:"zero_plural_form_enabled,omitempty"`
	// (Optional) Autopilot, requires machine_translation_enabled. [Read more](https://support.phrase.com/hc/en-us/articles/5822187934364)
	AutotranslateEnabled *bool `json:"autotranslate_enabled,omitempty"`
	// (Optional) Requires autotranslate_enabled to be true
	AutotranslateCheckNewTranslationKeys *bool `json:"autotranslate_check_new_translation_keys,omitempty"`
	// (Optional) Requires autotranslate_enabled to be true
	AutotranslateCheckNewUploads *bool `json:"autotranslate_check_new_uploads,omitempty"`
	// (Optional) Requires autotranslate_enabled to be true
	AutotranslateCheckNewLocales *bool `json:"autotranslate_check_new_locales,omitempty"`
	// (Optional) Requires autotranslate_enabled to be true
	AutotranslateMarkAsUnverified *bool `json:"autotranslate_mark_as_unverified,omitempty"`
	// (Optional) Requires autotranslate_enabled to be true
	AutotranslateUseMachineTranslation *bool `json:"autotranslate_use_machine_translation,omitempty"`
	// (Optional) Requires autotranslate_enabled to be true
	AutotranslateUseTranslationMemory *bool `json:"autotranslate_use_translation_memory,omitempty"`
}
