# WebhookCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUrl** | **string** | Callback URL to send requests to | [optional] 
**Secret** | **string** | Webhook secret used to calculate signature. If empty, the default project secret will be used. | [optional] 
**Description** | **string** | Webhook description | [optional] 
**Events** | **string** | List of event names to trigger the webhook (separated by comma) | [optional] 
**Active** | **bool** | Whether webhook is active or inactive | [optional] 
**IncludeBranches** | **bool** | If enabled, webhook will also be triggered for events from branches of the project specified. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


