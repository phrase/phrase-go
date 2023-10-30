# \JobTemplatesApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JobTemplateCreate**](JobTemplatesApi.md#JobTemplateCreate) | **Post** /projects/{project_id}/job_templates | Create a job template
[**JobTemplateDelete**](JobTemplatesApi.md#JobTemplateDelete) | **Delete** /projects/{project_id}/job_templates/{id} | Delete a job template
[**JobTemplateUpdate**](JobTemplatesApi.md#JobTemplateUpdate) | **Patch** /projects/{project_id}/job_templates/{id} | Update a job template
[**JobTemplatesList**](JobTemplatesApi.md#JobTemplatesList) | **Get** /projects/{project_id}/job_templates | List job templates
[**JobTemplatesShow**](JobTemplatesApi.md#JobTemplatesShow) | **Get** /projects/{project_id}/job_templates/{id} | Get a single job template



## JobTemplateCreate

> JobTemplateDetails JobTemplateCreate(ctx, projectId, jobTemplateCreateParameters, optional)

Create a job template

Create a new job template.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobTemplateCreateParameters** | [**JobTemplateCreateParameters**](JobTemplateCreateParameters.md)|  | 
 **optional** | ***JobTemplateCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobTemplateCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**JobTemplateDetails**](JobTemplateDetails.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobTemplateDelete

> JobTemplateDelete(ctx, projectId, id, optional)

Delete a job template

Delete an existing job template.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
 **optional** | ***JobTemplateDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobTemplateDeleteOpts struct


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


## JobTemplateUpdate

> JobTemplateDetails JobTemplateUpdate(ctx, projectId, id, jobTemplateUpdateParameters, optional)

Update a job template

Update an existing job template.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
**jobTemplateUpdateParameters** | [**JobTemplateUpdateParameters**](JobTemplateUpdateParameters.md)|  | 
 **optional** | ***JobTemplateUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobTemplateUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**JobTemplateDetails**](JobTemplateDetails.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobTemplatesList

> []JobTemplate JobTemplatesList(ctx, projectId, optional)

List job templates

List all job templates for the given project.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
 **optional** | ***JobTemplatesListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobTemplatesListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **page** | **optional.Int32**| Page number | 
 **perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 
 **branch** | **optional.String**| specify the branch to use | 

### Return type

[**[]JobTemplate**](JobTemplate.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobTemplatesShow

> JobTemplateDetails JobTemplatesShow(ctx, projectId, id, optional)

Get a single job template

Get details on a single job template for a given project.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
 **optional** | ***JobTemplatesShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobTemplatesShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **branch** | **optional.String**| specify the branch to use | 

### Return type

[**JobTemplateDetails**](JobTemplateDetails.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

