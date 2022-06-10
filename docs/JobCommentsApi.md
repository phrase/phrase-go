# \JobCommentsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JobCommentCreate**](JobCommentsApi.md#JobCommentCreate) | **Post** /projects/{project_id}/jobs/{job_id}/comments | Create a job comment
[**JobCommentDelete**](JobCommentsApi.md#JobCommentDelete) | **Delete** /projects/{project_id}/jobs/{job_id}/comments/{id} | Delete a job comment
[**JobCommentShow**](JobCommentsApi.md#JobCommentShow) | **Get** /projects/{project_id}/jobs/{job_id}/comments/{id} | Get a single job comment
[**JobCommentUpdate**](JobCommentsApi.md#JobCommentUpdate) | **Patch** /projects/{project_id}/jobs/{job_id}/comments/{id} | Update a job comment
[**JobCommentsList**](JobCommentsApi.md#JobCommentsList) | **Get** /projects/{project_id}/jobs/{job_id}/comments | List job comments



## JobCommentCreate

> JobComment JobCommentCreate(ctx, projectId, jobId, jobCommentCreateParameters, optional)

Create a job comment

Create a new comment for a job.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobId** | **string**| Job ID | 
**jobCommentCreateParameters** | [**JobCommentCreateParameters**](JobCommentCreateParameters.md)|  | 
 **optional** | ***JobCommentCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobCommentCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**JobComment**](job_comment.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobCommentDelete

> JobCommentDelete(ctx, projectId, jobId, id, optional)

Delete a job comment

Delete an existing job comment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobId** | **string**| Job ID | 
**id** | **string**| ID | 
 **optional** | ***JobCommentDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobCommentDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **branch** | **optional.String**| specify the branch to use | 

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


## JobCommentShow

> map[string]interface{} JobCommentShow(ctx, projectId, jobId, id, optional)

Get a single job comment

Get details on a single job comment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobId** | **string**| Job ID | 
**id** | **string**| ID | 
 **optional** | ***JobCommentShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobCommentShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **branch** | **optional.String**| specify the branch to use | 

### Return type

[**map[string]interface{}**](object.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobCommentUpdate

> map[string]interface{} JobCommentUpdate(ctx, projectId, keyId, id, jobCommentUpdateParameters, optional)

Update a job comment

Update an existing job comment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keyId** | **string**| Translation Key ID | 
**id** | **string**| ID | 
**jobCommentUpdateParameters** | [**JobCommentUpdateParameters**](JobCommentUpdateParameters.md)|  | 
 **optional** | ***JobCommentUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobCommentUpdateOpts struct


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


## JobCommentsList

> []map[string]interface{} JobCommentsList(ctx, projectId, jobId, optional)

List job comments

List all comments for a job.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobId** | **string**| Job ID | 
 **optional** | ***JobCommentsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobCommentsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **branch** | **optional.String**| specify the branch to use | 

### Return type

[**[]map[string]interface{}**](object.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

