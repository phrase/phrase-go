# phrase.GlossaryTermsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlossaryTermCreate**](GlossaryTermsApi.md#GlossaryTermCreate) | **Post** /accounts/{account_id}/glossaries/{glossary_id}/terms | Create a term
[**GlossaryTermDelete**](GlossaryTermsApi.md#GlossaryTermDelete) | **Delete** /accounts/{account_id}/glossaries/{glossary_id}/terms/{id} | Delete a term
[**GlossaryTermShow**](GlossaryTermsApi.md#GlossaryTermShow) | **Get** /accounts/{account_id}/glossaries/{glossary_id}/terms/{id} | Get a single term
[**GlossaryTermUpdate**](GlossaryTermsApi.md#GlossaryTermUpdate) | **Patch** /accounts/{account_id}/glossaries/{glossary_id}/terms/{id} | Update a term
[**GlossaryTermsList**](GlossaryTermsApi.md#GlossaryTermsList) | **Get** /accounts/{account_id}/glossaries/{glossary_id}/terms | List terms



## GlossaryTermCreate

> GlossaryTerm GlossaryTermCreate(ctx, accountId, glossaryId, glossaryTermCreateParameters, optional)

Create a term

Create a new term in a term base (previously: glossary).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**glossaryId** | **string**| Glossary ID | 
**glossaryTermCreateParameters** | [**GlossaryTermCreateParameters**](GlossaryTermCreateParameters.md)|  | 
 **optional** | ***GlossaryTermCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GlossaryTermCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**GlossaryTerm**](GlossaryTerm.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlossaryTermDelete

> GlossaryTermDelete(ctx, accountId, glossaryId, id, optional)

Delete a term

Delete an existing term in a term base (previously: glossary).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**glossaryId** | **string**| Glossary ID | 
**id** | **string**| ID | 
 **optional** | ***GlossaryTermDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GlossaryTermDeleteOpts struct


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


## GlossaryTermShow

> GlossaryTerm GlossaryTermShow(ctx, accountId, glossaryId, id, optional)

Get a single term

Get details for a single term in the term base (previously: glossary).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**glossaryId** | **string**| Glossary ID | 
**id** | **string**| ID | 
 **optional** | ***GlossaryTermShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GlossaryTermShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**GlossaryTerm**](GlossaryTerm.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlossaryTermUpdate

> GlossaryTerm GlossaryTermUpdate(ctx, accountId, glossaryId, id, glossaryTermUpdateParameters, optional)

Update a term

Update an existing term in a term base (previously: glossary).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**glossaryId** | **string**| Glossary ID | 
**id** | **string**| ID | 
**glossaryTermUpdateParameters** | [**GlossaryTermUpdateParameters**](GlossaryTermUpdateParameters.md)|  | 
 **optional** | ***GlossaryTermUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GlossaryTermUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**GlossaryTerm**](GlossaryTerm.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlossaryTermsList

> []GlossaryTerm GlossaryTermsList(ctx, accountId, glossaryId, optional)

List terms

List all terms in term bases (previously: glossary) that the current user has access to.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**glossaryId** | **string**| Glossary ID | 
 **optional** | ***GlossaryTermsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GlossaryTermsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**page** | **optional.Int32**| Page number | 
**perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 

### Return type

[**[]GlossaryTerm**](GlossaryTerm.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

