# phrase.ChecksApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckIssueDismiss**](ChecksApi.md#CheckIssueDismiss) | **Patch** /projects/{project_id}/checks/issues/{id}/dismiss | Dismiss a check issue
[**CheckIssuesList**](ChecksApi.md#CheckIssuesList) | **Get** /projects/{project_id}/checks/issues | List check issues



## CheckIssueDismiss

> CheckIssue CheckIssueDismiss(ctx, projectId, id, optional)

Dismiss a check issue

**Note:** The Checks API is still in development and might change in subsequent releases.  Mark a check issue as dismissed so it no longer appears on the list of active check issues.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| Check Issue ID | 
 **optional** | ***CheckIssueDismissOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CheckIssueDismissOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**CheckIssue**](CheckIssue.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckIssuesList

> []CheckIssue CheckIssuesList(ctx, projectId, optional)

List check issues

**Note:** The Checks API is still in development and might change in subsequent releases.  List check issues for the given project. Results can be filtered by locale, check name, and state.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
 **optional** | ***CheckIssuesListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CheckIssuesListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**page** | **optional.Int32**| Page number | 
**perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 
**state** | **optional.String**| Filter by state of the check issue. Can be one of: &#x60;active&#x60;, &#x60;solved&#x60;, &#x60;dismissed&#x60;, &#x60;all&#x60;. Defaults to &#x60;active&#x60;. | [default to &quot;active&quot;]
**localeIds** | [**optional.Interface of []string**](string.md)| Filter by one or more locale IDs. | 
**checkNames** | [**optional.Interface of []string**](string.md)| Filter by one or more check names. Valid values are:  - &#x60;translation_content_length&#x60; — the translation exceeds the maximum character limit configured for the key. - &#x60;translation_placeholder_usage&#x60; — the translation is missing placeholders present in the source, or contains unexpected ones. - &#x60;translation_glossary_usage&#x60; — the translation does not follow the glossary term translations. | 

### Return type

[**[]CheckIssue**](CheckIssue.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

