# \ReleaseTriggersApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReleaseTriggersCreate**](ReleaseTriggersApi.md#ReleaseTriggersCreate) | **Post** /accounts/{account_id}/distributions/{distribution_id}/release_triggers | Create a release trigger
[**ReleaseTriggersDestroy**](ReleaseTriggersApi.md#ReleaseTriggersDestroy) | **Delete** /accounts/{account_id}/distributions/{distribution_id}/release_triggers/{id} | Delete a single release trigger
[**ReleaseTriggersList**](ReleaseTriggersApi.md#ReleaseTriggersList) | **Get** /accounts/{account_id}/distributions/{distribution_id}/release_triggers | List release triggers
[**ReleaseTriggersShow**](ReleaseTriggersApi.md#ReleaseTriggersShow) | **Get** /accounts/{account_id}/distributions/{distribution_id}/release_triggers/{id} | Get a single release trigger
[**ReleaseTriggersUpdate**](ReleaseTriggersApi.md#ReleaseTriggersUpdate) | **Patch** /accounts/{account_id}/distributions/{distribution_id}/release_triggers/{id} | Update a release trigger



## ReleaseTriggersCreate

> ReleaseTrigger ReleaseTriggersCreate(ctx, accountId, distributionId, releaseCreateParameters1, optional)

Create a release trigger

Create a new recurring release. New releases will be published automatically, based on the cron schedule provided. Currently, only one release trigger can exist per distribution.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**distributionId** | **string**| Distribution ID | 
**releaseCreateParameters1** | [**ReleaseCreateParameters1**](ReleaseCreateParameters1.md)|  | 
 **optional** | ***ReleaseTriggersCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReleaseTriggersCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**ReleaseTrigger**](ReleaseTrigger.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleaseTriggersDestroy

> ReleaseTriggersDestroy(ctx, accountId, distributionId, id, optional)

Delete a single release trigger

Delete a single release trigger.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**distributionId** | **string**| Distribution ID | 
**id** | **string**| ID | 
 **optional** | ***ReleaseTriggersDestroyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReleaseTriggersDestroyOpts struct


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


## ReleaseTriggersList

> []ReleaseTrigger ReleaseTriggersList(ctx, accountId, distributionId, optional)

List release triggers

List all release triggers for the given distribution.  Note: Currently only one release trigger can exist per distribution. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**distributionId** | **string**| Distribution ID | 
 **optional** | ***ReleaseTriggersListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReleaseTriggersListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**[]ReleaseTrigger**](ReleaseTrigger.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleaseTriggersShow

> ReleaseTrigger ReleaseTriggersShow(ctx, accountId, distributionId, id, optional)

Get a single release trigger

Get details of a single release trigger.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**distributionId** | **string**| Distribution ID | 
**id** | **string**| ID | 
 **optional** | ***ReleaseTriggersShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReleaseTriggersShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**ReleaseTrigger**](ReleaseTrigger.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleaseTriggersUpdate

> ReleaseTrigger ReleaseTriggersUpdate(ctx, accountId, distributionId, id, releaseUpdateParameters1, optional)

Update a release trigger

Update a recurring release.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**distributionId** | **string**| Distribution ID | 
**id** | **string**| ID | 
**releaseUpdateParameters1** | [**ReleaseUpdateParameters1**](ReleaseUpdateParameters1.md)|  | 
 **optional** | ***ReleaseTriggersUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReleaseTriggersUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**ReleaseTrigger**](ReleaseTrigger.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

