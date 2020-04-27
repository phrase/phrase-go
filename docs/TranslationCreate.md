# TranslationCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | **string** | specify the branch to use | [optional] 
**LocaleId** | **string** | Locale. Can be the name or public id of the locale. Preferred is the public id. | [optional] 
**KeyId** | **string** | Key | [optional] 
**Content** | **string** | Translation content | [optional] 
**PluralSuffix** | **string** | Plural suffix. Can be one of: zero, one, two, few, many, other. Must be specified if the key associated to the translation is pluralized. | [optional] 
**Unverified** | **bool** | Indicates whether translation is unverified. Part of the &lt;a href&#x3D;\&quot;https://help.phrase.com/help/verification-and-proofreading\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Advanced Workflows&lt;/a&gt; feature. | [optional] 
**Excluded** | **bool** | Indicates whether translation is excluded. | [optional] 
**Autotranslate** | **bool** | Indicates whether the translation should be auto-translated. Responses with status 422 if provided for translation within a non-default locale or the project does not have the Autopilot feature enabled. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


