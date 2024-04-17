# \RepoSyncsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepoSyncActivate**](RepoSyncsApi.md#RepoSyncActivate) | **Post** /accounts/{account_id}/repo_syncs/{repo_sync_id}/activate | Activate a Repo Sync
[**RepoSyncDeactivate**](RepoSyncsApi.md#RepoSyncDeactivate) | **Post** /accounts/{account_id}/repo_syncs/{repo_sync_id}/deactivate | Deactivate a Repo Sync
[**RepoSyncEvents**](RepoSyncsApi.md#RepoSyncEvents) | **Get** /accounts/{account_id}/repo_syncs/{repo_sync_id}/events | Repository Syncs History
[**RepoSyncExport**](RepoSyncsApi.md#RepoSyncExport) | **Post** /accounts/{account_id}/repo_syncs/{repo_sync_id}/export | Export to code repository
[**RepoSyncImport**](RepoSyncsApi.md#RepoSyncImport) | **Post** /accounts/{account_id}/repo_syncs/{repo_sync_id}/import | Import from code repository
[**RepoSyncList**](RepoSyncsApi.md#RepoSyncList) | **Get** /accounts/{account_id}/repo_syncs | Get Repo Syncs
[**RepoSyncShow**](RepoSyncsApi.md#RepoSyncShow) | **Get** /accounts/{account_id}/repo_syncs/{repo_sync_id} | Get a single Repo Sync



## RepoSyncActivate

> RepoSync RepoSyncActivate(ctx, accountId, repoSyncId, optional)

Activate a Repo Sync

Activate a deactivated Repo Sync. Active syncs can be used to import and export translations, and imports to Phrase are automatically triggered by pushes to the repository, if configured.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**repoSyncId** | **string**| Repo Sync ID | 
 **optional** | ***RepoSyncActivateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RepoSyncActivateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**RepoSync**](RepoSync.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepoSyncDeactivate

> RepoSync RepoSyncDeactivate(ctx, accountId, repoSyncId, optional)

Deactivate a Repo Sync

Deactivate an active Repo Sync. Import and export can't be performed on deactivated syncs and the pushes to the repository won't trigger the import to Phrase.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**repoSyncId** | **string**| Repo Sync ID | 
 **optional** | ***RepoSyncDeactivateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RepoSyncDeactivateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**RepoSync**](RepoSync.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepoSyncEvents

> []RepoSyncEvent RepoSyncEvents(ctx, accountId, repoSyncId, optional)

Repository Syncs History

Get the history of a single Repo Sync. The history includes all imports and exports performed by the Repo Sync.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**repoSyncId** | **string**| Repo Sync ID | 
 **optional** | ***RepoSyncEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RepoSyncEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**[]RepoSyncEvent**](RepoSyncEvent.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepoSyncExport

> RepoSyncExport RepoSyncExport(ctx, accountId, repoSyncId, optional)

Export to code repository

> Beta: this feature will change in the future.  Export translations from Phrase Strings to repository provider according to the .phrase.yml file within the code repository.  *Export is done asynchronously and may take several seconds depending on the project size.*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**repoSyncId** | **string**| Repo Sync ID | 
 **optional** | ***RepoSyncExportOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RepoSyncExportOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**RepoSyncExport**](RepoSyncExport.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepoSyncImport

> RepoSyncImport RepoSyncImport(ctx, accountId, repoSyncId, optional)

Import from code repository

> Beta: this feature will change in the future.  Import translations from repository provider to Phrase Strings according to the .phrase.yml file within the code repository.  _Import is done asynchronously and may take several seconds depending on the project size._

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**repoSyncId** | **string**| Repo Sync ID | 
 **optional** | ***RepoSyncImportOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RepoSyncImportOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**RepoSyncImport**](RepoSyncImport.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepoSyncList

> []RepoSync RepoSyncList(ctx, accountId, optional)

Get Repo Syncs

Lists all Repo Syncs from an account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
 **optional** | ***RepoSyncListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RepoSyncListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**[]RepoSync**](RepoSync.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepoSyncShow

> RepoSync RepoSyncShow(ctx, accountId, repoSyncId, optional)

Get a single Repo Sync

Shows a single Repo Sync setting.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**repoSyncId** | **string**| Repo Sync ID | 
 **optional** | ***RepoSyncShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RepoSyncShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**RepoSync**](RepoSync.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

