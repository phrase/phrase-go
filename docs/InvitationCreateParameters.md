# InvitationCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email of the invited user. The &lt;code&gt;email&lt;/code&gt; can not be updated once created. Create a new invitation for each unique email. | [optional] 
**Role** | **string** | Invitiation role, can be any of Manager, Developer, Translator. | [optional] 
**ProjectIds** | **string** | List of project ids the invited user has access to. | [optional] 
**LocaleIds** | **string** | List of locale ids the invited user has access to. | [optional] 
**Permissions** | [**map[string]interface{}**](.md) | Additional permissions depending on invitation role. Available permissions are &lt;code&gt;create_upload&lt;/code&gt; and &lt;code&gt;review_translations&lt;/code&gt; | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


