# phrase.LinkedKeysApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KeyLinksBatchDestroy**](LinkedKeysApi.md#KeyLinksBatchDestroy) | **Delete** /projects/{project_id}/keys/{id}/key_links | Batch unlink child keys from a parent key
[**KeyLinksCreate**](LinkedKeysApi.md#KeyLinksCreate) | **Post** /projects/{project_id}/keys/{id}/key_links | Link child keys to a parent key
[**KeyLinksDestroy**](LinkedKeysApi.md#KeyLinksDestroy) | **Delete** /projects/{project_id}/keys/{id}/key_links/{child_key_id} | Unlink a child key from a parent key
[**KeyLinksIndex**](LinkedKeysApi.md#KeyLinksIndex) | **Get** /projects/{project_id}/keys/{id}/key_links | List child keys of a parent key



## KeyLinksBatchDestroy

> KeyLinksBatchDestroy(ctx, projectId, id, keyLinksBatchDestroyParameters, optional)

Batch unlink child keys from a parent key

Unlinks multiple child keys from a given parent key in a single operation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| Parent Translation Key ID | 
**keyLinksBatchDestroyParameters** | [**KeyLinksBatchDestroyParameters**](KeyLinksBatchDestroyParameters.md)|  | 
 **optional** | ***KeyLinksBatchDestroyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KeyLinksBatchDestroyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyLinksCreate

> KeyLink KeyLinksCreate(ctx, projectId, id, keyLinksCreateParameters, optional)

Link child keys to a parent key

Creates links between a given parent key and one or more child keys.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| Parent Translation Key ID | 
**keyLinksCreateParameters** | [**KeyLinksCreateParameters**](KeyLinksCreateParameters.md)|  | 
 **optional** | ***KeyLinksCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KeyLinksCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**KeyLink**](KeyLink.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyLinksDestroy

> KeyLinksDestroy(ctx, projectId, id, childKeyId, optional)

Unlink a child key from a parent key

Unlinks a single child key from a given parent key.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| Parent Translation Key ID | 
**childKeyId** | **string**| The ID of the child key to unlink. | 
 **optional** | ***KeyLinksDestroyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KeyLinksDestroyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyLinksIndex

> KeyLink KeyLinksIndex(ctx, projectId, id, optional)

List child keys of a parent key

Returns detailed information about a parent key, including its linked child keys.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| Parent Translation Key ID | 
 **optional** | ***KeyLinksIndexOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KeyLinksIndexOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**KeyLink**](KeyLink.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

