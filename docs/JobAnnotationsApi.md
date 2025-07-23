# phrase.JobAnnotationsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JobAnnotationDelete**](JobAnnotationsApi.md#JobAnnotationDelete) | **Delete** /projects/{project_id}/jobs/{job_id}/annotations/{id} | Delete a job annotation
[**JobAnnotationUpdate**](JobAnnotationsApi.md#JobAnnotationUpdate) | **Patch** /projects/{project_id}/jobs/{job_id}/annotations/{id} | Create/Update a job annotation
[**JobAnnotationsList**](JobAnnotationsApi.md#JobAnnotationsList) | **Get** /projects/{project_id}/jobs/{job_id}/annotations | List job annotations
[**JobLocaleAnnotationDelete**](JobAnnotationsApi.md#JobLocaleAnnotationDelete) | **Delete** /projects/{project_id}/jobs/{job_id}/locales/{job_locale_id}/annotations/{id} | Delete a job locale annotation
[**JobLocaleAnnotationUpdate**](JobAnnotationsApi.md#JobLocaleAnnotationUpdate) | **Patch** /projects/{project_id}/jobs/{job_id}/locales/{job_locale_id}/annotations/{id} | Create/Update a job locale annotation
[**JobLocaleAnnotationsList**](JobAnnotationsApi.md#JobLocaleAnnotationsList) | **Get** /projects/{project_id}/jobs/{job_id}/locales/{job_locale_id}/annotations | List job locale annotations



## JobAnnotationDelete

> JobAnnotationDelete(ctx, projectId, jobId, id, optional)

Delete a job annotation

Delete an annotation for a job.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobId** | **string**| Job ID | 
**id** | **string**| Name of the annotation to delete. | 
 **optional** | ***JobAnnotationDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobAnnotationDeleteOpts struct


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


## JobAnnotationUpdate

> JobAnnotation JobAnnotationUpdate(ctx, projectId, jobId, id, jobAnnotationUpdateParameters, optional)

Create/Update a job annotation

Create or update an annotation for a job. If the annotation already exists, it will be updated; otherwise, a new annotation will be created.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobId** | **string**| Job ID | 
**id** | **string**| Name of the annotation to set or update. | 
**jobAnnotationUpdateParameters** | [**JobAnnotationUpdateParameters**](JobAnnotationUpdateParameters.md)|  | 
 **optional** | ***JobAnnotationUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobAnnotationUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**JobAnnotation**](JobAnnotation.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobAnnotationsList

> []JobAnnotation JobAnnotationsList(ctx, projectId, jobId, optional)

List job annotations

Retrieve a list of annotations for a job.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobId** | **string**| Job ID | 
 **optional** | ***JobAnnotationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobAnnotationsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**branch** | **optional.String**| Branch to use | 

### Return type

[**[]JobAnnotation**](JobAnnotation.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobLocaleAnnotationDelete

> JobLocaleAnnotationDelete(ctx, projectId, jobId, jobLocaleId, id, optional)

Delete a job locale annotation

Delete an annotation for a job locale.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobId** | **string**| Job ID | 
**jobLocaleId** | **string**| Job Locale ID | 
**id** | **string**| Name of the annotation to delete. | 
 **optional** | ***JobLocaleAnnotationDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobLocaleAnnotationDeleteOpts struct


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


## JobLocaleAnnotationUpdate

> JobAnnotation JobLocaleAnnotationUpdate(ctx, projectId, jobId, jobLocaleId, id, jobAnnotationUpdateParameters, optional)

Create/Update a job locale annotation

Create or update an annotation for a job locale. If the annotation already exists, it will be updated; otherwise, a new annotation will be created.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobId** | **string**| Job ID | 
**jobLocaleId** | **string**| Job Locale ID | 
**id** | **string**| Name of the annotation to set or update. | 
**jobAnnotationUpdateParameters** | [**JobAnnotationUpdateParameters**](JobAnnotationUpdateParameters.md)|  | 
 **optional** | ***JobLocaleAnnotationUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobLocaleAnnotationUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**JobAnnotation**](JobAnnotation.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobLocaleAnnotationsList

> []JobAnnotation JobLocaleAnnotationsList(ctx, projectId, jobId, jobLocaleId, optional)

List job locale annotations

Retrieve a list of annotations for a job locale.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobId** | **string**| Job ID | 
**jobLocaleId** | **string**| Job Locale ID | 
 **optional** | ***JobLocaleAnnotationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobLocaleAnnotationsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**branch** | **optional.String**| Branch to use | 

### Return type

[**[]JobAnnotation**](JobAnnotation.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

