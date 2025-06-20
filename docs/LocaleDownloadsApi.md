# phrase.LocaleDownloadsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocaleDownloadCreate**](LocaleDownloadsApi.md#LocaleDownloadCreate) | **Post** /projects/{project_id}/locales/{locale_id}/downloads | Initiate async download of a locale
[**LocaleDownloadShow**](LocaleDownloadsApi.md#LocaleDownloadShow) | **Get** /projects/{project_id}/locales/{locale_id}/downloads/{id} | Show status of an async locale download



## LocaleDownloadCreate

> LocaleDownload LocaleDownloadCreate(ctx, projectId, localeId, localeDownloadCreateParameters, optional)

Initiate async download of a locale

Prepare a locale for download in a specific file format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**localeId** | **string**| Locale ID | 
**localeDownloadCreateParameters** | [**LocaleDownloadCreateParameters**](LocaleDownloadCreateParameters.md)|  | 
 **optional** | ***LocaleDownloadCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LocaleDownloadCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**ifModifiedSince** | **optional.String**| Last modified condition, see [Conditional GET requests / HTTP Caching](/en/api/strings/pagination#conditional-get-requests-%2F-http-caching) (optional) | 
**ifNoneMatch** | **optional.String**| ETag condition, see [Conditional GET requests / HTTP Caching](/en/api/strings/pagination#conditional-get-requests-%2F-http-caching) (optional) | 

### Return type

[**LocaleDownload**](LocaleDownload.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocaleDownloadShow

> LocaleDownload LocaleDownloadShow(ctx, projectId, localeId, id, optional)

Show status of an async locale download

Show status of already started async locale download. If the download is finished, the download link will be returned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**localeId** | **string**| Locale ID | 
**id** | **string**| ID | 
 **optional** | ***LocaleDownloadShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LocaleDownloadShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**ifModifiedSince** | **optional.String**| Last modified condition, see [Conditional GET requests / HTTP Caching](/en/api/strings/pagination#conditional-get-requests-%2F-http-caching) (optional) | 
**ifNoneMatch** | **optional.String**| ETag condition, see [Conditional GET requests / HTTP Caching](/en/api/strings/pagination#conditional-get-requests-%2F-http-caching) (optional) | 

### Return type

[**LocaleDownload**](LocaleDownload.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

