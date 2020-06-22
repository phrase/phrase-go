# \UploadsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadCreate**](UploadsApi.md#UploadCreate) | **Post** /projects/{project_id}/uploads | Upload a new file
[**UploadShow**](UploadsApi.md#UploadShow) | **Get** /projects/{project_id}/uploads/{id} | View upload details
[**UploadsList**](UploadsApi.md#UploadsList) | **Get** /projects/{project_id}/uploads | List uploads



## UploadCreate

> Upload UploadCreate(ctx, projectId, optional)

Upload a new file

Upload a new language file. Creates necessary resources in your project.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
 **optional** | ***UploadCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UploadCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **branch** | **optional.String**| specify the branch to use | 
 **file** | **optional.Interface of *os.File****optional.*os.File**| File to be imported | 
 **fileFormat** | **optional.String**| File format. Auto-detected when possible and not specified. See the [format guide](https://help.phrase.com/help/supported-platforms-and-formats) for all supported file formats. | 
 **localeId** | **optional.String**| Locale of the file&#39;s content. Can be the name or public id of the locale. Preferred is the public id. | 
 **tags** | **optional.String**| List of tags separated by comma to be associated with the new keys contained in the upload. | 
 **updateTranslations** | **optional.Bool**| Indicates whether existing translations should be updated with the file content. | 
 **updateDescriptions** | **optional.Bool**| Existing key descriptions will be updated with the file content. Empty descriptions overwrite existing descriptions. | 
 **convertEmoji** | **optional.Bool**| This option is obsolete. Providing the option will cause a bad request error. | 
 **skipUploadTags** | **optional.Bool**| Indicates whether the upload should not create upload tags. | 
 **skipUnverification** | **optional.Bool**| Indicates whether the upload should unverify updated translations. | 
 **fileEncoding** | **optional.String**| Enforces a specific encoding on the file contents. Valid options are \\\&quot;UTF-8\\\&quot;, \\\&quot;UTF-16\\\&quot; and \\\&quot;ISO-8859-1\\\&quot;. | 
 **localeMapping** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| Optional, format specific mapping between locale names and the columns the translations to those locales are contained in. | 
 **formatOptions** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| Additional options available for specific formats. See our format guide for complete list. | 
 **autotranslate** | **optional.Bool**| If set, translations for the uploaded language will be fetched automatically. | 
 **markReviewed** | **optional.Bool**| Indicated whether the imported translations should be marked as reviewed. This setting is available if the review workflow (currently beta) is enabled for the project. | 

### Return type

[**Upload**](upload.md)

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

View upload details

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

[**Upload**](upload.md)

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
 **perPage** | **optional.Int32**| allows you to specify a page size up to 100 items, 10 by default | 
 **branch** | **optional.String**| specify the branch to use | 

### Return type

[**[]Upload**](upload.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

