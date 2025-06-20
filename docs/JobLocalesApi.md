# phrase.JobLocalesApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JobLocaleComplete**](JobLocalesApi.md#JobLocaleComplete) | **Post** /projects/{project_id}/jobs/{job_id}/locales/{id}/complete | Complete a job locale
[**JobLocaleCompleteReview**](JobLocalesApi.md#JobLocaleCompleteReview) | **Post** /projects/{project_id}/jobs/{job_id}/locales/{id}/complete_review | Review a job locale
[**JobLocaleDelete**](JobLocalesApi.md#JobLocaleDelete) | **Delete** /projects/{project_id}/jobs/{job_id}/locales/{id} | Remove a target locale from a job
[**JobLocaleReopen**](JobLocalesApi.md#JobLocaleReopen) | **Post** /projects/{project_id}/jobs/{job_id}/locales/{id}/reopen | Reopen a job locale
[**JobLocaleShow**](JobLocalesApi.md#JobLocaleShow) | **Get** /projects/{project_id}/jobs/{job_id}/locales/{id} | Show single job target locale
[**JobLocaleUpdate**](JobLocalesApi.md#JobLocaleUpdate) | **Patch** /projects/{project_id}/jobs/{job_id}/locales/{id} | Update a job target locale
[**JobLocalesCreate**](JobLocalesApi.md#JobLocalesCreate) | **Post** /projects/{project_id}/jobs/{job_id}/locales | Add a target locale to a job
[**JobLocalesList**](JobLocalesApi.md#JobLocalesList) | **Get** /projects/{project_id}/jobs/{job_id}/locales | List job target locales



## JobLocaleComplete

> JobLocale JobLocaleComplete(ctx, projectId, jobId, id, jobLocaleCompleteParameters, optional)

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

[**JobLocale**](JobLocale.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobLocaleCompleteReview

> JobLocale JobLocaleCompleteReview(ctx, projectId, jobId, id, jobLocaleCompleteReviewParameters, optional)

Review a job locale

Mark job locale as reviewed.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobId** | **string**| Job ID | 
**id** | **string**| ID | 
**jobLocaleCompleteReviewParameters** | [**JobLocaleCompleteReviewParameters**](JobLocaleCompleteReviewParameters.md)|  | 
 **optional** | ***JobLocaleCompleteReviewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobLocaleCompleteReviewOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**JobLocale**](JobLocale.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobLocaleDelete

> JobLocaleDelete(ctx, projectId, jobId, id, optional)

Remove a target locale from a job

Removes a target locale from a job.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobId** | **string**| Job ID | 
**id** | **string**| ID | 
 **optional** | ***JobLocaleDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobLocaleDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**branch** | **optional.String**| Branch to use | 

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


## JobLocaleReopen

> JobLocale JobLocaleReopen(ctx, projectId, jobId, id, jobLocaleReopenParameters, optional)

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

[**JobLocale**](JobLocale.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobLocaleShow

> JobLocale JobLocaleShow(ctx, projectId, jobId, id, optional)

Show single job target locale

Get a single target locale for a given job.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobId** | **string**| Job ID | 
**id** | **string**| ID | 
 **optional** | ***JobLocaleShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobLocaleShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**branch** | **optional.String**| Branch to use | 

### Return type

[**JobLocale**](JobLocale.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobLocaleUpdate

> JobLocale JobLocaleUpdate(ctx, projectId, jobId, id, jobLocaleUpdateParameters, optional)

Update a job target locale

Update an existing job target locale.

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

[**JobLocale**](JobLocale.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobLocalesCreate

> JobLocale JobLocalesCreate(ctx, projectId, jobId, jobLocalesCreateParameters, optional)

Add a target locale to a job

Adds a target locale to a job.

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

[**JobLocale**](JobLocale.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobLocalesList

> []JobLocale JobLocalesList(ctx, projectId, jobId, optional)

List job target locales

List all target locales for a given job.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobId** | **string**| Job ID | 
 **optional** | ***JobLocalesListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobLocalesListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**page** | **optional.Int32**| Page number | 
**perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 
**branch** | **optional.String**| Branch to use | 

### Return type

[**[]JobLocale**](JobLocale.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

