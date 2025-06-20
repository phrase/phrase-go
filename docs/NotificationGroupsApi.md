# phrase.NotificationGroupsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationGroupsList**](NotificationGroupsApi.md#NotificationGroupsList) | **Get** /notification_groups | List notification groups
[**NotificationGroupsMarkAllAsRead**](NotificationGroupsApi.md#NotificationGroupsMarkAllAsRead) | **Patch** /notification_groups/mark_all_as_read | Mark all notification groups as read
[**NotificationGroupsMarkAsRead**](NotificationGroupsApi.md#NotificationGroupsMarkAsRead) | **Patch** /notification_groups/{id}/mark_as_read | Mark a notification group as read



## NotificationGroupsList

> []NotificationGroupDetail NotificationGroupsList(ctx, optional)

List notification groups

List all notification groups from the current user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NotificationGroupsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NotificationGroupsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**page** | **optional.Int32**| Page number | 
**perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 

### Return type

[**[]NotificationGroupDetail**](NotificationGroupDetail.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationGroupsMarkAllAsRead

> []NotificationGroupDetail NotificationGroupsMarkAllAsRead(ctx, optional)

Mark all notification groups as read

Mark all notification groups of the current user as read

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NotificationGroupsMarkAllAsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NotificationGroupsMarkAllAsReadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**[]NotificationGroupDetail**](NotificationGroupDetail.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationGroupsMarkAsRead

> NotificationGroupDetail NotificationGroupsMarkAsRead(ctx, id, optional)

Mark a notification group as read

Mark a notifications group of the current user as read

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| ID | 
 **optional** | ***NotificationGroupsMarkAsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NotificationGroupsMarkAsReadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**NotificationGroupDetail**](NotificationGroupDetail.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

