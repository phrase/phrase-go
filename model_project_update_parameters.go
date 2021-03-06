package phrase

import (
	"os"
)

// ProjectUpdateParameters struct for ProjectUpdateParameters
type ProjectUpdateParameters struct {
	// Required if the requesting user is a member of multiple accounts. Account ID to specify the actual account the project should be created in.
	AccountId string `json:"account_id,omitempty"`
	// (Optional) Name of the project
	Name string `json:"name,omitempty"`
	// (Optional) Main file format specified by its API Extension name. Used for locale downloads if no format is specified. For API Extension names of available file formats see <a href=\"https://help.phrase.com/help/supported-platforms-and-formats\">Format Guide</a> or our <a href=\"#formats\">Formats API Endpoint</a>.
	MainFormat string `json:"main_format,omitempty"`
	// (Optional) Indicates whether the project should share the account's translation memory
	SharesTranslationMemory *bool `json:"shares_translation_memory,omitempty"`
	// (Optional) Image to identify the project
	ProjectImage *os.File `json:"project_image,omitempty"`
	// (Optional) Indicates whether the project image should be deleted.
	RemoveProjectImage *bool `json:"remove_project_image,omitempty"`
	// (Optional) Review Workflow. \"simple\" / \"review\". <a href=\"https://help.phrase.com/help/advanced-review-workflow\">Read more</a>
	Workflow string `json:"workflow,omitempty"`
	// (Optional) Enable machine translation support in the project. Required for Autopilot and Smart Suggest
	MachineTranslationEnabled *bool `json:"machine_translation_enabled,omitempty"`
	// (Optional) Enable branching in the project
	EnableBranching *bool `json:"enable_branching,omitempty"`
	// (Optional) Protect the master branch in project where branching is enabled
	ProtectMasterBranch *bool `json:"protect_master_branch,omitempty"`
	// (Optional) Otherwise, translators are not allowed to edit translations other than strings
	EnableAllDataTypeTranslationKeysForTranslators *bool `json:"enable_all_data_type_translation_keys_for_translators,omitempty"`
	// (Optional) We can validate and highlight your ICU messages. <a href=\"https://help.phrase.com/help/icu-message-format\">Read more</a>
	EnableIcuMessageFormat *bool `json:"enable_icu_message_format,omitempty"`
	// (Optional) Displays the input fields for the 'ZERO' plural form for every key as well although only some languages require the 'ZERO' explicitly.
	ZeroPluralFormEnabled *bool `json:"zero_plural_form_enabled,omitempty"`
	// (Optional) Autopilot, requires machine_translation_enabled. <a href=\"https://help.phrase.com/help/autopilot\">Read more</a>
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
	// (Optional) Smart Suggest, requires machine_translation_enabled
	SmartSuggestEnabled *bool `json:"smart_suggest_enabled,omitempty"`
	// (Optional) Requires smart_suggest_enabled to be true
	SmartSuggestUseGlossary *bool `json:"smart_suggest_use_glossary,omitempty"`
	// (Optional) Requires smart_suggest_enabled to be true
	SmartSuggestUseMachineTranslation *bool `json:"smart_suggest_use_machine_translation,omitempty"`
}
