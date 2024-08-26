# LocaleDownloadCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileFormat** | **string** | File format name. See the &lt;a href&#x3D;\&quot;https://support.phrase.com/hc/en-us/sections/6111343326364\&quot;&gt;format guide&lt;/a&gt; for all supported file formats. | 
**Branch** | **string** | specify the branch to use | [optional] 
**Tags** | **string** | Limit results to keys tagged with a list of comma separated tag names. | [optional] 
**IncludeEmptyTranslations** | **bool** | Indicates whether keys without translations should be included in the output as well. | [optional] 
**ExcludeEmptyZeroForms** | **bool** | Indicates whether zero forms should be included when empty in pluralized keys. | [optional] 
**IncludeTranslatedKeys** | **bool** | Include translated keys in the locale file. Use in combination with include_empty_translations to obtain only untranslated keys. | [optional] 
**KeepNotranslateTags** | **bool** | Indicates whether [NOTRANSLATE] tags should be kept. | [optional] 
**FormatOptions** | **map[string]interface{}** | Additional formatting and render options. See the &lt;a href&#x3D;\&quot;https://support.phrase.com/hc/en-us/sections/6111343326364\&quot;&gt;format guide&lt;/a&gt; for a list of options available for each format. Specify format options like this: &lt;code&gt;...&amp;format_options[foo]&#x3D;bar&lt;/code&gt; | [optional] 
**Encoding** | **string** | Enforces a specific encoding on the file contents. Valid options are \&quot;UTF-8\&quot;, \&quot;UTF-16\&quot; and \&quot;ISO-8859-1\&quot;. | [optional] 
**IncludeUnverifiedTranslations** | **bool** | if set to false unverified translations are excluded | [optional] 
**UseLastReviewedVersion** | **bool** | If set to true the last reviewed version of a translation is used. This is only available if the review workflow is enabled for the project. | [optional] 
**LocaleIds** | **[]string** | Locale IDs or locale names | [optional] 
**FallbackLocaleId** | **string** | If a key has no translation in the locale being downloaded the translation in the fallback locale will be used. Provide the ID of the locale that should be used as the fallback. Requires include_empty_translations to be set to &lt;code&gt;true&lt;/code&gt;. | [optional] 
**SourceLocaleId** | **string** | Provides the source language of a corresponding job as the source language of the generated locale file. This parameter will be ignored unless used in combination with a &lt;code&gt;tag&lt;/code&gt; parameter indicating a specific job. | [optional] 
**CustomMetadataFilters** | **map[string]interface{}** | Custom metadata filters. Provide the name of the metadata field and the value to filter by. Only keys with matching metadata will be included in the download.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


