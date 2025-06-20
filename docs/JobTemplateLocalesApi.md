# phrase.JobTemplateLocalesApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JobTemplateLocaleDelete**](JobTemplateLocalesApi.md#JobTemplateLocaleDelete) | **Delete** /projects/{project_id}/job_templates/{job_template_id}/locales/{job_template_locale_id} | Delete a job template locale
[**JobTemplateLocaleShow**](JobTemplateLocalesApi.md#JobTemplateLocaleShow) | **Get** /projects/{project_id}/job_templates/{job_template_id}/locales/{job_template_locale_id} | Get a single job template locale
[**JobTemplateLocaleUpdate**](JobTemplateLocalesApi.md#JobTemplateLocaleUpdate) | **Patch** /projects/{project_id}/job_templates/{job_template_id}/locales/{job_template_locale_id} | Update a job template locale
[**JobTemplateLocalesCreate**](JobTemplateLocalesApi.md#JobTemplateLocalesCreate) | **Post** /projects/{project_id}/job_templates/{job_template_id}/locales | Create a job template locale
[**JobTemplateLocalesList**](JobTemplateLocalesApi.md#JobTemplateLocalesList) | **Get** /projects/{project_id}/job_templates/{job_template_id}/locales | List job template locales



## JobTemplateLocaleDelete

> JobTemplateLocaleDelete(ctx, projectId, jobTemplateId, jobTemplateLocaleId, optional)

Delete a job template locale

Delete an existing job template locale.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobTemplateId** | **string**| Job Template ID | 
**jobTemplateLocaleId** | **string**| Job Template Locale ID | 
 **optional** | ***JobTemplateLocaleDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobTemplateLocaleDeleteOpts struct


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


## JobTemplateLocaleShow

> JobTemplateLocales JobTemplateLocaleShow(ctx, projectId, jobTemplateId, jobTemplateLocaleId, optional)

Get a single job template locale

Get a single job template locale for a given job template.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobTemplateId** | **string**| Job Template ID | 
**jobTemplateLocaleId** | **string**| Job Template Locale ID | 
 **optional** | ***JobTemplateLocaleShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobTemplateLocaleShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**branch** | **optional.String**| specify the branch to use | 

### Return type

[**JobTemplateLocales**](JobTemplateLocales.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobTemplateLocaleUpdate

> JobTemplateLocales JobTemplateLocaleUpdate(ctx, projectId, jobTemplateId, jobTemplateLocaleId, jobTemplateLocaleUpdateParameters, optional)

Update a job template locale

Update an existing job template locale.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobTemplateId** | **string**| Job Template ID | 
**jobTemplateLocaleId** | **string**| Job Template Locale ID | 
**jobTemplateLocaleUpdateParameters** | [**JobTemplateLocaleUpdateParameters**](JobTemplateLocaleUpdateParameters.md)|  | 
 **optional** | ***JobTemplateLocaleUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobTemplateLocaleUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**JobTemplateLocales**](JobTemplateLocales.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobTemplateLocalesCreate

> JobTemplateLocales JobTemplateLocalesCreate(ctx, projectId, jobTemplateId, jobTemplateLocalesCreateParameters, optional)

Create a job template locale

Create a new job template locale.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobTemplateId** | **string**| Job Template ID | 
**jobTemplateLocalesCreateParameters** | [**JobTemplateLocalesCreateParameters**](JobTemplateLocalesCreateParameters.md)|  | 
 **optional** | ***JobTemplateLocalesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobTemplateLocalesCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**JobTemplateLocales**](JobTemplateLocales.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobTemplateLocalesList

> []JobTemplateLocales JobTemplateLocalesList(ctx, projectId, jobTemplateId, optional)

List job template locales

List all job template locales for a given job template.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**jobTemplateId** | **string**| Job Template ID | 
 **optional** | ***JobTemplateLocalesListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JobTemplateLocalesListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**page** | **optional.Int32**| Page number | 
**perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 
**branch** | **optional.String**| specify the branch to use | 

### Return type

[**[]JobTemplateLocales**](JobTemplateLocales.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

