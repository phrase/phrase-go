# JobCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | **string** | specify the branch to use | [optional] 
**Name** | **string** | Job name | [optional] 
**SourceLocaleId** | **string** | The API id of the source language | [optional] 
**Briefing** | **string** | Briefing for the translators | [optional] 
**DueDate** | Pointer to [**NullableTime**](time.Time.md) | Date the job should be finished | [optional] 
**TicketUrl** | **string** | URL to a ticket for this job (e.g. Jira, Trello) | [optional] 
**Tags** | **[]string** | tags of keys that should be included within the job | [optional] 
**TranslationKeyIds** | **[]string** | ids of keys that should be included within the job | [optional] 
**JobTemplateId** | **string** | id of a job template you would like to model the created job after. Any manually added parameters will take preference over template attributes. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


