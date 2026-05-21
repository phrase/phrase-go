# Automation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**Name** | **string** |  | [optional] 
**Status** | **string** |  | [optional] 
**Trigger** | **string** |  | [optional] 
**StatusFilters** | **[]string** | translation key statuses used to filter keys that are added to jobs | [optional] 
**ProjectId** | **string** |  | [optional] 
**ProjectIds** | **[]string** | All project IDs the automation applies to. Returned alongside the singular &#x60;project_id&#x60; for backwards compatibility. | [optional] 
**JobTemplateId** | **string** |  | [optional] 
**JobOwnerId** | Pointer to **NullableString** | User ID of the job owner that newly created jobs are assigned to. | [optional] 
**IncludeOnlyUpdatedLocales** | **bool** | When &#x60;true&#x60;, the automation only acts on locales that changed since its last run. | [optional] 
**Tags** | **[]string** |  | [optional] 
**CronSchedule** | **string** |  | [optional] 
**TimeZone** | **string** |  | [optional] 
**Account** | [**Account**](Account.md) |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


