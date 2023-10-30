# \VersionsHistoryApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VersionShow**](VersionsHistoryApi.md#VersionShow) | **Get** /projects/{project_id}/translations/{translation_id}/versions/{id} | Get a single version
[**VersionsList**](VersionsHistoryApi.md#VersionsList) | **Get** /projects/{project_id}/translations/{translation_id}/versions | List all versions



## VersionShow

> TranslationVersionWithUser VersionShow(ctx, projectId, translationId, id, optional)

Get a single version

Get details on a single version.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**translationId** | **string**| Translation ID | 
**id** | **string**| ID | 
 **optional** | ***VersionShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VersionShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **branch** | **optional.String**| specify the branch to use | 

### Return type

[**TranslationVersionWithUser**](TranslationVersionWithUser.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VersionsList

> []TranslationVersion VersionsList(ctx, projectId, translationId, optional)

List all versions

List all changes done to a given translation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**translationId** | **string**| Translation ID | 
 **optional** | ***VersionsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VersionsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **page** | **optional.Int32**| Page number | 
 **perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 
 **branch** | **optional.String**| specify the branch to use | 

### Return type

[**[]TranslationVersion**](TranslationVersion.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

