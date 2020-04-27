# \TranslationsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TranslationCreate**](TranslationsApi.md#TranslationCreate) | **Post** /projects/{project_id}/translations | Create a translation
[**TranslationExclude**](TranslationsApi.md#TranslationExclude) | **Patch** /projects/{project_id}/translations/{id}/exclude | Exclude a translation from export
[**TranslationInclude**](TranslationsApi.md#TranslationInclude) | **Patch** /projects/{project_id}/translations/{id}/include | Revoke exclusion of a translation in export
[**TranslationReview**](TranslationsApi.md#TranslationReview) | **Patch** /projects/{project_id}/translations/{id}/review | Review a translation
[**TranslationShow**](TranslationsApi.md#TranslationShow) | **Get** /projects/{project_id}/translations/{id} | Get a single translation
[**TranslationUnverify**](TranslationsApi.md#TranslationUnverify) | **Patch** /projects/{project_id}/translations/{id}/unverify | Mark a translation as unverified
[**TranslationUpdate**](TranslationsApi.md#TranslationUpdate) | **Patch** /projects/{project_id}/translations/{id} | Update a translation
[**TranslationVerify**](TranslationsApi.md#TranslationVerify) | **Patch** /projects/{project_id}/translations/{id}/verify | Verify a translation
[**TranslationsByKey**](TranslationsApi.md#TranslationsByKey) | **Get** /projects/{project_id}/keys/{key_id}/translations | List translations by key
[**TranslationsByLocale**](TranslationsApi.md#TranslationsByLocale) | **Get** /projects/{project_id}/locales/{locale_id}/translations | List translations by locale
[**TranslationsExclude**](TranslationsApi.md#TranslationsExclude) | **Patch** /projects/{project_id}/translations/exclude | Set exclude from export flag on translations selected by query
[**TranslationsInclude**](TranslationsApi.md#TranslationsInclude) | **Patch** /projects/{project_id}/translations/include | Remove exlude from import flag from translations selected by query
[**TranslationsList**](TranslationsApi.md#TranslationsList) | **Get** /projects/{project_id}/translations | List all translations
[**TranslationsReview**](TranslationsApi.md#TranslationsReview) | **Patch** /projects/{project_id}/translations/review | Review translations selected by query
[**TranslationsSearch**](TranslationsApi.md#TranslationsSearch) | **Post** /projects/{project_id}/translations/search | Search translations
[**TranslationsUnverify**](TranslationsApi.md#TranslationsUnverify) | **Patch** /projects/{project_id}/translations/unverify | Mark translations selected by query as unverified
[**TranslationsVerify**](TranslationsApi.md#TranslationsVerify) | **Patch** /projects/{project_id}/translations/verify | Verify translations selected by query



## TranslationCreate

> TranslationCreate(ctx, projectId, translationCreate, optional)

Create a translation

Create a translation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**translationCreate** | [**TranslationCreate**](TranslationCreate.md)|  | 
 **optional** | ***TranslationCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationExclude

> map[string]interface{} TranslationExclude(ctx, projectId, id, translationExclude, optional)

Exclude a translation from export

Set exclude from export flag on an existing translation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
**translationExclude** | [**TranslationExclude**](TranslationExclude.md)|  | 
 **optional** | ***TranslationExcludeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationExcludeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**map[string]interface{}**](object.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationInclude

> map[string]interface{} TranslationInclude(ctx, projectId, id, translationInclude, optional)

Revoke exclusion of a translation in export

Remove exclude from export flag from an existing translation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
**translationInclude** | [**TranslationInclude**](TranslationInclude.md)|  | 
 **optional** | ***TranslationIncludeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationIncludeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**map[string]interface{}**](object.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationReview

> map[string]interface{} TranslationReview(ctx, projectId, id, translationReview, optional)

Review a translation

Mark an existing translation as reviewed.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
**translationReview** | [**TranslationReview**](TranslationReview.md)|  | 
 **optional** | ***TranslationReviewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationReviewOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**map[string]interface{}**](object.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationShow

> map[string]interface{} TranslationShow(ctx, projectId, id, translationShow, optional)

Get a single translation

Get details on a single translation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
**translationShow** | [**TranslationShow**](TranslationShow.md)|  | 
 **optional** | ***TranslationShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

**map[string]interface{}**

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationUnverify

> map[string]interface{} TranslationUnverify(ctx, projectId, id, translationUnverify, optional)

Mark a translation as unverified

Mark an existing translation as unverified.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
**translationUnverify** | [**TranslationUnverify**](TranslationUnverify.md)|  | 
 **optional** | ***TranslationUnverifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationUnverifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**map[string]interface{}**](object.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationUpdate

> map[string]interface{} TranslationUpdate(ctx, projectId, id, translationUpdate, optional)

Update a translation

Update an existing translation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
**translationUpdate** | [**TranslationUpdate**](TranslationUpdate.md)|  | 
 **optional** | ***TranslationUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**map[string]interface{}**](object.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationVerify

> map[string]interface{} TranslationVerify(ctx, projectId, id, translationVerify, optional)

Verify a translation

Verify an existing translation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
**translationVerify** | [**TranslationVerify**](TranslationVerify.md)|  | 
 **optional** | ***TranslationVerifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationVerifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**map[string]interface{}**](object.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationsByKey

> []map[string]interface{} TranslationsByKey(ctx, projectId, keyId, translationsByKey, optional)

List translations by key

List translations for a specific key.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keyId** | **string**| Translation Key ID | 
**translationsByKey** | [**TranslationsByKey**](TranslationsByKey.md)|  | 
 **optional** | ***TranslationsByKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationsByKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **page** | **optional.Int32**| Page number | 
 **perPage** | **optional.Int32**| allows you to specify a page size up to 100 items, 10 by default | 

### Return type

[**[]map[string]interface{}**](object.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationsByLocale

> []map[string]interface{} TranslationsByLocale(ctx, projectId, localeId, translationsByLocale, optional)

List translations by locale

List translations for a specific locale. If you want to download all translations for one locale we recommend to use the <code>locales#download</code> endpoint.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**localeId** | **string**| Locale ID | 
**translationsByLocale** | [**TranslationsByLocale**](TranslationsByLocale.md)|  | 
 **optional** | ***TranslationsByLocaleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationsByLocaleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **page** | **optional.Int32**| Page number | 
 **perPage** | **optional.Int32**| allows you to specify a page size up to 100 items, 10 by default | 

### Return type

[**[]map[string]interface{}**](object.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationsExclude

> map[string]interface{} TranslationsExclude(ctx, projectId, translationsExclude, optional)

Set exclude from export flag on translations selected by query

Exclude translations matching query from locale export.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**translationsExclude** | [**TranslationsExclude**](TranslationsExclude.md)|  | 
 **optional** | ***TranslationsExcludeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationsExcludeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**map[string]interface{}**](object.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationsInclude

> map[string]interface{} TranslationsInclude(ctx, projectId, translationsInclude, optional)

Remove exlude from import flag from translations selected by query

Include translations matching query in locale export.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**translationsInclude** | [**TranslationsInclude**](TranslationsInclude.md)|  | 
 **optional** | ***TranslationsIncludeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationsIncludeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**map[string]interface{}**](object.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationsList

> []Translation TranslationsList(ctx, projectId, translationsList, optional)

List all translations

List translations for the given project. If you want to download all translations for one locale we recommend to use the <code>locales#download</code> endpoint.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**translationsList** | [**TranslationsList**](TranslationsList.md)|  | 
 **optional** | ***TranslationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **page** | **optional.Int32**| Page number | 
 **perPage** | **optional.Int32**| allows you to specify a page size up to 100 items, 10 by default | 

### Return type

[**[]Translation**](translation.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationsReview

> map[string]interface{} TranslationsReview(ctx, projectId, translationsReview, optional)

Review translations selected by query

Review translations matching query.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**translationsReview** | [**TranslationsReview**](TranslationsReview.md)|  | 
 **optional** | ***TranslationsReviewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationsReviewOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**map[string]interface{}**](object.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationsSearch

> []map[string]interface{} TranslationsSearch(ctx, projectId, translationsSearch, optional)

Search translations

Search translations for the given project. Provides the same search interface as <code>translations#index</code> but allows POST requests to avoid limitations imposed by GET requests. If you want to download all translations for one locale we recommend to use the <code>locales#download</code> endpoint.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**translationsSearch** | [**TranslationsSearch**](TranslationsSearch.md)|  | 
 **optional** | ***TranslationsSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationsSearchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **page** | **optional.Int32**| Page number | 
 **perPage** | **optional.Int32**| allows you to specify a page size up to 100 items, 10 by default | 

### Return type

[**[]map[string]interface{}**](object.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationsUnverify

> map[string]interface{} TranslationsUnverify(ctx, projectId, translationsUnverify, optional)

Mark translations selected by query as unverified

Mark translations matching query as unverified.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**translationsUnverify** | [**TranslationsUnverify**](TranslationsUnverify.md)|  | 
 **optional** | ***TranslationsUnverifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationsUnverifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**map[string]interface{}**](object.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationsVerify

> AffectedCount TranslationsVerify(ctx, projectId, translationsVerify, optional)

Verify translations selected by query

Verify translations matching query.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**translationsVerify** | [**TranslationsVerify**](TranslationsVerify.md)|  | 
 **optional** | ***TranslationsVerifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationsVerifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**AffectedCount**](affected_count.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

