# phrase.PreTranslationsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PreTranslationCreate**](PreTranslationsApi.md#PreTranslationCreate) | **Post** /projects/{project_id}/pre_translations | Create a pre-translation job
[**PreTranslationShow**](PreTranslationsApi.md#PreTranslationShow) | **Get** /projects/{project_id}/pre_translations/{id} | Get a single pre-translation job
[**PreTranslationsList**](PreTranslationsApi.md#PreTranslationsList) | **Get** /projects/{project_id}/pre_translations | List pre-translation jobs



## PreTranslationCreate

> PreTranslation PreTranslationCreate(ctx, projectId, preTranslationCreateParameters, optional)

Create a pre-translation job

Triggers a pre-translation job for a resource within a project, addressed by `translatable_type` (`locale`, `job`, `translation_key`, or `upload`) and `translatable_id` (its ID).  Enqueues machine translation using the project's configured MT engine. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**preTranslationCreateParameters** | [**PreTranslationCreateParameters**](PreTranslationCreateParameters.md)|  | 
 **optional** | ***PreTranslationCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PreTranslationCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**PreTranslation**](PreTranslation.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreTranslationShow

> PreTranslation PreTranslationShow(ctx, projectId, id, optional)

Get a single pre-translation job

Returns a single pre-translation job identified by its ID. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
 **optional** | ***PreTranslationShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PreTranslationShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**ifModifiedSince** | **optional.String**| Last modified condition, see [Conditional GET requests / HTTP Caching](/en/api/strings/pagination#conditional-get-requests-%2F-http-caching) (optional) | 
**ifNoneMatch** | **optional.String**| ETag condition, see [Conditional GET requests / HTTP Caching](/en/api/strings/pagination#conditional-get-requests-%2F-http-caching) (optional) | 

### Return type

[**PreTranslation**](PreTranslation.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreTranslationsList

> []PreTranslation PreTranslationsList(ctx, projectId, optional)

List pre-translation jobs

Returns all pre-translation jobs scoped to a project, ordered by creation date descending. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
 **optional** | ***PreTranslationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PreTranslationsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**ifModifiedSince** | **optional.String**| Last modified condition, see [Conditional GET requests / HTTP Caching](/en/api/strings/pagination#conditional-get-requests-%2F-http-caching) (optional) | 
**ifNoneMatch** | **optional.String**| ETag condition, see [Conditional GET requests / HTTP Caching](/en/api/strings/pagination#conditional-get-requests-%2F-http-caching) (optional) | 
**page** | **optional.Int32**| Page number | 
**perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 

### Return type

[**[]PreTranslation**](PreTranslation.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

