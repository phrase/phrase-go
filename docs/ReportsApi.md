# phrase.ReportsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportLocalesList**](ReportsApi.md#ReportLocalesList) | **Get** /projects/{project_id}/report/locales | List Locale Reports
[**ReportShow**](ReportsApi.md#ReportShow) | **Get** /projects/{project_id}/report | Get Project Report



## ReportLocalesList

> []LocaleReport ReportLocalesList(ctx, projectId, optional)

List Locale Reports

List all locale reports for the given project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
 **optional** | ***ReportLocalesListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReportLocalesListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**page** | **optional.Int32**| Page number | 
**perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 
**localeCodes** | **optional.String**| Locale Code | 
**tag** | **optional.String**| tag | 
**branch** | **optional.String**| specify the branch to use | 

### Return type

[**[]LocaleReport**](LocaleReport.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportShow

> ProjectReport ReportShow(ctx, projectId, optional)

Get Project Report

Get report of a single project.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
 **optional** | ***ReportShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReportShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**branch** | **optional.String**| specify the branch to use | 

### Return type

[**ProjectReport**](ProjectReport.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

