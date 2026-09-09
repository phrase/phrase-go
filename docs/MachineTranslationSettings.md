# MachineTranslationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultService** | Pointer to **NullableString** | The default machine translation engine configured for the account. Returns \&quot;microsoft_translate\&quot; when no service has been explicitly configured. Supported values: language_ai_translate, aita_translate, microsoft_translate, google_translate, amazon_translate, intento_translate, gpt_translate.  | [optional] 
**MachineTranslationUnitsUsed** | **int32** | Number of machine translation characters consumed in the current billing period. | [optional] 
**MachineTranslationUnitsTotal** | **int32** | Total machine translation character quota granted for the current billing period. | [optional] 
**LocaleProviderMappings** | [**[]MachineTranslationLocaleProviderMapping**](MachineTranslationLocaleProviderMapping.md) | Per-locale-pair provider overrides. When a matching mapping exists for a source/target locale pair, that provider takes precedence over the account default.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


