# \UsersApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShowUser**](UsersApi.md#ShowUser) | **Get** /user | Show current User



## ShowUser

> CurrentUser ShowUser(ctx, optional)

Show current User

Show details for current User.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ShowUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ShowUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**CurrentUser**](current_user.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

