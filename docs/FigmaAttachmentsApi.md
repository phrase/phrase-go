# \FigmaAttachmentsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FigmaAttachmentCreate**](FigmaAttachmentsApi.md#FigmaAttachmentCreate) | **Post** /projects/{project_id}/figma_attachments | Create a Figma attachment
[**FigmaAttachmentDelete**](FigmaAttachmentsApi.md#FigmaAttachmentDelete) | **Delete** /projects/{project_id}/figma_attachments/{id} | Delete a Figma attachment
[**FigmaAttachmentShow**](FigmaAttachmentsApi.md#FigmaAttachmentShow) | **Get** /projects/{project_id}/figma_attachments/{id} | Get a single Figma attachment
[**FigmaAttachmentUpdate**](FigmaAttachmentsApi.md#FigmaAttachmentUpdate) | **Patch** /projects/{project_id}/figma_attachments/{id} | Update a Figma attachment
[**FigmaAttachmentsList**](FigmaAttachmentsApi.md#FigmaAttachmentsList) | **Get** /projects/{project_id}/figma_attachments | List Figma attachments



## FigmaAttachmentCreate

> FigmaAttachment FigmaAttachmentCreate(ctx, projectId, figmaAttachmentCreateParameters, optional)

Create a Figma attachment

Create a new Figma attachment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**figmaAttachmentCreateParameters** | [**FigmaAttachmentCreateParameters**](FigmaAttachmentCreateParameters.md)|  | 
 **optional** | ***FigmaAttachmentCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FigmaAttachmentCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **branch** | **optional.String**| specify the branch to use | 

### Return type

[**FigmaAttachment**](figma_attachment.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FigmaAttachmentDelete

> FigmaAttachmentDelete(ctx, projectId, id, optional)

Delete a Figma attachment

Delete an existing Figma attachment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
 **optional** | ***FigmaAttachmentDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FigmaAttachmentDeleteOpts struct


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


## FigmaAttachmentShow

> FigmaAttachment FigmaAttachmentShow(ctx, projectId, id, optional)

Get a single Figma attachment

Get details on a single Figma attachment for a given project.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
 **optional** | ***FigmaAttachmentShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FigmaAttachmentShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **branch** | **optional.String**| specify the branch to use | 

### Return type

[**FigmaAttachment**](figma_attachment.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FigmaAttachmentUpdate

> FigmaAttachment FigmaAttachmentUpdate(ctx, projectId, id, figmaAttachmentUpdateParameters, optional)

Update a Figma attachment

Update an existing Figma attachment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
**figmaAttachmentUpdateParameters** | [**FigmaAttachmentUpdateParameters**](FigmaAttachmentUpdateParameters.md)|  | 
 **optional** | ***FigmaAttachmentUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FigmaAttachmentUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **branch** | **optional.String**| specify the branch to use | 

### Return type

[**FigmaAttachment**](figma_attachment.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FigmaAttachmentsList

> []FigmaAttachment FigmaAttachmentsList(ctx, projectId, optional)

List Figma attachments

List all Figma attachments for the given project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
 **optional** | ***FigmaAttachmentsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FigmaAttachmentsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **page** | **optional.Int32**| Page number | 
 **perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 
 **branch** | **optional.String**| specify the branch to use | 

### Return type

[**[]FigmaAttachment**](figma_attachment.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

