# JobDetails

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
**Owner** | [**UserPreview**](UserPreview.md) |  | [optional] 
**JobTagName** | **string** |  | [optional] 
**SourceTranslationsUpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**SourceLocale** | [**LocalePreview**](LocalePreview.md) |  | [optional] 
**Locales** | [**[]LocalePreview**](LocalePreview.md) |  | [optional] 
**Keys** | [**[]KeyPreview**](KeyPreview.md) |  | [optional] 
**Annotations** | [**[]JobAnnotationShort**](JobAnnotationShort.md) | Returned only when &#x60;include_annotations&#x3D;true&#x60; is supplied on the request. | [optional] 
**Locked** | **bool** | &#x60;true&#x60; if the job has been locked by the project&#39;s job-locking workflow (translations attached to the job are read-only until the job advances).  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


