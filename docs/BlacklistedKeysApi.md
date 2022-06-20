# \BlacklistedKeysApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlacklistedKeyCreate**](BlacklistedKeysApi.md#BlacklistedKeyCreate) | **Post** /projects/{project_id}/blacklisted_keys | Create a blocked key
[**BlacklistedKeyDelete**](BlacklistedKeysApi.md#BlacklistedKeyDelete) | **Delete** /projects/{project_id}/blacklisted_keys/{id} | Delete a blocked key
[**BlacklistedKeyShow**](BlacklistedKeysApi.md#BlacklistedKeyShow) | **Get** /projects/{project_id}/blacklisted_keys/{id} | Get a single blocked key
[**BlacklistedKeyUpdate**](BlacklistedKeysApi.md#BlacklistedKeyUpdate) | **Patch** /projects/{project_id}/blacklisted_keys/{id} | Update a blocked key
[**BlacklistedKeysList**](BlacklistedKeysApi.md#BlacklistedKeysList) | **Get** /projects/{project_id}/blacklisted_keys | List blocked keys



## BlacklistedKeyCreate

> BlacklistedKey BlacklistedKeyCreate(ctx, projectId, blacklistedKeyCreateParameters, optional)

Create a blocked key

Create a new rule for blocking keys.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**blacklistedKeyCreateParameters** | [**BlacklistedKeyCreateParameters**](BlacklistedKeyCreateParameters.md)|  | 
 **optional** | ***BlacklistedKeyCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BlacklistedKeyCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**BlacklistedKey**](blacklisted_key.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlacklistedKeyDelete

> BlacklistedKeyDelete(ctx, projectId, id, optional)

Delete a blocked key

Delete an existing rule for blocking keys.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
 **optional** | ***BlacklistedKeyDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BlacklistedKeyDeleteOpts struct


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


## BlacklistedKeyShow

> BlacklistedKey BlacklistedKeyShow(ctx, projectId, id, optional)

Get a single blocked key

Get details on a single rule for blocking keys for a given project.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
 **optional** | ***BlacklistedKeyShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BlacklistedKeyShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**BlacklistedKey**](blacklisted_key.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlacklistedKeyUpdate

> BlacklistedKey BlacklistedKeyUpdate(ctx, projectId, id, blacklistedKeyUpdateParameters, optional)

Update a blocked key

Update an existing rule for blocking keys.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
**blacklistedKeyUpdateParameters** | [**BlacklistedKeyUpdateParameters**](BlacklistedKeyUpdateParameters.md)|  | 
 **optional** | ***BlacklistedKeyUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BlacklistedKeyUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**BlacklistedKey**](blacklisted_key.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlacklistedKeysList

> []BlacklistedKey BlacklistedKeysList(ctx, projectId, optional)

List blocked keys

List all rules for blocking keys for the given project.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
 **optional** | ***BlacklistedKeysListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BlacklistedKeysListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **page** | **optional.Int32**| Page number | 
 **perPage** | **optional.Int32**| allows you to specify a page size up to 100 items, 25 by default | 
 **branch** | **optional.String**| specify the branch to use | 

### Return type

[**[]BlacklistedKey**](blacklisted_key.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

