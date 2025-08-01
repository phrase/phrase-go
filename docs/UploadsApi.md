# phrase.UploadsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadCreate**](UploadsApi.md#UploadCreate) | **Post** /projects/{project_id}/uploads | Upload a new file
[**UploadShow**](UploadsApi.md#UploadShow) | **Get** /projects/{project_id}/uploads/{id} | Get a single upload
[**UploadsList**](UploadsApi.md#UploadsList) | **Get** /projects/{project_id}/uploads | List uploads



## UploadCreate

> Upload UploadCreate(ctx, projectId, file, fileFormat, localeId, optional)

Upload a new file

Upload a new language file. Creates necessary resources in your project.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**file** | ***os.File*****os.File**| File to be imported | 
**fileFormat** | **string**| File format. Auto-detected when possible and not specified. | 
**localeId** | **string**| Locale of the file&#39;s content. Can be the name or id of the locale. Preferred is id. | 
 **optional** | ***UploadCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UploadCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**branch** | **optional.String**| specify the branch to use | 
**tags** | **optional.String**| List of tags separated by comma to be associated with the new keys contained in the upload. | 
**updateTranslations** | **optional.Bool**| Indicates whether existing translations should be updated with the file content. | 
**updateCustomMetadata** | **optional.Bool**| Indicates whether existing custom metadata properties should be updated with the file content | [default to true]
**updateTranslationKeys** | **optional.Bool**| Pass &#x60;false&#x60; here to prevent new keys from being created and existing keys updated. | [default to true]
**updateTranslationsOnSourceMatch** | **optional.Bool**| Update target translations only if the source translations of the uploaded multilingual file match the stored translations. | [default to false]
**updateDescriptions** | **optional.Bool**| Existing key descriptions will be updated with the file content. Empty descriptions overwrite existing descriptions. | 
**convertEmoji** | **optional.Bool**| This option is obsolete. Providing the option will cause a bad request error. | 
**skipUploadTags** | **optional.Bool**| Indicates whether the upload should not create upload tags. | 
**skipUnverification** | **optional.Bool**| Indicates whether the upload should unverify updated translations. | 
**fileEncoding** | **optional.String**| Enforces a specific encoding on the file contents. Valid options are \\\&quot;UTF-8\\\&quot;, \\\&quot;UTF-16\\\&quot; and \\\&quot;ISO-8859-1\\\&quot;. | 
**localeMapping** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| Mapping between locale names and translation columns. Required in some formats like CSV or XLSX. | 
**formatOptions** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| Additional options available for specific formats. See our format guide for the [complete list](https://support.phrase.com/hc/en-us/articles/9652464547740-List-of-Supported-File-Types-Strings). | 
**autotranslate** | **optional.Bool**| If set, translations for the uploaded language will be fetched automatically. | 
**verifyMentionedTranslations** | **optional.Bool**| Indicates whether all translations mentioned in the upload should be verified. | [default to false]
**markReviewed** | **optional.Bool**| Indicated whether the imported translations should be marked as reviewed. This setting is available if the review workflow is enabled for the project. | 
**tagOnlyAffectedKeys** | **optional.Bool**| Indicates whether only keys affected (created or updated) by the upload should be tagged. The default is &#x60;false&#x60; | [default to false]
**translationKeyPrefix** | **optional.String**| This prefix will be added to all uploaded translation key names to prevent collisions. Use a meaningful prefix related to your project or file to keep key names organized. | 

### Return type

[**Upload**](Upload.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadShow

> Upload UploadShow(ctx, projectId, id, optional)

Get a single upload

View details and summary for a single upload.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
 **optional** | ***UploadShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UploadShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**branch** | **optional.String**| specify the branch to use | 

### Return type

[**Upload**](Upload.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadsList

> []Upload UploadsList(ctx, projectId, optional)

List uploads

List all uploads for the given project.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
 **optional** | ***UploadsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UploadsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**page** | **optional.Int32**| Page number | 
**perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 
**branch** | **optional.String**| specify the branch to use | 

### Return type

[**[]Upload**](Upload.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

