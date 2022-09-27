# \LocalesApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountLocales**](LocalesApi.md#AccountLocales) | **Get** /accounts/{account_id}/locales | List locales used in account
[**LocaleCreate**](LocalesApi.md#LocaleCreate) | **Post** /projects/{project_id}/locales | Create a locale
[**LocaleDelete**](LocalesApi.md#LocaleDelete) | **Delete** /projects/{project_id}/locales/{id} | Delete a locale
[**LocaleDownload**](LocalesApi.md#LocaleDownload) | **Get** /projects/{project_id}/locales/{id}/download | Download a locale
[**LocaleShow**](LocalesApi.md#LocaleShow) | **Get** /projects/{project_id}/locales/{id} | Get a single locale
[**LocaleUpdate**](LocalesApi.md#LocaleUpdate) | **Patch** /projects/{project_id}/locales/{id} | Update a locale
[**LocalesList**](LocalesApi.md#LocalesList) | **Get** /projects/{project_id}/locales | List locales



## AccountLocales

> []LocalePreview1 AccountLocales(ctx, id, optional)

List locales used in account

List all locales unique by locale code used across all projects within an account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| ID | 
 **optional** | ***AccountLocalesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AccountLocalesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **page** | **optional.Int32**| Page number | 
 **perPage** | **optional.Int32**| allows you to specify a page size up to 100 items, 25 by default | 

### Return type

[**[]LocalePreview1**](locale_preview_1.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocaleCreate

> LocaleDetails LocaleCreate(ctx, projectId, localeCreateParameters, optional)

Create a locale

Create a new locale.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**localeCreateParameters** | [**LocaleCreateParameters**](LocaleCreateParameters.md)|  | 
 **optional** | ***LocaleCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LocaleCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**LocaleDetails**](locale_details.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocaleDelete

> LocaleDelete(ctx, projectId, id, optional)

Delete a locale

Delete an existing locale.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
 **optional** | ***LocaleDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LocaleDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **branch** | **optional.String**| specify the branch to use | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocaleDownload

> *os.File LocaleDownload(ctx, projectId, id, optional)

Download a locale

Download a locale in a specific file format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
 **optional** | ***LocaleDownloadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LocaleDownloadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **branch** | **optional.String**| specify the branch to use | 
 **fileFormat** | **optional.String**| File format name. See the &lt;a href&#x3D;\&quot;https://support.phrase.com/hc/en-us/articles/5784070560412\&quot;&gt;format guide&lt;/a&gt; for all supported file formats. | 
 **tags** | **optional.String**| Limit results to keys tagged with a list of comma separated tag names. | 
 **tag** | **optional.String**| Limit download to tagged keys. This parameter is deprecated. Please use the \&quot;tags\&quot; parameter instead | 
 **includeEmptyTranslations** | **optional.Bool**| Indicates whether keys without translations should be included in the output as well. | 
 **excludeEmptyZeroForms** | **optional.Bool**| Indicates whether zero forms should be included when empty in pluralized keys. | 
 **includeTranslatedKeys** | **optional.Bool**| Include translated keys in the locale file. Use in combination with include_empty_translations to obtain only untranslated keys. | 
 **keepNotranslateTags** | **optional.Bool**| Indicates whether [NOTRANSLATE] tags should be kept. | 
 **convertEmoji** | **optional.Bool**| This option is obsolete. Projects that were created on or after Nov 29th 2019 or that did not contain emoji by then will not require this flag any longer since emoji are now supported natively. | 
 **formatOptions** | [**optional.Interface of map[string]interface{}**](.md)| Additional formatting and render options. See the &lt;a href&#x3D;\&quot;https://support.phrase.com/hc/en-us/articles/5784070560412\&quot;&gt;format guide&lt;/a&gt; for a list of options available for each format. Specify format options like this: &lt;code&gt;...&amp;format_options[foo]&#x3D;bar&lt;/code&gt; | 
 **encoding** | **optional.String**| Enforces a specific encoding on the file contents. Valid options are \&quot;UTF-8\&quot;, \&quot;UTF-16\&quot; and \&quot;ISO-8859-1\&quot;. | 
 **skipUnverifiedTranslations** | **optional.Bool**| Indicates whether the locale file should skip all unverified translations. This parameter is deprecated and should be replaced with &lt;code&gt;include_unverified_translations&lt;/code&gt;. | 
 **includeUnverifiedTranslations** | **optional.Bool**| if set to false unverified translations are excluded | 
 **useLastReviewedVersion** | **optional.Bool**| If set to true the last reviewed version of a translation is used. This is only available if the review workflow (currently in beta) is enabled for the project. | 
 **fallbackLocaleId** | **optional.String**| If a key has no translation in the locale being downloaded the translation in the fallback locale will be used. Provide the public ID of the locale that should be used as the fallback. Requires include_empty_translations to be set to &lt;code&gt;true&lt;/code&gt;. | 
 **sourceLocaleId** | **optional.String**| Provides the source language of a corresponding job as the source language of the generated locale file. This parameter will be ignored unless used in combination with a &lt;code&gt;tag&lt;/code&gt; parameter indicating a specific job. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: *

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocaleShow

> LocaleDetails LocaleShow(ctx, projectId, id, optional)

Get a single locale

Get details on a single locale for a given project.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
 **optional** | ***LocaleShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LocaleShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **branch** | **optional.String**| specify the branch to use | 

### Return type

[**LocaleDetails**](locale_details.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocaleUpdate

> LocaleDetails LocaleUpdate(ctx, projectId, id, localeUpdateParameters, optional)

Update a locale

Update an existing locale.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
**localeUpdateParameters** | [**LocaleUpdateParameters**](LocaleUpdateParameters.md)|  | 
 **optional** | ***LocaleUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LocaleUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**LocaleDetails**](locale_details.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocalesList

> []Locale LocalesList(ctx, projectId, optional)

List locales

List all locales for the given project.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
 **optional** | ***LocalesListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LocalesListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **page** | **optional.Int32**| Page number | 
 **perPage** | **optional.Int32**| allows you to specify a page size up to 100 items, 25 by default | 
 **sortBy** | **optional.String**| Sort locales. Valid options are \&quot;name_asc\&quot;, \&quot;name_desc\&quot;, \&quot;default_asc\&quot;, \&quot;default_desc\&quot;. | 
 **branch** | **optional.String**| specify the branch to use | 

### Return type

[**[]Locale**](locale.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

