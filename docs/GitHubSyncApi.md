# \GitHubSyncApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GithubSyncExport**](GitHubSyncApi.md#GithubSyncExport) | **Post** /github_syncs/export | Export from Phrase Strings to GitHub
[**GithubSyncImport**](GitHubSyncApi.md#GithubSyncImport) | **Post** /github_syncs/import | Import to Phrase Strings from GitHub



## GithubSyncExport

> GithubSyncExport(ctx, githubSyncExportParameters, optional)

Export from Phrase Strings to GitHub

Export translations from Phrase Strings to GitHub according to the .phraseapp.yml file within the GitHub repository. <br><br><i>Note: Export is done asynchronously and may take several seconds depending on the project size.</i>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**githubSyncExportParameters** | [**GithubSyncExportParameters**](GithubSyncExportParameters.md)|  | 
 **optional** | ***GithubSyncExportOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GithubSyncExportOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GithubSyncImport

> GithubSyncImport(ctx, githubSyncImportParameters, optional)

Import to Phrase Strings from GitHub

Import files to Phrase Strings from your connected GitHub repository. <br><br><i>Note: Import is done asynchronously and may take several seconds depending on the project size.</i>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**githubSyncImportParameters** | [**GithubSyncImportParameters**](GithubSyncImportParameters.md)|  | 
 **optional** | ***GithubSyncImportOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GithubSyncImportOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

