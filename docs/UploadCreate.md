# UploadCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | **string** | specify the branch to use | [optional] 
**File** | **string** | File to be imported | [optional] 
**FileFormat** | **string** | File format. Auto-detected when possible and not specified. | [optional] 
**LocaleId** | **string** | Locale of the file&#39;s content. Can be the name or public id of the locale. Preferred is the public id. | [optional] 
**Tags** | **string** | List of tags separated by comma to be associated with the new keys contained in the upload. | [optional] 
**UpdateTranslations** | **string** | Indicates whether existing translations should be updated with the file content. | [optional] 
**UpdateDescriptions** | **string** | Existing key descriptions will be updated with the file content. Empty descriptions overwrite existing descriptions. | [optional] 
**ConvertEmoji** | **string** | This option is obsolete. Providing the option will cause a bad request error. | [optional] 
**SkipUploadTags** | **string** | Indicates whether the upload should not create upload tags. | [optional] 
**SkipUnverification** | **string** | Indicates whether the upload should unverify updated translations. | [optional] 
**FileEncoding** | **string** | Enforces a specific encoding on the file contents. Valid options are \&quot;UTF-8\&quot;, \&quot;UTF-16\&quot; and \&quot;ISO-8859-1\&quot;. | [optional] 
**LocaleMapping** | **string** | Optional, format specific mapping between locale names and the columns the translations to those locales are contained in. | [optional] 
**FormatOptions** | **string** | Additional options available for specific formats. See our format guide for complete list. | [optional] 
**Autotranslate** | **string** | If set, translations for the uploaded language will be fetched automatically. | [optional] 
**MarkReviewed** | **string** | Indicated whether the imported translations should be marked as reviewed. This setting is available if the review workflow (currently beta) is enabled for the project. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


