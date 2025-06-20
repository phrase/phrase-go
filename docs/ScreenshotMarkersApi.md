# phrase.ScreenshotMarkersApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScreenshotMarkerCreate**](ScreenshotMarkersApi.md#ScreenshotMarkerCreate) | **Post** /projects/{project_id}/screenshots/{screenshot_id}/markers | Create a screenshot marker
[**ScreenshotMarkerDelete**](ScreenshotMarkersApi.md#ScreenshotMarkerDelete) | **Delete** /projects/{project_id}/screenshots/{screenshot_id}/markers | Delete a screenshot marker
[**ScreenshotMarkerShow**](ScreenshotMarkersApi.md#ScreenshotMarkerShow) | **Get** /projects/{project_id}/screenshots/{screenshot_id}/markers/{id} | Get a single screenshot marker
[**ScreenshotMarkerUpdate**](ScreenshotMarkersApi.md#ScreenshotMarkerUpdate) | **Patch** /projects/{project_id}/screenshots/{screenshot_id}/markers | Update a screenshot marker
[**ScreenshotMarkersList**](ScreenshotMarkersApi.md#ScreenshotMarkersList) | **Get** /projects/{project_id}/screenshots/{id}/markers | List screenshot markers



## ScreenshotMarkerCreate

> ScreenshotMarker ScreenshotMarkerCreate(ctx, projectId, screenshotId, screenshotMarkerCreateParameters, optional)

Create a screenshot marker

Create a new screenshot marker.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**screenshotId** | **string**| Screenshot ID | 
**screenshotMarkerCreateParameters** | [**ScreenshotMarkerCreateParameters**](ScreenshotMarkerCreateParameters.md)|  | 
 **optional** | ***ScreenshotMarkerCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScreenshotMarkerCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**ScreenshotMarker**](ScreenshotMarker.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScreenshotMarkerDelete

> ScreenshotMarkerDelete(ctx, projectId, screenshotId, optional)

Delete a screenshot marker

Delete an existing screenshot marker.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**screenshotId** | **string**| Screenshot ID | 
 **optional** | ***ScreenshotMarkerDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScreenshotMarkerDeleteOpts struct


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


## ScreenshotMarkerShow

> ScreenshotMarker ScreenshotMarkerShow(ctx, projectId, screenshotId, id, optional)

Get a single screenshot marker

Get details on a single screenshot marker for a given project.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**screenshotId** | **string**| Screenshot ID | 
**id** | **string**| ID | 
 **optional** | ***ScreenshotMarkerShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScreenshotMarkerShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**branch** | **optional.String**| specify the branch to use | 

### Return type

[**ScreenshotMarker**](ScreenshotMarker.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScreenshotMarkerUpdate

> ScreenshotMarker ScreenshotMarkerUpdate(ctx, projectId, screenshotId, screenshotMarkerUpdateParameters, optional)

Update a screenshot marker

Update an existing screenshot marker.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**screenshotId** | **string**| Screenshot ID | 
**screenshotMarkerUpdateParameters** | [**ScreenshotMarkerUpdateParameters**](ScreenshotMarkerUpdateParameters.md)|  | 
 **optional** | ***ScreenshotMarkerUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScreenshotMarkerUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**ScreenshotMarker**](ScreenshotMarker.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScreenshotMarkersList

> []ScreenshotMarker ScreenshotMarkersList(ctx, projectId, id, optional)

List screenshot markers

List all screenshot markers for the given project.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
 **optional** | ***ScreenshotMarkersListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScreenshotMarkersListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**page** | **optional.Int32**| Page number | 
**perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 
**branch** | **optional.String**| specify the branch to use | 

### Return type

[**[]ScreenshotMarker**](ScreenshotMarker.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

