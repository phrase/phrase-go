# JobCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | **string** | specify the branch to use | [optional] 
**Name** | **string** | Job name | 
**SourceLocaleId** | **string** | The API id of the source language | [optional] 
**Briefing** | **string** | Briefing for the translators | [optional] 
**DueDate** | Pointer to [**NullableTime**](time.Time.md) | Date the job should be finished | [optional] 
**TicketUrl** | **string** | URL to a ticket for this job (e.g. Jira, Trello) | [optional] 
**Tags** | **[]string** | tags of keys that should be included within the job.  *Note: a tag matches every key currently carrying that tag, not just the ones you just tagged. For example, if hundreds of pre-existing keys already share the tag &#x60;myUploadTag&#x60;, adding it here pulls in every one of them, not only the key you just tagged. Use &#x60;translation_key_ids&#x60; to scope the job to specific keys instead.*  | [optional] 
**TranslationKeyIds** | **[]string** | ids of keys that should be included within the job | [optional] 
**TargetLocaleIds** | **[]string** | List of target locales for the job. Mutually exclusive with &#x60;job_template_id&#x60;. | [optional] 
**JobTemplateId** | **string** | id of a job template you would like to model the created job after. Any manually added parameters will take preference over template attributes. Mutually exclusive with &#x60;target_locale_ids&#x60;. | [optional] 
**Autotranslate** | **bool** | Automatically translate the job using machine translation. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


