# phrase.CommentReactionsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReactionCreate**](CommentReactionsApi.md#ReactionCreate) | **Post** /projects/{project_id}/keys/{key_id}/comments/{comment_id}/reactions | Create a reaction
[**ReactionDelete**](CommentReactionsApi.md#ReactionDelete) | **Delete** /projects/{project_id}/keys/{key_id}/comments/{comment_id}/reactions/{id} | Delete a reaction
[**ReactionShow**](CommentReactionsApi.md#ReactionShow) | **Get** /projects/{project_id}/keys/{key_id}/comments/{comment_id}/reactions/{id} | Get a single reaction
[**ReactionsList**](CommentReactionsApi.md#ReactionsList) | **Get** /projects/{project_id}/keys/{key_id}/comments/{comment_id}/reactions | List reactions



## ReactionCreate

> CommentReaction ReactionCreate(ctx, projectId, keyId, commentId, optional)

Create a reaction

Create a new reaction for a comment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keyId** | **string**| Translation Key ID | 
**commentId** | **string**| Comment ID | 
 **optional** | ***ReactionCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReactionCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**branch** | **optional.String**| specify the branch to use | 
**emoji** | **optional.String**| specify the emoji for the reaction | 

### Return type

[**CommentReaction**](CommentReaction.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactionDelete

> ReactionDelete(ctx, projectId, keyId, commentId, id, optional)

Delete a reaction

Delete an existing reaction.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keyId** | **string**| Translation Key ID | 
**commentId** | **string**| Comment ID | 
**id** | **string**| ID | 
 **optional** | ***ReactionDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReactionDeleteOpts struct


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


## ReactionShow

> CommentReaction ReactionShow(ctx, projectId, keyId, commentId, id, optional)

Get a single reaction

Get details on a single reaction.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keyId** | **string**| Translation Key ID | 
**commentId** | **string**| Comment ID | 
**id** | **string**| ID | 
 **optional** | ***ReactionShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReactionShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**branch** | **optional.String**| specify the branch to use | 

### Return type

[**CommentReaction**](CommentReaction.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactionsList

> []CommentReaction ReactionsList(ctx, projectId, keyId, commentId, optional)

List reactions

List all reactions for a comment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keyId** | **string**| Translation Key ID | 
**commentId** | **string**| Comment ID | 
 **optional** | ***ReactionsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReactionsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**page** | **optional.Int32**| Page number | 
**perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 
**branch** | **optional.String**| specify the branch to use | 

### Return type

[**[]CommentReaction**](CommentReaction.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

