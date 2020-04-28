# DistributionCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the distribution | [optional] 
**ProjectId** | **string** | Project id the distribution should be assigned to. | [optional] 
**Platforms** | **[]string** | List of platforms the distribution should support. | [optional] 
**FormatOptions** | [**map[string]interface{}**](.md) | Additional formatting and render options. Only &lt;code&gt;enclose_in_cdata&lt;/code&gt; is available for platform &lt;code&gt;android&lt;/code&gt;. | [optional] 
**FallbackToNonRegionalLocale** | **bool** | Indicates whether to fallback to non regional locale when locale can not be found | [optional] 
**FallbackToDefaultLocale** | **bool** | Indicates whether to fallback to projects default locale when locale can not be found | [optional] 
**UseLastReviewedVersion** | **bool** | Use last reviewed instead of latest translation in a project | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


