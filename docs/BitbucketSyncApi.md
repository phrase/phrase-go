# \BitbucketSyncApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BitbucketSyncExport**](BitbucketSyncApi.md#BitbucketSyncExport) | **Post** /bitbucket_syncs/{id}/export | Export from Phrase to Bitbucket
[**BitbucketSyncImport**](BitbucketSyncApi.md#BitbucketSyncImport) | **Post** /bitbucket_syncs/{id}/import | Import to Phrase from Bitbucket
[**BitbucketSyncsList**](BitbucketSyncApi.md#BitbucketSyncsList) | **Get** /bitbucket_syncs | List Bitbucket syncs



## BitbucketSyncExport

> BitbucketSyncExportResponse BitbucketSyncExport(ctx, id, bitbucketSyncExportParameters, optional)

Export from Phrase to Bitbucket

Export translations from Phrase to Bitbucket according to the .phraseapp.yml file within the Bitbucket Repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| ID | 
**bitbucketSyncExportParameters** | [**BitbucketSyncExportParameters**](BitbucketSyncExportParameters.md)|  | 
 **optional** | ***BitbucketSyncExportOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BitbucketSyncExportOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**BitbucketSyncExportResponse**](bitbucket_sync_export_response.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BitbucketSyncImport

> BitbucketSyncImport(ctx, id, bitbucketSyncImportParameters, optional)

Import to Phrase from Bitbucket

Import translations from Bitbucket to Phrase according to the .phraseapp.yml file within the Bitbucket repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| ID | 
**bitbucketSyncImportParameters** | [**BitbucketSyncImportParameters**](BitbucketSyncImportParameters.md)|  | 
 **optional** | ***BitbucketSyncImportOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BitbucketSyncImportOpts struct


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


## BitbucketSyncsList

> []BitbucketSync BitbucketSyncsList(ctx, optional)

List Bitbucket syncs

List all Bitbucket repositories for which synchronisation with Phrase is activated.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BitbucketSyncsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BitbucketSyncsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **accountId** | **optional.String**| Account ID to specify the actual account the project should be created in. Required if the requesting user is a member of multiple accounts. | 

### Return type

[**[]BitbucketSync**](bitbucket_sync.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

