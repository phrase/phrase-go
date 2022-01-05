# \CommentsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommentCreate**](CommentsApi.md#CommentCreate) | **Post** /projects/{project_id}/keys/{key_id}/comments | Create a comment
[**CommentDelete**](CommentsApi.md#CommentDelete) | **Delete** /projects/{project_id}/keys/{key_id}/comments/{id} | Delete a comment
[**CommentMarkCheck**](CommentsApi.md#CommentMarkCheck) | **Get** /projects/{project_id}/keys/{key_id}/comments/{id}/read | Check if comment is read
[**CommentMarkRead**](CommentsApi.md#CommentMarkRead) | **Patch** /projects/{project_id}/keys/{key_id}/comments/{id}/read | Mark a comment as read
[**CommentMarkUnread**](CommentsApi.md#CommentMarkUnread) | **Delete** /projects/{project_id}/keys/{key_id}/comments/{id}/read | Mark a comment as unread
[**CommentShow**](CommentsApi.md#CommentShow) | **Get** /projects/{project_id}/keys/{key_id}/comments/{id} | Get a single comment
[**CommentUpdate**](CommentsApi.md#CommentUpdate) | **Patch** /projects/{project_id}/keys/{key_id}/comments/{id} | Update a comment
[**CommentsList**](CommentsApi.md#CommentsList) | **Get** /projects/{project_id}/keys/{key_id}/comments | List comments



## CommentCreate

> Comment CommentCreate(ctx, projectId, keyId, commentCreateParameters, optional)

Create a comment

Create a new comment for a key.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keyId** | **string**| Translation Key ID | 
**commentCreateParameters** | [**CommentCreateParameters**](CommentCreateParameters.md)|  | 
 **optional** | ***CommentCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommentCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**Comment**](comment.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommentDelete

> CommentDelete(ctx, projectId, keyId, id, optional)

Delete a comment

Delete an existing comment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keyId** | **string**| Translation Key ID | 
**id** | **string**| ID | 
 **optional** | ***CommentDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommentDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **branch** | **optional.String**| specify the branch to use | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommentMarkCheck

> CommentMarkCheck(ctx, projectId, keyId, id, optional)

Check if comment is read

Check if comment was marked as read. Returns 204 if read, 404 if unread.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keyId** | **string**| Translation Key ID | 
**id** | **string**| ID | 
 **optional** | ***CommentMarkCheckOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommentMarkCheckOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **branch** | **optional.String**| specify the branch to use | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommentMarkRead

> CommentMarkRead(ctx, projectId, keyId, id, commentMarkReadParameters, optional)

Mark a comment as read

Mark a comment as read.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keyId** | **string**| Translation Key ID | 
**id** | **string**| ID | 
**commentMarkReadParameters** | [**CommentMarkReadParameters**](CommentMarkReadParameters.md)|  | 
 **optional** | ***CommentMarkReadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommentMarkReadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommentMarkUnread

> CommentMarkUnread(ctx, projectId, keyId, id, optional)

Mark a comment as unread

Mark a comment as unread.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keyId** | **string**| Translation Key ID | 
**id** | **string**| ID | 
 **optional** | ***CommentMarkUnreadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommentMarkUnreadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **branch** | **optional.String**| specify the branch to use | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommentShow

> Comment CommentShow(ctx, projectId, keyId, id, optional)

Get a single comment

Get details on a single comment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keyId** | **string**| Translation Key ID | 
**id** | **string**| ID | 
 **optional** | ***CommentShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommentShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **branch** | **optional.String**| specify the branch to use | 

### Return type

[**Comment**](comment.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommentUpdate

> Comment CommentUpdate(ctx, projectId, keyId, id, commentUpdateParameters, optional)

Update a comment

Update an existing comment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keyId** | **string**| Translation Key ID | 
**id** | **string**| ID | 
**commentUpdateParameters** | [**CommentUpdateParameters**](CommentUpdateParameters.md)|  | 
 **optional** | ***CommentUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommentUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**Comment**](comment.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommentsList

> []Comment CommentsList(ctx, projectId, keyId, optional)

List comments

List all comments for a key.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keyId** | **string**| Translation Key ID | 
 **optional** | ***CommentsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommentsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **page** | **optional.Int32**| Page number | 
 **perPage** | **optional.Int32**| allows you to specify a page size up to 100 items, 25 by default | 
 **branch** | **optional.String**| specify the branch to use | 

### Return type

[**[]Comment**](comment.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

