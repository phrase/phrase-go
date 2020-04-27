package phrase
import (
	"os"
)
// KeyUpdate struct for KeyUpdate
type KeyUpdate struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Key name
	Name string `json:"name,omitempty"`
	// Key description (usually includes contextual information for translators)
	Description string `json:"description,omitempty"`
	// Indicates whether key supports pluralization
	Plural bool `json:"plural,omitempty"`
	// Plural name for the key (used in some file formats, e.g. Gettext)
	NamePlural string `json:"name_plural,omitempty"`
	// Type of the key. Can be one of the following: string, number, boolean, array, markdown.
	DataType string `json:"data_type,omitempty"`
	// List of tags separated by comma to be associated with the key.
	Tags string `json:"tags,omitempty"`
	// Max. number of characters translations for this key can have.
	MaxCharactersAllowed int32 `json:"max_characters_allowed,omitempty"`
	// Screenshot/image for the key. This parameter is deprecated. Please use the Screenshots endpoint instead.
	Screenshot *os.File `json:"screenshot,omitempty"`
	// Indicates whether the screenshot will be deleted. This parameter is deprecated. Please use the Screenshots endpoint instead.
	RemoveScreenshot bool `json:"remove_screenshot,omitempty"`
	// Indicates whether the key should be exported as \"unformatted\". Supported by Android XML and other formats.
	Unformatted bool `json:"unformatted,omitempty"`
	// Indicates whether the key should be exported with \"xml:space=preserve\". Supported by several XML-based formats.
	XmlSpacePreserve bool `json:"xml_space_preserve,omitempty"`
	// Original file attribute. Used in some formats, e.g. XLIFF.
	OriginalFile string `json:"original_file,omitempty"`
	// NSStringLocalizedFormatKey attribute. Used in .stringsdict format.
	LocalizedFormatString string `json:"localized_format_string,omitempty"`
	// NSStringLocalizedFormatKey attribute. Used in .stringsdict format.
	LocalizedFormatKey string `json:"localized_format_key,omitempty"`
}
