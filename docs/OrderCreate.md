# OrderCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | **string** | specify the branch to use | [optional] 
**Lsp** | **string** | Name of the LSP that should process this order. Can be one of gengo, textmaster. | [optional] 
**SourceLocaleId** | **string** | Source locale for the order. Can be the name or public id of the source locale. Preferred is the public id. | [optional] 
**TargetLocaleIds** | **string** | List of target locales you want the source content translate to. Can be the name or public id of the target locales. Preferred is the public id. | [optional] 
**TranslationType** | **string** | Name of the quality level, availability depends on the LSP. Can be one of:  standard, pro (for orders processed by Gengo) and one of regular, premium, enterprise (for orders processed by TextMaster) | [optional] 
**Tag** | **string** | Tag you want to order translations for. | [optional] 
**Message** | **string** | Message that is displayed to the translators for description. | [optional] 
**StyleguideId** | **string** | Style guide for translators to be sent with the order. | [optional] 
**UnverifyTranslationsUponDelivery** | **string** | Unverify translations upon delivery. | [optional] 
**IncludeUntranslatedKeys** | **string** | Order translations for keys with untranslated content in the selected target locales. | [optional] 
**IncludeUnverifiedTranslations** | **string** | Order translations for keys with unverified content in the selected target locales. | [optional] 
**Category** | **string** | Category to use (required for orders processed by TextMaster). | [optional] 
**Quality** | **string** | Extra proofreading option to ensure consistency in vocabulary and style. Only available for orders processed by TextMaster. | [optional] 
**Priority** | **string** | Indicates whether the priority option should be ordered which decreases turnaround time by 30%. Available only for orders processed by TextMaster. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


