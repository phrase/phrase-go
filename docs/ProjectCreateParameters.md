# ProjectCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the project | [optional] 
**MainFormat** | **string** | Main file format specified by its API Extension name. Used for locale downloads if no format is specified. For API Extension names of available file formats see &lt;a href&#x3D;\&quot;https://help.phrase.com/help/supported-platforms-and-formats\&quot;&gt;Format Guide&lt;/a&gt; or our &lt;a href&#x3D;\&quot;#formats\&quot;&gt;Formats API Endpoint&lt;/a&gt;. | [optional] 
**SharesTranslationMemory** | **bool** | Indicates whether the project should share the account&#39;s translation memory | [optional] 
**ProjectImage** | [***os.File**](*os.File.md) | Image to identify the project | [optional] 
**RemoveProjectImage** | **bool** | Indicates whether the project image should be deleted. | [optional] 
**AccountId** | **string** | Account ID to specify the actual account the project should be created in. Required if the requesting user is a member of multiple accounts. | [optional] 
**SourceProjectId** | **string** | When a source project ID is given, a clone of that project will be created, including all locales, keys and translations as well as the main project settings if they are not defined otherwise through the params. | [optional] 
**MachineTranslationEnabled** | **bool** | (Optional) Enable machine translation support in the project. Required for Autopilot and Smart Suggest | [optional] 
**EnableBranching** | **bool** | (Optional) Enable branching in the project | [optional] 
**ProtectMasterBranch** | **bool** | (Optional) Protect the master branch in project where branching is enabled | [optional] 
**EnableAllDataTypeTranslationKeysForTranslators** | **bool** | (Optional) Otherwise, translators are not allowed to edit translations other than strings | [optional] 
**EnableIcuMessageFormat** | **bool** | (Optional) We can validate and highlight your ICU messages. &lt;a href&#x3D;\&quot;https://help.phrase.com/help/icu-message-format\&quot;&gt;Read more&lt;/a&gt; | [optional] 
**ZeroPluralFormEnabled** | **bool** | (Optional) Displays the input fields for the &#39;ZERO&#39; plural form for every key as well although only some languages require the &#39;ZERO&#39; explicitly. | [optional] 
**AutotranslateEnabled** | **bool** | (Optional) Autopilot, requires machine_translation_enabled. &lt;a href&#x3D;\&quot;https://help.phrase.com/help/autopilot\&quot;&gt;Read more&lt;/a&gt; | [optional] 
**AutotranslateCheckNewTranslationKeys** | **bool** | (Optional) Requires autotranslate_enabled to be true | [optional] 
**AutotranslateCheckNewUploads** | **bool** | (Optional) Requires autotranslate_enabled to be true | [optional] 
**AutotranslateCheckNewLocales** | **bool** | (Optional) Requires autotranslate_enabled to be true | [optional] 
**AutotranslateMarkAsUnverified** | **bool** | (Optional) Requires autotranslate_enabled to be true | [optional] 
**AutotranslateUseMachineTranslation** | **bool** | (Optional) Requires autotranslate_enabled to be true | [optional] 
**AutotranslateUseTranslationMemory** | **bool** | (Optional) Requires autotranslate_enabled to be true | [optional] 
**SmartSuggestEnabled** | **bool** | (Optional) Smart Suggest, requires machine_translation_enabled | [optional] 
**SmartSuggestUseGlossary** | **bool** | (Optional) Requires smart_suggest_enabled to be true | [optional] 
**SmartSuggestUseMachineTranslation** | **bool** | (Optional) Requires smart_suggest_enabled to be true | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


