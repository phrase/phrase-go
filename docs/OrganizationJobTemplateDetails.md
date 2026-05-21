# OrganizationJobTemplateDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**Name** | **string** |  | [optional] 
**Briefing** | **string** |  | [optional] 
**AutotranslateEnabled** | **bool** | When &#x60;true&#x60;, jobs created from this template are auto-translated on creation. Maps to the &#x60;autotranslate&#x60; field on the request body.  | [optional] 
**SourceLocaleId** | Pointer to **NullableString** | Optional. ID of the source locale used by jobs created from this template. When omitted, the project&#39;s default source locale is used.  | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**Owner** | [**UserPreview**](UserPreview.md) |  | [optional] 
**Creator** | [**UserPreview**](UserPreview.md) |  | [optional] 
**Locales** | [**[]LocalePreview**](LocalePreview.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


