# phrase.KeyFormatAnnotationsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KeyFormatAnnotationsList**](KeyFormatAnnotationsApi.md#KeyFormatAnnotationsList) | **Get** /projects/{project_id}/keys/{id}/format_annotations | List format annotations for a key



## KeyFormatAnnotationsList

> []KeyFormatAnnotationsList200ResponseInner KeyFormatAnnotationsList(ctx, projectId, id, optional)

List format annotations for a key

Returns the format annotations stored on a translation key. Format annotations capture file-format data recorded when the key was imported — for example, an ARB placeholder block or an XLIFF note.  Results are limited to 1,000 entries. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
 **optional** | ***KeyFormatAnnotationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KeyFormatAnnotationsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**branch** | **optional.String**| Branch to use | 

### Return type

[**[]KeyFormatAnnotationsList200ResponseInner**](KeyFormatAnnotationsList200ResponseInner.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

