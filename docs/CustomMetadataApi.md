# \CustomMetadataApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomMetadataPropertiesDelete**](CustomMetadataApi.md#CustomMetadataPropertiesDelete) | **Delete** /accounts/{account_id}/custom_metadata/properties/{id} | Destroy property
[**CustomMetadataPropertiesList**](CustomMetadataApi.md#CustomMetadataPropertiesList) | **Get** /accounts/{account_id}/custom_metadata/properties | List properties
[**CustomMetadataPropertyCreate**](CustomMetadataApi.md#CustomMetadataPropertyCreate) | **Post** /accounts/{account_id}/custom_metadata/properties | Create a property
[**CustomMetadataPropertyShow**](CustomMetadataApi.md#CustomMetadataPropertyShow) | **Get** /accounts/{account_id}/custom_metadata/properties/{id} | Get a single property
[**CustomMetadataPropertyUpdate**](CustomMetadataApi.md#CustomMetadataPropertyUpdate) | **Patch** /accounts/{account_id}/custom_metadata/properties/{id} | Update a property



## CustomMetadataPropertiesDelete

> CustomMetadataPropertiesDelete(ctx, accountId, id, optional)

Destroy property

Destroy a custom metadata property of an account.  This endpoint is only available to accounts with advanced plans or above. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**id** | **string**| ID | 
 **optional** | ***CustomMetadataPropertiesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CustomMetadataPropertiesDeleteOpts struct


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


## CustomMetadataPropertiesList

> []CustomMetadataProperty CustomMetadataPropertiesList(ctx, accountId, optional)

List properties

List all custom metadata properties for an account.  This endpoint is only available to accounts with advanced plans or above. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
 **optional** | ***CustomMetadataPropertiesListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CustomMetadataPropertiesListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **dataType** | [**optional.Interface of CustomMetadataDataType**](.md)| Data Type of Custom Metadata Property | 
 **projectId** | **optional.String**| id of project that the properties belong to | 
 **page** | **optional.Int32**| Page number | 
 **perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 
 **sort** | **optional.String**| Sort criteria. Can be one of: name, data_type, created_at. | 
 **order** | **optional.String**| Order direction. Can be one of: asc, desc. | 

### Return type

[**[]CustomMetadataProperty**](CustomMetadataProperty.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomMetadataPropertyCreate

> CustomMetadataProperty CustomMetadataPropertyCreate(ctx, accountId, customMetadataPropertiesCreateParameters, optional)

Create a property

Create a new custom metadata property.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**customMetadataPropertiesCreateParameters** | [**CustomMetadataPropertiesCreateParameters**](CustomMetadataPropertiesCreateParameters.md)|  | 
 **optional** | ***CustomMetadataPropertyCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CustomMetadataPropertyCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**CustomMetadataProperty**](CustomMetadataProperty.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomMetadataPropertyShow

> CustomMetadataProperty CustomMetadataPropertyShow(ctx, accountId, id, optional)

Get a single property

Get details of a single custom property.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**id** | **string**| ID | 
 **optional** | ***CustomMetadataPropertyShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CustomMetadataPropertyShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**CustomMetadataProperty**](CustomMetadataProperty.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomMetadataPropertyUpdate

> CustomMetadataProperty CustomMetadataPropertyUpdate(ctx, accountId, id, customMetadataPropertiesUpdateParameters, optional)

Update a property

Update an existing custom metadata property.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**id** | **string**| ID | 
**customMetadataPropertiesUpdateParameters** | [**CustomMetadataPropertiesUpdateParameters**](CustomMetadataPropertiesUpdateParameters.md)|  | 
 **optional** | ***CustomMetadataPropertyUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CustomMetadataPropertyUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**CustomMetadataProperty**](CustomMetadataProperty.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

