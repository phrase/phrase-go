# phrase.DocumentsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DocumentDelete**](DocumentsApi.md#DocumentDelete) | **Delete** /projects/{project_id}/documents/{id} | Delete document
[**DocumentsList**](DocumentsApi.md#DocumentsList) | **Get** /projects/{project_id}/documents | List documents



## DocumentDelete

> DocumentDelete(ctx, projectId, id, optional)

Delete document

Delete an existing document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
 **optional** | ***DocumentDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DocumentDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

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


## DocumentsList

> []Document DocumentsList(ctx, projectId, optional)

List documents

Returns all documents in a project that the authenticated user has read access to. A Document is a source file — an HTML or DOCX file — that has been uploaded to Phrase Strings and whose content is segmented into translation keys for localization.  Use this endpoint to enumerate documents before downloading, previewing, or triggering translation workflows for individual files.  The q parameter performs a prefix match on the document name (case-insensitive). For example, passing q=invoice returns documents whose names begin with \"invoice\" but not documents containing \"invoice\" elsewhere in the name. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
 **optional** | ***DocumentsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DocumentsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**page** | **optional.Int32**| Page number | 
**perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 
**q** | **optional.String**| Filter documents by name prefix. Returns documents whose name starts with the given value (case-insensitive). | 

### Return type

[**[]Document**](Document.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

