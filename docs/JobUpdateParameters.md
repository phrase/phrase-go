# JobUpdateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Job name | [optional] 
**Briefing** | **string** | Briefing for the translators | [optional] 
**DueDate** | Pointer to [**NullableTime**](time.Time.md) | Date the job should be finished | [optional] 
**ReviewDueDate** | [**time.Time**](time.Time.md) | Date the review of the job should be finished. Must not be before &#x60;due_date&#x60;. | [optional] 
**TicketUrl** | **string** | URL to a ticket for this job (e.g. Jira, Trello) | [optional] 
**TargetLocaleIds** | **[]string** | List of target locales for the job. | [optional] 
**Autotranslate** | **bool** | Automatically translate the job using machine translation | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


