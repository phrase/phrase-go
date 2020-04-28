# \JobLocalesApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JobLocaleComplete**](JobLocalesApi.md#JobLocaleComplete) | **Post** /projects/{project_id}/jobs/{job_id}/locales/{id}/complete | Complete a job locale
[**JobLocaleDelete**](JobLocalesApi.md#JobLocaleDelete) | **Delete** /projects/{project_id}/jobs/{job_id}/locales/{id} | Delete a job locale
[**JobLocaleReopen**](JobLocalesApi.md#JobLocaleReopen) | **Post** /projects/{project_id}/jobs/{job_id}/locales/{id}/reopen | Reopen a job locale
[**JobLocaleShow**](JobLocalesApi.md#JobLocaleShow) | **Get** /projects/{project_id}/jobs/{job_id}/locale/{id} | Get a single job locale
[**JobLocaleUpdate**](JobLocalesApi.md#JobLocaleUpdate) | **Patch** /projects/{project_id}/jobs/{job_id}/locales/{id} | Update a job locale
[**JobLocalesCreate**](JobLocalesApi.md#JobLocalesCreate) | **Post** /projects/{project_id}/jobs/{job_id}/locales | Create a job locale
[**JobLocalesList**](JobLocalesApi.md#JobLocalesList) | **Get** /projects/{project_id}/jobs/{job_id}/locales | List job locales



## JobLocaleComplete

> map[string]interface{} JobLocaleComplete(ctx, projectId, jobId, id, jobLocaleCompleteParameters, optional)

Complete a job locale

Mark a job locale as completed.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobId** | **string**| Job ID | 
**id** | **string**| ID | 
**jobLocaleCompleteParameters** | [**JobLocaleCompleteParameters**](JobLocaleCompleteParameters.md)|  | 
 **optional** | ***JobLocaleCompleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobLocaleCompleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**map[string]interface{}**](object.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobLocaleDelete

> JobLocaleDelete(ctx, projectId, jobId, id, jobLocaleDeleteParameters, optional)

Delete a job locale

Delete an existing job locale.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobId** | **string**| Job ID | 
**id** | **string**| ID | 
**jobLocaleDeleteParameters** | [**JobLocaleDeleteParameters**](JobLocaleDeleteParameters.md)|  | 
 **optional** | ***JobLocaleDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobLocaleDeleteOpts struct


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


## JobLocaleReopen

> map[string]interface{} JobLocaleReopen(ctx, projectId, jobId, id, jobLocaleReopenParameters, optional)

Reopen a job locale

Mark a job locale as uncompleted.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobId** | **string**| Job ID | 
**id** | **string**| ID | 
**jobLocaleReopenParameters** | [**JobLocaleReopenParameters**](JobLocaleReopenParameters.md)|  | 
 **optional** | ***JobLocaleReopenOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobLocaleReopenOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**map[string]interface{}**](object.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobLocaleShow

> JobLocale JobLocaleShow(ctx, projectId, jobId, id, jobLocaleShowParameters, optional)

Get a single job locale

Get a single job locale for a given job.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobId** | **string**| Job ID | 
**id** | **string**| ID | 
**jobLocaleShowParameters** | [**JobLocaleShowParameters**](JobLocaleShowParameters.md)|  | 
 **optional** | ***JobLocaleShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobLocaleShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**JobLocale**](job_locale.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobLocaleUpdate

> map[string]interface{} JobLocaleUpdate(ctx, projectId, jobId, id, jobLocaleUpdateParameters, optional)

Update a job locale

Update an existing job locale.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobId** | **string**| Job ID | 
**id** | **string**| ID | 
**jobLocaleUpdateParameters** | [**JobLocaleUpdateParameters**](JobLocaleUpdateParameters.md)|  | 
 **optional** | ***JobLocaleUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobLocaleUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**map[string]interface{}**](object.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobLocalesCreate

> JobLocalesCreate(ctx, projectId, jobId, jobLocalesCreateParameters, optional)

Create a job locale

Create a new job locale.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobId** | **string**| Job ID | 
**jobLocalesCreateParameters** | [**JobLocalesCreateParameters**](JobLocalesCreateParameters.md)|  | 
 **optional** | ***JobLocalesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobLocalesCreateOpts struct


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


## JobLocalesList

> []map[string]interface{} JobLocalesList(ctx, projectId, jobId, jobLocalesListParameters, optional)

List job locales

List all job locales for a given job.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobId** | **string**| Job ID | 
**jobLocalesListParameters** | [**JobLocalesListParameters**](JobLocalesListParameters.md)|  | 
 **optional** | ***JobLocalesListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobLocalesListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **page** | **optional.Int32**| Page number | 
 **perPage** | **optional.Int32**| allows you to specify a page size up to 100 items, 10 by default | 

### Return type

[**[]map[string]interface{}**](object.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

