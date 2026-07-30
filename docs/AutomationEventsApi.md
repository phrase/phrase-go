# phrase.AutomationEventsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountAutomationEventsList**](AutomationEventsApi.md#AccountAutomationEventsList) | **Get** /accounts/{account_id}/automation_events | List automation events for an account
[**AutomationEventsList**](AutomationEventsApi.md#AutomationEventsList) | **Get** /accounts/{account_id}/automations/{automation_id}/events | List events for an automation



## AccountAutomationEventsList

> []AutomationEvent AccountAutomationEventsList(ctx, accountId, optional)

List automation events for an account

Returns the run history across all automations in the account, newest-first.  Use `automation_id` to narrow results to a single automation. Use `project_id` or `project_ids` to narrow by project.  For feature availability, see [Jobs (Strings)](https://support.phrase.com/hc/en-us/articles/5784100517788-Jobs-Strings). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
 **optional** | ***AccountAutomationEventsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AccountAutomationEventsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**page** | **optional.Int32**| Page number | 
**perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 
**automationId** | **optional.String**| Filter events to a single automation by its ID. | 
**state** | **optional.String**| Filter events by outcome state. Unrecognized values are ignored. | 
**triggeredBy** | **optional.String**| Filter events by what triggered the automation run. Unrecognized values are ignored. | 
**projectId** | **optional.String**| Filter events by project ID. Accepts a single ID or a comma-separated list of IDs. | 
**projectIds** | [**optional.Interface of []string**](string.md)| Filter events by one or more project IDs. | 
**createdAfter** | **optional.String**| Return only events created after this ISO 8601 timestamp. Returns 400 if the value is not a valid date-time. | 
**createdBefore** | **optional.String**| Return only events created before this ISO 8601 timestamp. Returns 400 if the value is not a valid date-time. | 

### Return type

[**[]AutomationEvent**](AutomationEvent.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationEventsList

> []AutomationEvent AutomationEventsList(ctx, accountId, id, optional)

List events for an automation

Returns the run history for a specific automation, newest-first.  For feature availability, see [Jobs (Strings)](https://support.phrase.com/hc/en-us/articles/5784100517788-Jobs-Strings). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**id** | **string**| ID | 
 **optional** | ***AutomationEventsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AutomationEventsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**page** | **optional.Int32**| Page number | 
**perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 
**state** | **optional.String**| Filter events by outcome state. Unrecognized values are ignored. | 
**triggeredBy** | **optional.String**| Filter events by what triggered the automation run. Unrecognized values are ignored. | 
**projectId** | **optional.String**| Filter events by project ID. Accepts a single ID or a comma-separated list of IDs. | 
**projectIds** | [**optional.Interface of []string**](string.md)| Filter events by one or more project IDs. | 
**createdAfter** | **optional.String**| Return only events created after this ISO 8601 timestamp. Returns 400 if the value is not a valid date-time. | 
**createdBefore** | **optional.String**| Return only events created before this ISO 8601 timestamp. Returns 400 if the value is not a valid date-time. | 

### Return type

[**[]AutomationEvent**](AutomationEvent.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

