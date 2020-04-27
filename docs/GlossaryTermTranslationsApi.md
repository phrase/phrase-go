# \GlossaryTermTranslationsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlossaryTermTranslationCreate**](GlossaryTermTranslationsApi.md#GlossaryTermTranslationCreate) | **Post** /accounts/{account_id}/glossaries/{glossary_id}/terms/{term_id}/translations | Create a glossary term translation
[**GlossaryTermTranslationDelete**](GlossaryTermTranslationsApi.md#GlossaryTermTranslationDelete) | **Delete** /accounts/{account_id}/glossaries/{glossary_id}/terms/{term_id}/translations/{id} | Delete a glossary term translation
[**GlossaryTermTranslationUpdate**](GlossaryTermTranslationsApi.md#GlossaryTermTranslationUpdate) | **Patch** /accounts/{account_id}/glossaries/{glossary_id}/terms/{term_id}/translations/{id} | Update a glossary term translation



## GlossaryTermTranslationCreate

> GlossaryTermTranslationCreate(ctx, accountId, glossaryId, termId, glossaryTermTranslationCreate, optional)

Create a glossary term translation

Create a new glossary term translation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**glossaryId** | **string**| Glossary ID | 
**termId** | **string**| Term ID | 
**glossaryTermTranslationCreate** | [**GlossaryTermTranslationCreate**](GlossaryTermTranslationCreate.md)|  | 
 **optional** | ***GlossaryTermTranslationCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GlossaryTermTranslationCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlossaryTermTranslationDelete

> GlossaryTermTranslationDelete(ctx, accountId, glossaryId, termId, id, optional)

Delete a glossary term translation

Delete an existing glossary term translation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**glossaryId** | **string**| Glossary ID | 
**termId** | **string**| Term ID | 
**id** | **string**| ID | 
 **optional** | ***GlossaryTermTranslationDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GlossaryTermTranslationDeleteOpts struct


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


## GlossaryTermTranslationUpdate

> GlossaryTermTranslation GlossaryTermTranslationUpdate(ctx, accountId, glossaryId, termId, id, glossaryTermTranslationUpdate, optional)

Update a glossary term translation

Update an existing glossary term translation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**glossaryId** | **string**| Glossary ID | 
**termId** | **string**| Term ID | 
**id** | **string**| ID | 
**glossaryTermTranslationUpdate** | [**GlossaryTermTranslationUpdate**](GlossaryTermTranslationUpdate.md)|  | 
 **optional** | ***GlossaryTermTranslationUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GlossaryTermTranslationUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**GlossaryTermTranslation**](glossary_term_translation.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

