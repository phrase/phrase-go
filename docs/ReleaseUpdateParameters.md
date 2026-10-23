# ReleaseUpdateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Description of the release. Send an empty string or &#x60;null&#x60; to clear the value; omit the field to leave it unchanged. The Phrase CLI drops &#x60;null&#x60; values before sending, so use an empty string there. | [optional] 
**Platforms** | **[]string** | List of platforms the release should support. | [optional] 
**AppMinVersion** | Pointer to **NullableString** | Minimum version of the app that the release supports in semver format. Send an empty string or &#x60;null&#x60; to clear the value; omit the field to leave it unchanged. The Phrase CLI drops &#x60;null&#x60; values before sending, so use an empty string there. | [optional] 
**AppMaxVersion** | Pointer to **NullableString** | Maximum version of the app that the release supports in semver format. Send an empty string or &#x60;null&#x60; to clear the value; omit the field to leave it unchanged. The Phrase CLI drops &#x60;null&#x60; values before sending, so use an empty string there. | [optional] 
**Branch** | **string** | Branch used for release | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


