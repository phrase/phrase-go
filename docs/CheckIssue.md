# CheckIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**CheckName** | **string** | Identifier of the check that reported this issue. One of: &#x60;translation_content_length&#x60;, &#x60;translation_placeholder_usage&#x60;, &#x60;translation_glossary_usage&#x60;. | [optional] 
**State** | **string** | Current state of the check issue. One of: &#x60;active&#x60;, &#x60;solved&#x60;, &#x60;dismissed&#x60;. | [optional] 
**Description** | **string** | Human-readable description of the reported issue, always in English. This message is intended for display only. Its wording may change at any time and it should not be parsed or relied upon programmatically. | [optional] 
**DismissedAt** | Pointer to [**NullableTime**](time.Time.md) |  | [optional] 
**SolvedAt** | Pointer to [**NullableTime**](time.Time.md) |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**Translation** | [**Translation**](Translation.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


