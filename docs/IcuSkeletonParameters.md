# IcuSkeletonParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | Source content to derive skeletons from. Mutually exclusive with &#x60;id&#x60;; exactly one of the two must be provided.  | [optional] 
**Id** | **string** | Translation code to source content from. Mutually exclusive with &#x60;content&#x60;; exactly one of the two must be provided.  | [optional] 
**LocaleCodes** | **[]string** | Locale codes | [optional] 
**KeepContent** | **bool** | Keep the content and add missing plural forms for each locale | [optional] 
**ZeroFormEnabled** | **bool** | Indicates whether the zero form should be included or excluded in the returned skeletons | [optional] 
**CldrVersion** | **string** | Strings supports two CLDR variants, when it comes to pluralization rules. \\ You can choose which one you want to use when constructing the skeletons. Possible values \\ are &#x60;legacy&#x60; and &#x60;cldr_41&#x60;. Default value is &#x60;legacy&#x60;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


