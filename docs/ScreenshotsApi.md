# phrase.ScreenshotsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScreenshotCreate**](ScreenshotsApi.md#ScreenshotCreate) | **Post** /projects/{project_id}/screenshots | Create a screenshot
[**ScreenshotDelete**](ScreenshotsApi.md#ScreenshotDelete) | **Delete** /projects/{project_id}/screenshots/{id} | Delete a screenshot
[**ScreenshotShow**](ScreenshotsApi.md#ScreenshotShow) | **Get** /projects/{project_id}/screenshots/{id} | Get a single screenshot
[**ScreenshotUpdate**](ScreenshotsApi.md#ScreenshotUpdate) | **Patch** /projects/{project_id}/screenshots/{id} | Update a screenshot
[**ScreenshotsList**](ScreenshotsApi.md#ScreenshotsList) | **Get** /projects/{project_id}/screenshots | List screenshots



## ScreenshotCreate

> Screenshot ScreenshotCreate(ctx, projectId, filename, optional)

Create a screenshot

Creates a screenshot in a project to provide visual context for in-context translation. Attach translation keys to regions of the uploaded image so translators can see where each string appears in your UI.  This endpoint accepts a multipart/form-data request with a binary file upload, unlike most Phrase API endpoints that use JSON. Use a multipart form client or the -F flag in curl rather than a JSON body.  The screenshot name must be unique within the project (case-insensitive). When name is omitted, it is derived from the uploaded filename. The account must have the Screenshots feature enabled; requests to projects on accounts without it return 403. Creating a screenshot requires a token with the write scope and manage access to the project. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**filename** | ***os.File*****os.File**| Image file to upload. Accepted formats are JPEG (jpg/jpeg), GIF, and PNG. Maximum file size is 10 MB. Submitting an unsupported format or a file exceeding the size limit returns 422. | 
 **optional** | ***ScreenshotCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScreenshotCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**branch** | **optional.String**| specify the branch to use | 
**name** | **optional.String**| Display name for the screenshot. Must be unique within the project (case-insensitive). When omitted, the name is derived from the uploaded filename. | 
**description** | **optional.String**| Optional free-text description of the screenshot. | 

### Return type

[**Screenshot**](Screenshot.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScreenshotDelete

> ScreenshotDelete(ctx, projectId, id, optional)

Delete a screenshot

Permanently removes a screenshot and all its associated markers from the project. Use this when you need to fully remove a screenshot that is no longer relevant — for example, after a UI redesign renders the captured screen obsolete. This is a hard delete: the screenshot record and every key-to-region marker linked to it are destroyed together and cannot be recovered. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
 **optional** | ***ScreenshotDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScreenshotDeleteOpts struct


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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScreenshotShow

> Screenshot ScreenshotShow(ctx, projectId, id, optional)

Get a single screenshot

Returns a single screenshot belonging to the specified project. Use this to retrieve the screenshot's name, description, hosted image URL, and marker count after uploading, or before creating, updating, or inspecting its markers. The response is a synchronous, idempotent read — repeated calls return the same record without side effects.  The Attachable Screenshots feature must be enabled on the account. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
 **optional** | ***ScreenshotShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScreenshotShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**branch** | **optional.String**| specify the branch to use | 

### Return type

[**Screenshot**](Screenshot.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScreenshotUpdate

> Screenshot ScreenshotUpdate(ctx, projectId, id, screenshotUpdateParameters, optional)

Update a screenshot

Update an existing screenshot.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
**screenshotUpdateParameters** | [**ScreenshotUpdateParameters**](ScreenshotUpdateParameters.md)|  | 
 **optional** | ***ScreenshotUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScreenshotUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**Screenshot**](Screenshot.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScreenshotsList

> []Screenshot ScreenshotsList(ctx, projectId, optional)

List screenshots

List all screenshots for the given project.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
 **optional** | ***ScreenshotsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScreenshotsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**page** | **optional.Int32**| Page number | 
**perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 
**branch** | **optional.String**| specify the branch to use | 
**keyId** | **optional.String**| filter by key | 

### Return type

[**[]Screenshot**](Screenshot.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

