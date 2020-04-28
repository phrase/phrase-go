# LocaleDownloadParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | **string** | specify the branch to use | [optional] 
**FileFormat** | **string** | File format name. See the format guide for all supported file formats. | [optional] 
**Tags** | **string** | Limit results to keys tagged with a list of comma separated tag names. | [optional] 
**Tag** | **string** | Limit download to tagged keys. This parameter is deprecated. Please use the \&quot;tags\&quot; parameter instead | [optional] 
**IncludeEmptyTranslations** | **bool** | Indicates whether keys without translations should be included in the output as well. | [optional] 
**IncludeTranslatedKeys** | **bool** | Include translated keys in the locale file. Use in combination with include_empty_translations to obtain only untranslated keys. | [optional] 
**KeepNotranslateTags** | **bool** | Indicates whether [NOTRANSLATE] tags should be kept. | [optional] 
**ConvertEmoji** | **bool** | This option is obsolete. Projects that were created on or after Nov 29th 2019 or that did not contain emoji by then will not require this flag any longer since emoji are now supported natively. | [optional] 
**FormatOptions** | [**map[string]interface{}**](.md) | Additional formatting and render options. See the &lt;a href&#x3D;\&quot;https://help.phrase.com/help/supported-platforms-and-formats\&quot;&gt;format guide&lt;/a&gt; for a list of options available for each format. Specify format options like this: &lt;code&gt;...&amp;format_options[foo]&#x3D;bar&lt;/code&gt; | [optional] 
**Encoding** | **string** | Enforces a specific encoding on the file contents. Valid options are \&quot;UTF-8\&quot;, \&quot;UTF-16\&quot; and \&quot;ISO-8859-1\&quot;. | [optional] 
**SkipUnverifiedTranslations** | **bool** | Indicates whether the locale file should skip all unverified translations. This parameter is deprecated and should be replaced with &lt;code&gt;include_unverified_translations&lt;/code&gt;. | [optional] 
**IncludeUnverifiedTranslations** | **bool** | if set to false unverified translations are excluded | [optional] 
**UseLastReviewedVersion** | **bool** | If set to true the last reviewed version of a translation is used. This is only available if the review workflow (currently in beta) is enabled for the project. | [optional] 
**FallbackLocaleId** | **string** | If a key has no translation in the locale being downloaded the translation in the fallback locale will be used. Provide the public ID of the locale that should be used as the fallback. Requires include_empty_translations to be set to &lt;code&gt;true&lt;/code&gt;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


