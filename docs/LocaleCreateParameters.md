# LocaleCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | **string** | specify the branch to use | [optional] 
**Name** | **string** | Locale name. Must be unique per project. | 
**Code** | **string** | Locale ISO code. Unlike &#x60;name&#x60;, &#x60;code&#x60; is not required to be unique per project - creating a locale whose &#x60;code&#x60; duplicates an existing locale&#39;s &#x60;code&#x60; in the same project will succeed rather than error, resulting in two locales that share the same code. | 
**Default** | **bool** | Indicates whether locale is the default locale. If set to true, the previous default locale the project is no longer the default locale. | [optional] 
**Main** | **bool** | Indicates whether locale is a main locale. Main locales are part of the [Verification System](https://support.phrase.com/hc/en-us/articles/5784094755484) feature. | [optional] 
**Rtl** | **bool** | Indicates whether locale is a RTL (Right-to-Left) locale. | [optional] 
**SourceLocaleId** | **string** | Source locale. Can be the name or id of the locale. Preferred is id. | [optional] 
**FallbackLocaleId** | **string** | Fallback locale for empty translations. Can be a locale name or id. | [optional] 
**UnverifyNewTranslations** | **bool** | Indicates that new translations for this locale should be marked as unverified. Part of the [Advanced Workflows](https://support.phrase.com/hc/en-us/articles/5784094755484) feature. | [optional] 
**UnverifyUpdatedTranslations** | **bool** | Indicates that updated translations for this locale should be marked as unverified. Part of the [Advanced Workflows](https://support.phrase.com/hc/en-us/articles/5784094755484) feature. | [optional] 
**UnverifyOnSourceChanges** | **bool** | Indicates that translations for this locale should be marked as unverified when the source language has been changed. | [optional] 
**Autotranslate** | **bool** | If set, translations for this locale will be fetched automatically, right after creation. | [optional] 
**LanguageAiProfile** | **string** | Identifier of the Language AI profile to use for this locale. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


