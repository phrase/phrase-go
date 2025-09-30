# AutomationsCreateParameters1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name of the automation | 
**Trigger** | **string** |  | 
**ProjectIds** | **[]string** | List of project IDs to associate with the automation. Currently, only the first ID in the array is used. The array format leaves room for future support of multiple projects.  | 
**JobTemplateId** | **string** | id of job template that the automation uses to create jobs from | [optional] 
**StatusFilters** | **[]string** | translation key statuses used to filter keys that are added to jobs | 
**Tags** | **[]string** | used to filter which keys are added to jobs | [optional] 
**CronSchedule** | **string** | along with time_zone, specifies when the scheduled automation is supposed to run | [optional] 
**TimeZone** | **string** | along with cron_schedule, specifies when the scheduled automation is supposed to run | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


