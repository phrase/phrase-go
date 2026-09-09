# phrase.MachineTranslationApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MachineTranslationLocaleProviderMappingsCreate**](MachineTranslationApi.md#MachineTranslationLocaleProviderMappingsCreate) | **Post** /accounts/{account_id}/machine_translation_locale_provider_mappings | Create a locale provider mapping
[**MachineTranslationLocaleProviderMappingsDestroy**](MachineTranslationApi.md#MachineTranslationLocaleProviderMappingsDestroy) | **Delete** /accounts/{account_id}/machine_translation_locale_provider_mappings | Delete a locale provider mapping
[**MachineTranslationSettingsShow**](MachineTranslationApi.md#MachineTranslationSettingsShow) | **Get** /accounts/{account_id}/machine_translation_settings | Get machine translation settings
[**MachineTranslationSettingsUpdate**](MachineTranslationApi.md#MachineTranslationSettingsUpdate) | **Patch** /accounts/{account_id}/machine_translation_settings | Update machine translation settings



## MachineTranslationLocaleProviderMappingsCreate

> MachineTranslationLocaleProviderMapping MachineTranslationLocaleProviderMappingsCreate(ctx, accountId, machineTranslationLocaleProviderMappingsCreateParameters, optional)

Create a locale provider mapping

Creates a locale-pair-specific machine translation provider override for the account. When a mapping exists for a given source/target locale pair, that provider is used instead of the account default. Only one mapping may exist per source/target locale pair; attempting to create a duplicate returns a validation error. The source and target locale codes must differ. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**machineTranslationLocaleProviderMappingsCreateParameters** | [**MachineTranslationLocaleProviderMappingsCreateParameters**](MachineTranslationLocaleProviderMappingsCreateParameters.md)|  | 
 **optional** | ***MachineTranslationLocaleProviderMappingsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MachineTranslationLocaleProviderMappingsCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**MachineTranslationLocaleProviderMapping**](MachineTranslationLocaleProviderMapping.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineTranslationLocaleProviderMappingsDestroy

> MachineTranslationLocaleProviderMappingsDestroy(ctx, accountId, sourceLocaleCode, targetLocaleCode, optional)

Delete a locale provider mapping

Removes the machine translation provider override for the specified source and target locale pair. The mapping is identified by locale codes supplied as query parameters rather than a path ID. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**sourceLocaleCode** | **string**| The locale code of the source language of the mapping to delete. | 
**targetLocaleCode** | **string**| The locale code of the target language of the mapping to delete. | 
 **optional** | ***MachineTranslationLocaleProviderMappingsDestroyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MachineTranslationLocaleProviderMappingsDestroyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineTranslationSettingsShow

> MachineTranslationSettings MachineTranslationSettingsShow(ctx, accountId, optional)

Get machine translation settings

Returns the machine translation configuration for the account, including the default translation service, current machine translation unit usage, and any locale-pair-specific provider mappings. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
 **optional** | ***MachineTranslationSettingsShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MachineTranslationSettingsShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**MachineTranslationSettings**](MachineTranslationSettings.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineTranslationSettingsUpdate

> MachineTranslationSettings MachineTranslationSettingsUpdate(ctx, accountId, machineTranslationSettingsUpdateParameters, optional)

Update machine translation settings

Sets the default machine translation service for the account. Requires write access to the account's machine translation settings. Passing an empty or absent value for `default_service` resets the account to its plan default (Microsoft Translate). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**machineTranslationSettingsUpdateParameters** | [**MachineTranslationSettingsUpdateParameters**](MachineTranslationSettingsUpdateParameters.md)|  | 
 **optional** | ***MachineTranslationSettingsUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MachineTranslationSettingsUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**MachineTranslationSettings**](MachineTranslationSettings.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

