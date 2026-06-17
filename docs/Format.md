# Format

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable display name of the format. | 
**ApiName** | **string** | Identifier used to reference this format in API requests, such as the file_format parameter on the uploads and downloads endpoints. | 
**Description** | **string** | Human-readable summary of the format and its typical use case. | 
**Extension** | **string** | Default file extension associated with this format. | 
**DefaultEncoding** | **string** | Default character encoding used when reading or writing files in this format. | 
**Importable** | **bool** | Whether locale files can be imported using this format. | 
**Exportable** | **bool** | Whether locale files can be exported using this format. | 
**DefaultFile** | **string** | Conventional file path pattern for this format. Contains locale_name as a placeholder for the locale identifier. | 
**RendersDefaultLocale** | **bool** | When true, exported files contain the default locale&#39;s content for any key that has no translation in the target locale. | 
**IncludesLocaleInformation** | **bool** | When true, files in this format embed locale information so Phrase can detect the locale automatically on import. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


