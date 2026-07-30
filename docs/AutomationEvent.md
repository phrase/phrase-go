# AutomationEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the automation event. | [optional] 
**AutomationId** | **string** | Identifier of the automation that produced this event. | [optional] 
**State** | **string** | Outcome of the automation run. | [optional] 
**TriggeredBy** | **string** | What caused the automation to run. | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | Timestamp when the event was created. | [optional] 
**JobsCreated** | **int32** | Number of jobs created during this automation run. | [optional] 
**JobIds** | **[]string** | Identifiers of the jobs created during this automation run. | [optional] 
**Project** | [**AutomationEventProject**](AutomationEventProject.md) |  | [optional] 
**Details** | Pointer to **NullableString** | Error message describing the failure when state is &#x60;failure&#x60;; null otherwise. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


