# \OrganizationJobTemplateLocalesApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationJobTemplateLocaleDelete**](OrganizationJobTemplateLocalesApi.md#OrganizationJobTemplateLocaleDelete) | **Delete** /accounts/{account_id}/job_templates/{job_template_id}/locales/{job_template_locale_id} | Delete an organization job template locale
[**OrganizationJobTemplateLocaleShow**](OrganizationJobTemplateLocalesApi.md#OrganizationJobTemplateLocaleShow) | **Get** /accounts/{account_id}/job_templates/{job_template_id}/locales/{job_template_locale_id} | Get a single organization job template locale
[**OrganizationJobTemplateLocaleUpdate**](OrganizationJobTemplateLocalesApi.md#OrganizationJobTemplateLocaleUpdate) | **Patch** /accounts/{account_id}/job_templates/{job_template_id}/locales/{job_template_locale_id} | Update an organization job template locale
[**OrganizationJobTemplateLocalesCreate**](OrganizationJobTemplateLocalesApi.md#OrganizationJobTemplateLocalesCreate) | **Post** /accounts/{account_id}/job_templates/{job_template_id}/locales | Create an organization job template locale
[**OrganizationJobTemplateLocalesList**](OrganizationJobTemplateLocalesApi.md#OrganizationJobTemplateLocalesList) | **Get** /accounts/{account_id}/job_templates/{job_template_id}/locales | List organization job template locales



## OrganizationJobTemplateLocaleDelete

> OrganizationJobTemplateLocaleDelete(ctx, accountId, jobTemplateId, jobTemplateLocaleId, optional)

Delete an organization job template locale

Delete an existing organization job template locale.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**jobTemplateId** | **string**| Job Template ID | 
**jobTemplateLocaleId** | **string**| Job Template Locale ID | 
 **optional** | ***OrganizationJobTemplateLocaleDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganizationJobTemplateLocaleDeleteOpts struct


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


## OrganizationJobTemplateLocaleShow

> JobTemplateLocales OrganizationJobTemplateLocaleShow(ctx, accountId, jobTemplateId, jobTemplateLocaleId, optional)

Get a single organization job template locale

Get a single job template locale for a given organization job template.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**jobTemplateId** | **string**| Job Template ID | 
**jobTemplateLocaleId** | **string**| Job Template Locale ID | 
 **optional** | ***OrganizationJobTemplateLocaleShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganizationJobTemplateLocaleShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**JobTemplateLocales**](job_template_locales.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationJobTemplateLocaleUpdate

> JobTemplateLocales OrganizationJobTemplateLocaleUpdate(ctx, accountId, jobTemplateId, jobTemplateLocaleId, organizationJobTemplateLocaleUpdateParameters, optional)

Update an organization job template locale

Update an existing organization job template locale.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**jobTemplateId** | **string**| Job Template ID | 
**jobTemplateLocaleId** | **string**| Job Template Locale ID | 
**organizationJobTemplateLocaleUpdateParameters** | [**OrganizationJobTemplateLocaleUpdateParameters**](OrganizationJobTemplateLocaleUpdateParameters.md)|  | 
 **optional** | ***OrganizationJobTemplateLocaleUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganizationJobTemplateLocaleUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**JobTemplateLocales**](job_template_locales.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationJobTemplateLocalesCreate

> JobTemplateLocales OrganizationJobTemplateLocalesCreate(ctx, accountId, jobTemplateId, organizationJobTemplateLocalesCreateParameters, optional)

Create an organization job template locale

Create a new organization job template locale.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**jobTemplateId** | **string**| Job Template ID | 
**organizationJobTemplateLocalesCreateParameters** | [**OrganizationJobTemplateLocalesCreateParameters**](OrganizationJobTemplateLocalesCreateParameters.md)|  | 
 **optional** | ***OrganizationJobTemplateLocalesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganizationJobTemplateLocalesCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**JobTemplateLocales**](job_template_locales.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationJobTemplateLocalesList

> []JobTemplateLocales OrganizationJobTemplateLocalesList(ctx, accountId, jobTemplateId, optional)

List organization job template locales

List all job template locales for a given organization job template.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**jobTemplateId** | **string**| Job Template ID | 
 **optional** | ***OrganizationJobTemplateLocalesListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganizationJobTemplateLocalesListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **page** | **optional.Int32**| Page number | 
 **perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 

### Return type

[**[]JobTemplateLocales**](job_template_locales.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

