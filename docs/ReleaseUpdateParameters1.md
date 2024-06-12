# ReleaseUpdateParameters1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronSchedule** | **string** | Cron schedule for the scheduler. Read more about the format of this field at https://en.wikipedia.org/wiki/Cron | [optional] 
**TimeZone** | **string** | Time zone for the scheduler | [optional] 
**LocaleIds** | **[]string** | List of locale ids that will be included in the release. | [optional] 
**Tags** | **[]string** | Only include tagged keys in the release. For a key to be included it must be tagged with all tags provided | [optional] 
**Branch** | **string** | Branch used for release | [optional] 
**AppMinVersion** | **string** | The created releases will be available only for apps with version greater or equal to this value | [optional] 
**AppMaxVersion** | **string** | The created releases will be available only for apps with version less or equal to this value | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


