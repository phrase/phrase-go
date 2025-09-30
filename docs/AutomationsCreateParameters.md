# AutomationsCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name of the automation | 
**Trigger** | **string** |  | 
**ProjectIds** | **[]string** | List of project IDs to associate with the automation. Currently, only the first ID in the array is used. The array format leaves room for future support of multiple projects.  | 
**JobTemplateId** | **string** | id of job template that the automation uses to create jobs from | [optional] 
**StatusFilters** | **[]string** | Translation states used when selecting keys for a job.  States are derived from associated translations, not the keys themselves.  When review workflow is enabled, &#x60;ready_for_review&#x60; is internally treated as &#x60;translated&#x60;.  | 
**Tags** | **[]string** | used to filter which keys are added to jobs | [optional] 
**CronSchedule** | **string** | along with time_zone, specifies when the scheduled automation is supposed to run | [optional] 
**TimeZone** | **string** | along with cron_schedule, specifies when the scheduled automation is supposed to run | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


