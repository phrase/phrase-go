# ProjectCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the project | 
**MainFormat** | **string** | Main file format specified by its API Extension name. Used for locale downloads if no format is specified. For API Extension names of available file formats see [Format Guide](https://support.phrase.com/hc/en-us/sections/6111343326364) or our [Formats API Endpoint](/en/api/strings/formats/list-formats). | [optional] 
**Media** | **string** | (Optional) Main technology stack used in the project. It affects for example the suggested placeholder style. Predefined values include: &#x60;Ruby&#x60;, &#x60;JavaScript&#x60;, &#x60;AngularJS&#x60;, &#x60;React&#x60;, &#x60;iOS&#x60;, &#x60;Android&#x60;, &#x60;Python&#x60;, &#x60;PHP&#x60;, &#x60;Java&#x60;, &#x60;Go&#x60;, &#x60;Windows Phone&#x60;, &#x60;Rails&#x60;, &#x60;Node.js&#x60;, &#x60;.NET&#x60;, &#x60;Django&#x60;, &#x60;Symfony&#x60;, &#x60;Yii Framework&#x60;, &#x60;Zend Framework&#x60;, &#x60;Apple App Store Description&#x60;, &#x60;Google Play Description&#x60;, but it can also take any other value. | [optional] 
**SharesTranslationMemory** | **bool** | Indicates whether the project should share the account&#39;s translation memory | [optional] 
**ProjectImage** | [***os.File**](*os.File.md) | Image to identify the project | [optional] 
**RemoveProjectImage** | **bool** | Indicates whether the project image should be deleted. | [optional] 
**AccountId** | **string** | Account ID to specify the actual account the project should be created in. Required if the requesting user is a member of multiple accounts. | [optional] 
**PointOfContact** | **string** | (Optional) User ID of the point of contact for the project. | [optional] 
**SourceProjectId** | **string** | When a source project ID is given, a clone of that project will be created, including all locales, keys and translations as well as the main project settings if they are not defined otherwise through the params. | [optional] 
**Workflow** | **string** | (Optional) Review Workflow. \&quot;simple\&quot; / \&quot;review\&quot;. [Read more](https://support.phrase.com/hc/en-us/articles/5784094755484) | [optional] 
**MachineTranslationEnabled** | **bool** | (Optional) Enable machine translation support in the project. Required for Pre-Translation | [optional] 
**EnableBranching** | **bool** | (Optional) Enable branching in the project | [optional] 
**ProtectMasterBranch** | **bool** | (Optional) Protect the master branch in project where branching is enabled | [optional] 
**EnableAllDataTypeTranslationKeysForTranslators** | **bool** | (Optional) Otherwise, translators are not allowed to edit translations other than strings | [optional] 
**EnableIcuMessageFormat** | **bool** | (Optional) We can validate and highlight your ICU messages. [Read more](https://support.phrase.com/hc/en-us/articles/5822319545116) | [optional] 
**ZeroPluralFormEnabled** | **bool** | (Optional) Displays the input fields for the &#39;ZERO&#39; plural form for every key as well although only some languages require the &#39;ZERO&#39; explicitly. | [optional] 
**AutotranslateEnabled** | **bool** | (Optional) Autopilot, requires machine_translation_enabled. [Read more](https://support.phrase.com/hc/en-us/articles/5822187934364) | [optional] 
**AutotranslateCheckNewTranslationKeys** | **bool** | (Optional) Requires autotranslate_enabled to be true | [optional] 
**AutotranslateCheckNewUploads** | **bool** | (Optional) Requires autotranslate_enabled to be true | [optional] 
**AutotranslateCheckNewLocales** | **bool** | (Optional) Requires autotranslate_enabled to be true | [optional] 
**AutotranslateMarkAsUnverified** | **bool** | (Optional) Requires autotranslate_enabled to be true | [optional] 
**AutotranslateUseMachineTranslation** | **bool** | (Optional) Requires autotranslate_enabled to be true | [optional] 
**AutotranslateUseTranslationMemory** | **bool** | (Optional) Requires autotranslate_enabled to be true | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


