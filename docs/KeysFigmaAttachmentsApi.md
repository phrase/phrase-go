# phrase.KeysFigmaAttachmentsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FigmaAttachmentAttachToKey**](KeysFigmaAttachmentsApi.md#FigmaAttachmentAttachToKey) | **Post** /projects/{project_id}/figma_attachments/{figma_attachment_id}/keys | Attach the Figma attachment to a key
[**FigmaAttachmentDetachFromKey**](KeysFigmaAttachmentsApi.md#FigmaAttachmentDetachFromKey) | **Delete** /projects/{project_id}/figma_attachments/{figma_attachment_id}/keys/{id} | Detach the Figma attachment from a key



## FigmaAttachmentAttachToKey

> FigmaAttachmentAttachToKey(ctx, projectId, figmaAttachmentId, id, optional)

Attach the Figma attachment to a key

Attach the Figma attachment to a key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**figmaAttachmentId** | **string**| Figma attachment ID | 
**id** | **string**| ID | 
 **optional** | ***FigmaAttachmentAttachToKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FigmaAttachmentAttachToKeyOpts struct


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


## FigmaAttachmentDetachFromKey

> FigmaAttachmentDetachFromKey(ctx, projectId, figmaAttachmentId, id, optional)

Detach the Figma attachment from a key

Detach the Figma attachment from a key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**figmaAttachmentId** | **string**| Figma attachment ID | 
**id** | **string**| ID | 
 **optional** | ***FigmaAttachmentDetachFromKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FigmaAttachmentDetachFromKeyOpts struct


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

