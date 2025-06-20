# phrase.ReleasesApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReleaseCreate**](ReleasesApi.md#ReleaseCreate) | **Post** /accounts/{account_id}/distributions/{distribution_id}/releases | Create a release
[**ReleaseDelete**](ReleasesApi.md#ReleaseDelete) | **Delete** /accounts/{account_id}/distributions/{distribution_id}/releases/{id} | Delete a release
[**ReleasePublish**](ReleasesApi.md#ReleasePublish) | **Post** /accounts/{account_id}/distributions/{distribution_id}/releases/{id}/publish | Publish a release
[**ReleaseShow**](ReleasesApi.md#ReleaseShow) | **Get** /accounts/{account_id}/distributions/{distribution_id}/releases/{id} | Get a single release
[**ReleaseUpdate**](ReleasesApi.md#ReleaseUpdate) | **Patch** /accounts/{account_id}/distributions/{distribution_id}/releases/{id} | Update a release
[**ReleasesList**](ReleasesApi.md#ReleasesList) | **Get** /accounts/{account_id}/distributions/{distribution_id}/releases | List releases



## ReleaseCreate

> Release ReleaseCreate(ctx, accountId, distributionId, releaseCreateParameters, optional)

Create a release

Create a new release.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**distributionId** | **string**| Distribution ID | 
**releaseCreateParameters** | [**ReleaseCreateParameters**](ReleaseCreateParameters.md)|  | 
 **optional** | ***ReleaseCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReleaseCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**Release**](Release.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleaseDelete

> ReleaseDelete(ctx, accountId, distributionId, id, optional)

Delete a release

Delete an existing release.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**distributionId** | **string**| Distribution ID | 
**id** | **string**| ID | 
 **optional** | ***ReleaseDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReleaseDeleteOpts struct


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


## ReleasePublish

> Release ReleasePublish(ctx, accountId, distributionId, id, optional)

Publish a release

Publish a release for production.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**distributionId** | **string**| Distribution ID | 
**id** | **string**| ID | 
 **optional** | ***ReleasePublishOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReleasePublishOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**Release**](Release.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleaseShow

> Release ReleaseShow(ctx, accountId, distributionId, id, optional)

Get a single release

Get details on a single release.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**distributionId** | **string**| Distribution ID | 
**id** | **string**| ID | 
 **optional** | ***ReleaseShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReleaseShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**Release**](Release.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleaseUpdate

> Release ReleaseUpdate(ctx, accountId, distributionId, id, releaseUpdateParameters, optional)

Update a release

Update an existing release.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**distributionId** | **string**| Distribution ID | 
**id** | **string**| ID | 
**releaseUpdateParameters** | [**ReleaseUpdateParameters**](ReleaseUpdateParameters.md)|  | 
 **optional** | ***ReleaseUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReleaseUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**Release**](Release.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleasesList

> []ReleasePreview ReleasesList(ctx, accountId, distributionId, optional)

List releases

List all releases for the given distribution.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**distributionId** | **string**| Distribution ID | 
 **optional** | ***ReleasesListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReleasesListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**page** | **optional.Int32**| Page number | 
**perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 

### Return type

[**[]ReleasePreview**](ReleasePreview.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

