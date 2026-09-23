# JobTemplateCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | **string** | specify the branch to use | [optional] 
**Name** | **string** | Job template name | 
**Briefing** | **string** | Briefing for the translators | [optional] 
**Autotranslate** | **bool** | Automatically translate the job using machine translation. | [optional] 
**SourceLocaleId** | **string** | The API id of the source language. This locale will be set as source locale for the job template. If not provided, the project default locale will be used. | [optional] 
**OwnerId** | **string** | Code of the account member to set as the job template owner. The referenced user must also be a member of the project; passing the code of an account member who is not a project member returns a 404. When omitted or blank, no owner is pre-set; the user who creates a job from this template is assigned as owner at job-creation time.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


