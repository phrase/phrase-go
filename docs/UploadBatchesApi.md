# phrase.UploadBatchesApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadBatchesCreate**](UploadBatchesApi.md#UploadBatchesCreate) | **Post** /projects/{project_id}/upload_batches | Create upload batch



## UploadBatchesCreate

> UploadBatch UploadBatchesCreate(ctx, projectId, uploadBatchesCreateParameters, optional)

Create upload batch

Groups multiple file uploads into a single batch. Optionally, launches the deletion of unmentioned translation keys after all uploads in the batch are completed. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**uploadBatchesCreateParameters** | [**UploadBatchesCreateParameters**](UploadBatchesCreateParameters.md)|  | 
 **optional** | ***UploadBatchesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UploadBatchesCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**UploadBatch**](UploadBatch.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

