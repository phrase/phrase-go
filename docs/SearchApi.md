# \SearchApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchInAccount**](SearchApi.md#SearchInAccount) | **Post** /accounts/{account_id}/search | Search across projects



## SearchInAccount

> AccountSearchResult SearchInAccount(ctx, accountId, searchInAccountParameters, optional)

Search across projects

Search for keys and translations in all account projects <br><br><i>Note: Search is limited to 10000 results and may not include recently updated data depending on the project sizes.</i>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**searchInAccountParameters** | [**SearchInAccountParameters**](SearchInAccountParameters.md)|  | 
 **optional** | ***SearchInAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchInAccountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**AccountSearchResult**](account_search_result.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

