# ReleaseCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the release | [optional] 
**Platforms** | **[]string** | List of platforms the release should support. | [optional] 
**LocaleIds** | **[]string** | List of locale ids that will be included in the release. If empty, distribution locales will be used | [optional] 
**Tags** | **[]string** | Only include tagged keys in the release. For a key to be included it must be tagged with all tags provided | [optional] 
**AppMinVersion** | **string** | Minimum version of the app that the release supports in semver format | [optional] 
**AppMaxVersion** | **string** | Maximum version of the app that the release supports in semver format | [optional] 
**Branch** | **string** | Branch used for release | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


