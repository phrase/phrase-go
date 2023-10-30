# \CommentRepliesApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepliesList**](CommentRepliesApi.md#RepliesList) | **Get** /projects/{project_id}/keys/{key_id}/comments/{comment_id}/replies | List replies
[**ReplyCreate**](CommentRepliesApi.md#ReplyCreate) | **Post** /projects/{project_id}/keys/{key_id}/comments/{comment_id}/replies | Create a reply
[**ReplyDelete**](CommentRepliesApi.md#ReplyDelete) | **Delete** /projects/{project_id}/keys/{key_id}/comments/{comment_id}/replies/{id} | Delete a reply
[**ReplyMarkAsRead**](CommentRepliesApi.md#ReplyMarkAsRead) | **Patch** /projects/{project_id}/keys/{key_id}/comments/{comment_id}/replies/{id}/mark_as_read | Mark a reply as read
[**ReplyMarkAsUnread**](CommentRepliesApi.md#ReplyMarkAsUnread) | **Patch** /projects/{project_id}/keys/{key_id}/comments/{comment_id}/replies/{id}/mark_as_unread | Mark a reply as unread
[**ReplyShow**](CommentRepliesApi.md#ReplyShow) | **Get** /projects/{project_id}/keys/{key_id}/comments/{comment_id}/replies/{id} | Get a single reply



## RepliesList

> []Comment RepliesList(ctx, projectId, keyId, commentId, repliesListParameters, optional)

List replies

List all replies for a comment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keyId** | **string**| Translation Key ID | 
**commentId** | **string**| Comment ID | 
**repliesListParameters** | [**RepliesListParameters**](RepliesListParameters.md)|  | 
 **optional** | ***RepliesListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RepliesListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **page** | **optional.Int32**| Page number | 
 **perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 
 **branch** | **optional.String**| specify the branch to use | 
 **query** | **optional.String**| Search query for comment messages | 
 **filters** | [**optional.Interface of []string**](string.md)| Specify the filter for the comments | 
 **order** | **optional.String**| Order direction. Can be one of: asc, desc. | 

### Return type

[**[]Comment**](Comment.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplyCreate

> Comment ReplyCreate(ctx, projectId, keyId, commentId, optional)

Create a reply

Create a new reply for a comment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keyId** | **string**| Translation Key ID | 
**commentId** | **string**| Comment ID | 
 **optional** | ***ReplyCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReplyCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **branch** | **optional.String**| specify the branch to use | 
 **message** | **optional.String**| specify the message for the comment | 

### Return type

[**Comment**](Comment.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplyDelete

> ReplyDelete(ctx, projectId, keyId, commentId, id, optional)

Delete a reply

Delete an existing reply.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keyId** | **string**| Translation Key ID | 
**commentId** | **string**| Comment ID | 
**id** | **string**| ID | 
 **optional** | ***ReplyDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReplyDeleteOpts struct


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


## ReplyMarkAsRead

> ReplyMarkAsRead(ctx, projectId, keyId, commentId, id, optional)

Mark a reply as read

Mark a reply as read.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keyId** | **string**| Translation Key ID | 
**commentId** | **string**| Comment ID | 
**id** | **string**| ID | 
 **optional** | ***ReplyMarkAsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReplyMarkAsReadOpts struct


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


## ReplyMarkAsUnread

> ReplyMarkAsUnread(ctx, projectId, keyId, commentId, id, optional)

Mark a reply as unread

Mark a reply as unread.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keyId** | **string**| Translation Key ID | 
**commentId** | **string**| Comment ID | 
**id** | **string**| ID | 
 **optional** | ***ReplyMarkAsUnreadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReplyMarkAsUnreadOpts struct


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


## ReplyShow

> Comment ReplyShow(ctx, projectId, keyId, commentId, id, optional)

Get a single reply

Get details on a single reply.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keyId** | **string**| Translation Key ID | 
**commentId** | **string**| Comment ID | 
**id** | **string**| ID | 
 **optional** | ***ReplyShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReplyShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **branch** | **optional.String**| specify the branch to use | 

### Return type

[**Comment**](Comment.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

