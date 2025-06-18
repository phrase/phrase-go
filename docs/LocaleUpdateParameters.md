# LocaleUpdateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | **string** | specify the branch to use | [optional] 
**Name** | **string** | Locale name | [optional] 
**Code** | **string** | Locale ISO code | [optional] 
**Default** | **bool** | Indicates whether locale is the default locale. If set to true, the previous default locale the project is no longer the default locale. | [optional] 
**Main** | **bool** | Indicates whether locale is a main locale. Main locales are part of the [Verification System](https://support.phrase.com/hc/en-us/articles/5784094755484) feature. | [optional] 
**Rtl** | **bool** | Indicates whether locale is a RTL (Right-to-Left) locale. | [optional] 
**SourceLocaleId** | **string** | Source locale. Can be the name or id of the locale. Preferred is id. | [optional] 
**FallbackLocaleId** | **string** | Fallback locale for empty translations. Can be a locale name or id. | [optional] 
**UnverifyNewTranslations** | **bool** | Indicates that new translations for this locale should be marked as unverified. Part of the [Advanced Workflows](https://support.phrase.com/hc/en-us/articles/5784094755484) feature. | [optional] 
**UnverifyUpdatedTranslations** | **bool** | Indicates that updated translations for this locale should be marked as unverified. Part of the [Advanced Workflows](https://support.phrase.com/hc/en-us/articles/5784094755484) feature. | [optional] 
**Autotranslate** | **bool** | If set, translations for this locale will be fetched automatically, right after creation. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


