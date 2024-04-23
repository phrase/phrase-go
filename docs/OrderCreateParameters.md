# OrderCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | **string** | specify the branch to use | [optional] 
**Name** | **string** | the name of the order, default name is: Translation order from &#39;current datetime&#39; | 
**Lsp** | **string** | Name of the LSP that should process this order. Can be one of gengo, textmaster. | 
**SourceLocaleId** | **string** | Source locale for the order. Can be the name or id of the source locale. Preferred is id. | [optional] 
**TargetLocaleIds** | **[]string** | List of target locales you want the source content translate to. Can be the name or id of the target locales. Preferred is id. | [optional] 
**TranslationType** | **string** | Name of the quality level, availability depends on the LSP. Can be one of:  standard, pro (for orders processed by Gengo) and one of regular, premium, enterprise (for orders processed by TextMaster) | [optional] 
**Tag** | **string** | Tag you want to order translations for. | [optional] 
**Message** | **string** | Message that is displayed to the translators for description. | [optional] 
**StyleguideId** | **string** | Style guide for translators to be sent with the order. | [optional] 
**UnverifyTranslationsUponDelivery** | **bool** | Unverify translations upon delivery. | [optional] 
**IncludeUntranslatedKeys** | **bool** | Order translations for keys with untranslated content in the selected target locales. | [optional] 
**IncludeUnverifiedTranslations** | **bool** | Order translations for keys with unverified content in the selected target locales. | [optional] 
**Category** | **string** | Category to use (required for orders processed by TextMaster). | [optional] 
**Quality** | **bool** | Extra proofreading option to ensure consistency in vocabulary and style. Only available for orders processed by TextMaster. | [optional] 
**Priority** | **bool** | Indicates whether the priority option should be ordered which decreases turnaround time by 30%. Available only for orders processed by TextMaster. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


