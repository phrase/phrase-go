# Upload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**Filename** | **string** |  | [optional] 
**Format** | **string** |  | [optional] 
**State** | **string** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** | A user-facing message explaining why the upload failed, or &#x60;null&#x60; if the upload did not fail.  This message is intended for display only. Its wording may change at any time and it should not be parsed or relied upon programmatically.  | [optional] 
**Tag** | **string** | Unique tag of the upload  | [optional] 
**Tags** | **[]string** | List of tags that were assigned to the uploaded keys  | [optional] 
**Url** | **string** | The URL to the upload in Phrase Strings app.  | [optional] 
**User** | [**UserPreview**](.md) |  | [optional] 
**Summary** | [**UploadSummary**](UploadSummary.md) |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


