# phrase.AuthorizationsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizationCreate**](AuthorizationsApi.md#AuthorizationCreate) | **Post** /authorizations | Create an authorization
[**AuthorizationDelete**](AuthorizationsApi.md#AuthorizationDelete) | **Delete** /authorizations/{id} | Delete an authorization
[**AuthorizationShow**](AuthorizationsApi.md#AuthorizationShow) | **Get** /authorizations/{id} | Get a single authorization
[**AuthorizationUpdate**](AuthorizationsApi.md#AuthorizationUpdate) | **Patch** /authorizations/{id} | Update an authorization
[**AuthorizationsList**](AuthorizationsApi.md#AuthorizationsList) | **Get** /authorizations | List authorizations



## AuthorizationCreate

> AuthorizationWithToken AuthorizationCreate(ctx, authorizationCreateParameters, optional)

Create an authorization

Create a new authorization.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorizationCreateParameters** | [**AuthorizationCreateParameters**](AuthorizationCreateParameters.md)|  | 
 **optional** | ***AuthorizationCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthorizationCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**AuthorizationWithToken**](AuthorizationWithToken.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizationDelete

> AuthorizationDelete(ctx, id, optional)

Delete an authorization

Delete an existing authorization. API calls using that token will stop working.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| ID | 
 **optional** | ***AuthorizationDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthorizationDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizationShow

> Authorization AuthorizationShow(ctx, id, optional)

Get a single authorization

Get details on a single authorization.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| ID | 
 **optional** | ***AuthorizationShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthorizationShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**Authorization**](Authorization.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizationUpdate

> Authorization AuthorizationUpdate(ctx, id, authorizationUpdateParameters, optional)

Update an authorization

Update an existing authorization.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| ID | 
**authorizationUpdateParameters** | [**AuthorizationUpdateParameters**](AuthorizationUpdateParameters.md)|  | 
 **optional** | ***AuthorizationUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthorizationUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**Authorization**](Authorization.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizationsList

> []Authorization AuthorizationsList(ctx, optional)

List authorizations

List all your authorizations.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthorizationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthorizationsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**page** | **optional.Int32**| Page number | 
**perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 

### Return type

[**[]Authorization**](Authorization.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

