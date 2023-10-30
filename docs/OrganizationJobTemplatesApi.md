# \OrganizationJobTemplatesApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationJobTemplateCreate**](OrganizationJobTemplatesApi.md#OrganizationJobTemplateCreate) | **Post** /accounts/{account_id}/job_templates | Create an organization job template
[**OrganizationJobTemplateDelete**](OrganizationJobTemplatesApi.md#OrganizationJobTemplateDelete) | **Delete** /accounts/{account_id}/job_templates/{id} | Delete an organization job template
[**OrganizationJobTemplateUpdate**](OrganizationJobTemplatesApi.md#OrganizationJobTemplateUpdate) | **Patch** /accounts/{account_id}/job_templates/{id} | Update an organization job template
[**OrganizationJobTemplatesList**](OrganizationJobTemplatesApi.md#OrganizationJobTemplatesList) | **Get** /accounts/{account_id}/job_templates | List organization job templates
[**OrganizationJobTemplatesShow**](OrganizationJobTemplatesApi.md#OrganizationJobTemplatesShow) | **Get** /accounts/{account_id}/job_templates/{id} | Get a single organization job template



## OrganizationJobTemplateCreate

> OrganizationJobTemplateDetails OrganizationJobTemplateCreate(ctx, accountId, organizationJobTemplateCreateParameters, optional)

Create an organization job template

Create a new organization job template.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**organizationJobTemplateCreateParameters** | [**OrganizationJobTemplateCreateParameters**](OrganizationJobTemplateCreateParameters.md)|  | 
 **optional** | ***OrganizationJobTemplateCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganizationJobTemplateCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**OrganizationJobTemplateDetails**](OrganizationJobTemplateDetails.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationJobTemplateDelete

> OrganizationJobTemplateDelete(ctx, accountId, id, optional)

Delete an organization job template

Delete an existing organization job template.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**id** | **string**| ID | 
 **optional** | ***OrganizationJobTemplateDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganizationJobTemplateDeleteOpts struct


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


## OrganizationJobTemplateUpdate

> OrganizationJobTemplateDetails OrganizationJobTemplateUpdate(ctx, accountId, id, organizationJobTemplateUpdateParameters, optional)

Update an organization job template

Update an existing organization job template.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**id** | **string**| ID | 
**organizationJobTemplateUpdateParameters** | [**OrganizationJobTemplateUpdateParameters**](OrganizationJobTemplateUpdateParameters.md)|  | 
 **optional** | ***OrganizationJobTemplateUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganizationJobTemplateUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**OrganizationJobTemplateDetails**](OrganizationJobTemplateDetails.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationJobTemplatesList

> []OrganizationJobTemplate OrganizationJobTemplatesList(ctx, accountId, optional)

List organization job templates

List all job templates for the given account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
 **optional** | ***OrganizationJobTemplatesListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganizationJobTemplatesListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **page** | **optional.Int32**| Page number | 
 **perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 

### Return type

[**[]OrganizationJobTemplate**](OrganizationJobTemplate.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationJobTemplatesShow

> OrganizationJobTemplateDetails OrganizationJobTemplatesShow(ctx, accountId, id, optional)

Get a single organization job template

Get details on a single organization job template for a given account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**id** | **string**| ID | 
 **optional** | ***OrganizationJobTemplatesShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganizationJobTemplatesShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**OrganizationJobTemplateDetails**](OrganizationJobTemplateDetails.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

