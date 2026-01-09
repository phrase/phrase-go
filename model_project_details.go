package phrase

import (
	"time"
)

// ProjectDetails struct for ProjectDetails
type ProjectDetails struct {
	Id                                             string      `json:"id,omitempty"`
	Name                                           string      `json:"name,omitempty"`
	Slug                                           string      `json:"slug,omitempty"`
	MainFormat                                     string      `json:"main_format,omitempty"`
	ProjectImageUrl                                string      `json:"project_image_url,omitempty"`
	Media                                          string      `json:"media,omitempty"`
	Account                                        Account     `json:"account,omitempty"`
	Space                                          Space1      `json:"space,omitempty"`
	PointOfContact                                 UserPreview `json:"point_of_contact,omitempty"`
	CreatedAt                                      time.Time   `json:"created_at,omitempty"`
	UpdatedAt                                      time.Time   `json:"updated_at,omitempty"`
	SharesTranslationMemory                        *bool       `json:"shares_translation_memory,omitempty"`
	MachineTranslationEnabled                      *bool       `json:"machine_translation_enabled,omitempty"`
	ZeroPluralFormEnabled                          *bool       `json:"zero_plural_form_enabled,omitempty"`
	EnableAllDataTypeTranslationKeysForTranslators *bool       `json:"enable_all_data_type_translation_keys_for_translators,omitempty"`
	EnableIcuMessageFormat                         *bool       `json:"enable_icu_message_format,omitempty"`
	EnableBranching                                *bool       `json:"enable_branching,omitempty"`
	ProtectMasterBranch                            *bool       `json:"protect_master_branch,omitempty"`
	AutotranslateEnabled                           *bool       `json:"autotranslate_enabled,omitempty"`
	AutotranslateCheckNewTranslationKeys           *bool       `json:"autotranslate_check_new_translation_keys,omitempty"`
	AutotranslateCheckNewUploads                   *bool       `json:"autotranslate_check_new_uploads,omitempty"`
	AutotranslateCheckNewLocales                   *bool       `json:"autotranslate_check_new_locales,omitempty"`
	AutotranslateMarkAsUnverified                  *bool       `json:"autotranslate_mark_as_unverified,omitempty"`
	AutotranslateUseMachineTranslation             *bool       `json:"autotranslate_use_machine_translation,omitempty"`
	AutotranslateUseTranslationMemory              *bool       `json:"autotranslate_use_translation_memory,omitempty"`
	DefaultEncoding                                string      `json:"default_encoding,omitempty"`
}
