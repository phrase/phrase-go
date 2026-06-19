# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**Name** | **string** |  | [optional] 
**Briefing** | **string** |  | [optional] 
**DueDate** | Pointer to [**NullableTime**](time.Time.md) |  | [optional] 
**State** | **string** |  | [optional] 
**TicketUrl** | **string** |  | [optional] 
**Project** | [**ProjectShort**](ProjectShort.md) |  | [optional] 
**Branch** | [**BranchName**](BranchName.md) |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**AutomationId** | Pointer to **NullableString** | The ID of the automation that created this job, or null if the job was created manually. | [optional] 
**JobTemplateId** | Pointer to **NullableString** | The ID of the job template this job was created from, or null if no template was used. | [optional] 
**ReviewDueDate** | Pointer to [**NullableTime**](time.Time.md) | The review due date for this job. Returns &#x60;null&#x60; when the project does not have review workflow enabled. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


