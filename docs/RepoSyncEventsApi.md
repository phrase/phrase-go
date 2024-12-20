# \RepoSyncEventsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepoSyncEventList**](RepoSyncEventsApi.md#RepoSyncEventList) | **Get** /accounts/{account_id}/repo_syncs/{id}/events | Repository Syncs History
[**RepoSyncEventShow**](RepoSyncEventsApi.md#RepoSyncEventShow) | **Get** /accounts/{account_id}/repo_syncs/{repo_sync_id}/events/{id} | Get a single Repo Sync Event



## RepoSyncEventList

> []RepoSyncEvent RepoSyncEventList(ctx, accountId, id, optional)

Repository Syncs History

Get the history of a single Repo Sync. The history includes all imports and exports performed by the Repo Sync.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**id** | **string**| ID | 
 **optional** | ***RepoSyncEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RepoSyncEventListOpts struct


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


## RepoSyncEventShow

> RepoSyncEvent RepoSyncEventShow(ctx, accountId, repoSyncId, id, optional)

Get a single Repo Sync Event

Shows a single Repo Sync event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**repoSyncId** | **string**| Repo Sync ID | 
**id** | **string**| ID | 
 **optional** | ***RepoSyncEventShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RepoSyncEventShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**RepoSyncEvent**](RepoSyncEvent.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

