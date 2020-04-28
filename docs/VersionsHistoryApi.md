# \VersionsHistoryApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VersionShow**](VersionsHistoryApi.md#VersionShow) | **Get** /projects/{project_id}/translations/{translation_id}/versions/{id} | Get a single version
[**VersionsList**](VersionsHistoryApi.md#VersionsList) | **Get** /projects/{project_id}/translations/{translation_id}/versions | List all versions



## VersionShow

> TranslationVersionWithUser VersionShow(ctx, projectId, translationId, id, versionShowParameters, optional)

Get a single version

Get details on a single version.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**translationId** | **string**| Translation ID | 
**id** | **string**| ID | 
**versionShowParameters** | [**VersionShowParameters**](VersionShowParameters.md)|  | 
 **optional** | ***VersionShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VersionShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**TranslationVersionWithUser**](translation_version_with_user.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VersionsList

> []TranslationVersion VersionsList(ctx, projectId, translationId, versionsListParameters, optional)

List all versions

List all versions for the given translation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**translationId** | **string**| Translation ID | 
**versionsListParameters** | [**VersionsListParameters**](VersionsListParameters.md)|  | 
 **optional** | ***VersionsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VersionsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **page** | **optional.Int32**| Page number | 
 **perPage** | **optional.Int32**| allows you to specify a page size up to 100 items, 10 by default | 

### Return type

[**[]TranslationVersion**](translation_version.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

