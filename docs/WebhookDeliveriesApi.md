# phrase.WebhookDeliveriesApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhookDeliveriesList**](WebhookDeliveriesApi.md#WebhookDeliveriesList) | **Get** /projects/{project_id}/webhooks/{webhook_id}/deliveries | List webhook deliveries
[**WebhookDeliveriesRedeliver**](WebhookDeliveriesApi.md#WebhookDeliveriesRedeliver) | **Post** /projects/{project_id}/webhooks/{webhook_id}/deliveries/{id}/redeliver | Redeliver a single webhook delivery
[**WebhookDeliveriesShow**](WebhookDeliveriesApi.md#WebhookDeliveriesShow) | **Get** /projects/{project_id}/webhooks/{webhook_id}/deliveries/{id} | Get a single webhook delivery



## WebhookDeliveriesList

> []WebhookDelivery WebhookDeliveriesList(ctx, projectId, webhookId, optional)

List webhook deliveries

List all webhook deliveries for the given webhook_id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**webhookId** | **string**| Webhook ID | 
 **optional** | ***WebhookDeliveriesListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WebhookDeliveriesListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**responseStatusCodes** | **optional.String**| List of Response Status Codes | 
**status** | **optional.String**| Filter deliveries by outcome. &#x60;success&#x60; returns deliveries with a 2xx response code or a soft error. &#x60;failed&#x60; returns all other deliveries, including exceptions and timeouts without a response code.  | 

### Return type

[**[]WebhookDelivery**](WebhookDelivery.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDeliveriesRedeliver

> WebhookDelivery WebhookDeliveriesRedeliver(ctx, projectId, webhookId, id, optional)

Redeliver a single webhook delivery

Trigger an individual webhook delivery to be redelivered.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**webhookId** | **string**| Webhook ID | 
**id** | **string**| ID | 
 **optional** | ***WebhookDeliveriesRedeliverOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WebhookDeliveriesRedeliverOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**WebhookDelivery**](WebhookDelivery.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDeliveriesShow

> WebhookDelivery WebhookDeliveriesShow(ctx, projectId, webhookId, id, optional)

Get a single webhook delivery

Get all information about a single webhook delivery for the given ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**webhookId** | **string**| Webhook ID | 
**id** | **string**| ID | 
 **optional** | ***WebhookDeliveriesShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WebhookDeliveriesShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**WebhookDelivery**](WebhookDelivery.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

