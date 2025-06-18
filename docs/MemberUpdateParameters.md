# MemberUpdateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Strategy** | **string** | Update strategy, can be any of set, add, remove. If provided, it will set, add or remove given spaces, projects and locale ids from users access list. | [optional] 
**Role** | **string** | Member role, can be any of of Admin, ProjectManager, Developer, Designer, Translator | [optional] 
**ProjectIds** | **string** | List of project ids the user has access to.  | [optional] 
**LocaleIds** | **string** | List of locale ids the user has access to. | [optional] 
**DefaultLocaleCodes** | **[]string** | List of default locales for the user. | [optional] 
**SpaceIds** | **[]string** | List of spaces the user is assigned to. | [optional] 
**Permissions** | **map[string]string** | Additional permissions depending on member role. Available permissions are &#x60;create_upload&#x60; and &#x60;review_translations&#x60; | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


