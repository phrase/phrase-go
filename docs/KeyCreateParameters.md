# KeyCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | **string** | specify the branch to use | [optional] 
**Name** | **string** | Key name | 
**Description** | **string** | Key description (usually includes contextual information for translators) | [optional] 
**Plural** | **bool** | Indicates whether key supports pluralization | [optional] 
**NamePlural** | **string** | Plural name for the key (used in some file formats, e.g. Gettext) | [optional] 
**DataType** | **string** | Type of the key. Can be one of the following: string, number, boolean, array, markdown. | [optional] 
**Tags** | **string** | List of tags separated by comma to be associated with the key. | [optional] 
**MaxCharactersAllowed** | **int32** | Max. number of characters translations for this key can have. | [optional] 
**Screenshot** | [***os.File**](*os.File.md) | Screenshot/image for the key. This parameter is deprecated. Please use the Screenshots endpoint instead. | [optional] 
**RemoveScreenshot** | **bool** | Indicates whether the screenshot will be deleted. This parameter is deprecated. Please use the Screenshots endpoint instead. | [optional] 
**Unformatted** | **bool** | Indicates whether the key should be exported as \&quot;unformatted\&quot;. Supported by Android XML and other formats. | [optional] 
**DefaultTranslationContent** | **string** | Creates a translation in the default locale with the specified content | [optional] 
**XmlSpacePreserve** | **bool** | Indicates whether the key should be exported with \&quot;xml:space&#x3D;preserve\&quot;. Supported by several XML-based formats. | [optional] 
**OriginalFile** | **string** | Original file attribute. Used in some formats, e.g. XLIFF. | [optional] 
**LocalizedFormatString** | **string** | NSStringLocalizedFormatKey attribute. Used in .stringsdict format. | [optional] 
**LocalizedFormatKey** | **string** | NSStringLocalizedFormatKey attribute. Used in .stringsdict format. | [optional] 
**CustomMetadata** | **map[string]interface{}** | Custom metadata property name and value pairs to be associated with key. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


