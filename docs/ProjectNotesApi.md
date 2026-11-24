# phrase.ProjectNotesApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectNoteDestroy**](ProjectNotesApi.md#ProjectNoteDestroy) | **Delete** /projects/{project_id}/note | Delete a project note
[**ProjectNoteShow**](ProjectNotesApi.md#ProjectNoteShow) | **Get** /projects/{project_id}/note | Get a project note
[**ProjectNoteUpdate**](ProjectNotesApi.md#ProjectNoteUpdate) | **Patch** /projects/{project_id}/note | Create or update a project note



## ProjectNoteDestroy

> ProjectNoteDestroy(ctx, projectId, optional)

Delete a project note

Deletes the note associated with a project. This operation is idempotent — it succeeds silently even when no note exists. The note is resolved against the main project — branch projects delete the main project's note. Requires manage permissions on the project and the `write` OAuth scope. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
 **optional** | ***ProjectNoteDestroyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProjectNoteDestroyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**branch** | **optional.String**| Branch to use | 

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


## ProjectNoteShow

> Note ProjectNoteShow(ctx, projectId, optional)

Get a project note

Returns the Markdown-formatted note attached to a project. The note is resolved against the main project — branch projects return the main project's note. Returns 404 when no note has been created yet. Requires the `read` OAuth scope. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
 **optional** | ***ProjectNoteShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProjectNoteShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**branch** | **optional.String**| Branch to use | 

### Return type

[**Note**](Note.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectNoteUpdate

> Note ProjectNoteUpdate(ctx, projectId, projectNoteUpdateParameters, optional)

Create or update a project note

Creates or updates the Markdown-formatted note for a project. If a note already exists it is updated in place; otherwise a new note is created. The note is stored on the main project — branch projects write to the main project's note. `body` is required. Requires manage permissions on the project and the `write` OAuth scope. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**projectNoteUpdateParameters** | [**ProjectNoteUpdateParameters**](ProjectNoteUpdateParameters.md)|  | 
 **optional** | ***ProjectNoteUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProjectNoteUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**branch** | **optional.String**| Branch to use | 

### Return type

[**Note**](Note.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

