# InvitationCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email of the invited user. The &#x60;email&#x60; can not be updated once created. Create a new invitation for each unique email. | 
**Role** | **string** | Invitiation role, can be any of Manager, Developer, Translator. | 
**ProjectIds** | **string** | List of project ids the invited user has access to. | [optional] 
**LocaleIds** | **string** | List of locale ids the invited user has access to. | [optional] 
**SpaceIds** | **[]string** | List of spaces the user is assigned to. | [optional] 
**TeamIds** | **[]string** | List of teams the user is assigned to. | [optional] 
**DefaultLocaleCodes** | **[]string** | List of default locales for the user. | [optional] 
**Permissions** | **map[string]string** | Additional permissions depending on invitation role. Available permissions are &#x60;create_upload&#x60; and &#x60;review_translations&#x60; | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


