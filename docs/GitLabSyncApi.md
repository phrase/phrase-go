# \GitLabSyncApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GitlabSyncDelete**](GitLabSyncApi.md#GitlabSyncDelete) | **Delete** /gitlab_syncs/{id} | Delete single Sync Setting
[**GitlabSyncExport**](GitLabSyncApi.md#GitlabSyncExport) | **Post** /gitlab_syncs/{gitlab_sync_id}/export | Export from Phrase to GitLab
[**GitlabSyncHistory**](GitLabSyncApi.md#GitlabSyncHistory) | **Get** /gitlab_syncs/{gitlab_sync_id}/history | History of single Sync Setting
[**GitlabSyncImport**](GitLabSyncApi.md#GitlabSyncImport) | **Post** /gitlab_syncs/{gitlab_sync_id}/import | Import from GitLab to Phrase
[**GitlabSyncList**](GitLabSyncApi.md#GitlabSyncList) | **Get** /gitlab_syncs | List GitLab syncs
[**GitlabSyncShow**](GitLabSyncApi.md#GitlabSyncShow) | **Get** /gitlab_syncs/{id} | Get single Sync Setting
[**GitlabSyncUpdate**](GitLabSyncApi.md#GitlabSyncUpdate) | **Put** /gitlab_syncs/{id} | Update single Sync Setting



## GitlabSyncDelete

> GitlabSyncDelete(ctx, id, optional)

Delete single Sync Setting

Deletes a single GitLab Sync Setting.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| ID | 
 **optional** | ***GitlabSyncDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GitlabSyncDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **accountId** | **optional.String**| Account ID to specify the actual account the GitLab Sync should be created in. Required if the requesting user is a member of multiple accounts. | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitlabSyncExport

> GitlabSyncExport GitlabSyncExport(ctx, gitlabSyncId, gitlabSyncExportParameters, optional)

Export from Phrase to GitLab

Export translations from Phrase to GitLab according to the .phraseapp.yml file within the GitLab repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gitlabSyncId** | **string**| Gitlab Sync ID | 
**gitlabSyncExportParameters** | [**GitlabSyncExportParameters**](GitlabSyncExportParameters.md)|  | 
 **optional** | ***GitlabSyncExportOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GitlabSyncExportOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**GitlabSyncExport**](gitlab_sync_export.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitlabSyncHistory

> []GitlabSyncHistory GitlabSyncHistory(ctx, gitlabSyncId, optional)

History of single Sync Setting

List history for a single Sync Setting.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gitlabSyncId** | **string**| Gitlab Sync ID | 
 **optional** | ***GitlabSyncHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GitlabSyncHistoryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **page** | **optional.Int32**| Page number | 
 **perPage** | **optional.Int32**| allows you to specify a page size up to 100 items, 25 by default | 
 **accountId** | **optional.String**| Account ID to specify the actual account the GitLab Sync should be created in. Required if the requesting user is a member of multiple accounts. | 

### Return type

[**[]GitlabSyncHistory**](gitlab_sync_history.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitlabSyncImport

> []Upload GitlabSyncImport(ctx, gitlabSyncId, gitlabSyncImportParameters, optional)

Import from GitLab to Phrase

Import translations from GitLab to Phrase according to the .phraseapp.yml file within the GitLab repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gitlabSyncId** | **string**| Gitlab Sync ID | 
**gitlabSyncImportParameters** | [**GitlabSyncImportParameters**](GitlabSyncImportParameters.md)|  | 
 **optional** | ***GitlabSyncImportOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GitlabSyncImportOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**[]Upload**](upload.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitlabSyncList

> []GitlabSync GitlabSyncList(ctx, optional)

List GitLab syncs

List all GitLab Sync Settings for which synchronisation with Phrase and GitLab is activated.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GitlabSyncListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GitlabSyncListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **accountId** | **optional.String**| Account ID to specify the actual account the GitLab Sync should be created in. Required if the requesting user is a member of multiple accounts. | 

### Return type

[**[]GitlabSync**](gitlab_sync.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitlabSyncShow

> GitlabSync GitlabSyncShow(ctx, id, optional)

Get single Sync Setting

Shows a single GitLab Sync Setting.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| ID | 
 **optional** | ***GitlabSyncShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GitlabSyncShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **accountId** | **optional.String**| Account ID to specify the actual account the GitLab Sync should be created in. Required if the requesting user is a member of multiple accounts. | 

### Return type

[**GitlabSync**](gitlab_sync.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitlabSyncUpdate

> GitlabSync GitlabSyncUpdate(ctx, id, optional)

Update single Sync Setting

Updates a single GitLab Sync Setting.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| ID | 
 **optional** | ***GitlabSyncUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GitlabSyncUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **accountId** | **optional.String**| Account ID to specify the actual account the GitLab Sync should be created in. Required if the requesting user is a member of multiple accounts. | 
 **phraseProjectCode** | **optional.String**| Code of the related Phrase Project. | 
 **gitlabProjectId** | **optional.Int32**| ID of the related GitLab Project. | 
 **gitlabBranchName** | **optional.String**| Name of the GitLab Branch. | 

### Return type

[**GitlabSync**](gitlab_sync.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

