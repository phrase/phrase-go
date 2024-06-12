# ReleaseTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**Branch** | **string** |  | [optional] 
**CronSchedule** | **string** | Cron schedule for the scheduler. Read more about the format of this field at https://en.wikipedia.org/wiki/Cron | [optional] 
**TimeZone** | **string** | Time zone for the scheduler | [optional] 
**NextRunAt** | [**time.Time**](time.Time.md) | The next time a release will be created for this trigger | [optional] 
**AppMinVersion** | **string** |  | [optional] 
**AppMaxVersion** | **string** |  | [optional] 
**Locales** | [**[]LocalePreview**](LocalePreview.md) |  | [optional] 
**Tags** | **[]string** |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


