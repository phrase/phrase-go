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
[**TranslationsExcludeCollection**](TranslationsApi.md#TranslationsExcludeCollection) | **Patch** /projects/{project_id}/translations/exclude | Set exclude from export flag on translations selected by query
[**TranslationsIncludeCollection**](TranslationsApi.md#TranslationsIncludeCollection) | **Patch** /projects/{project_id}/translations/include | Remove exlude from import flag from translations selected by query
[**TranslationsList**](TranslationsApi.md#TranslationsList) | **Get** /projects/{project_id}/translations | List all translations
[**TranslationsReviewCollection**](TranslationsApi.md#TranslationsReviewCollection) | **Patch** /projects/{project_id}/translations/review | Review translations selected by query
[**TranslationsSearch**](TranslationsApi.md#TranslationsSearch) | **Post** /projects/{project_id}/translations/search | Search translations
[**TranslationsUnverifyCollection**](TranslationsApi.md#TranslationsUnverifyCollection) | **Patch** /projects/{project_id}/translations/unverify | Mark translations selected by query as unverified
[**TranslationsVerifyCollection**](TranslationsApi.md#TranslationsVerifyCollection) | **Patch** /projects/{project_id}/translations/verify | Verify translations selected by query



## TranslationCreate

> TranslationDetails TranslationCreate(ctx, projectId, translationCreateParameters, optional)

Create a translation

Create a translation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**translationCreateParameters** | [**TranslationCreateParameters**](TranslationCreateParameters.md)|  | 
 **optional** | ***TranslationCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**TranslationDetails**](translation_details.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationExclude

> TranslationDetails TranslationExclude(ctx, projectId, id, translationExcludeParameters, optional)

Exclude a translation from export

Set exclude from export flag on an existing translation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
**translationExcludeParameters** | [**TranslationExcludeParameters**](TranslationExcludeParameters.md)|  | 
 **optional** | ***TranslationExcludeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationExcludeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**TranslationDetails**](translation_details.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationInclude

> TranslationDetails TranslationInclude(ctx, projectId, id, translationIncludeParameters, optional)

Revoke exclusion of a translation in export

Remove exclude from export flag from an existing translation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
**translationIncludeParameters** | [**TranslationIncludeParameters**](TranslationIncludeParameters.md)|  | 
 **optional** | ***TranslationIncludeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationIncludeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**TranslationDetails**](translation_details.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationReview

> TranslationDetails TranslationReview(ctx, projectId, id, translationReviewParameters, optional)

Review a translation

Mark an existing translation as reviewed.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
**translationReviewParameters** | [**TranslationReviewParameters**](TranslationReviewParameters.md)|  | 
 **optional** | ***TranslationReviewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationReviewOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**TranslationDetails**](translation_details.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationShow

> TranslationDetails TranslationShow(ctx, projectId, id, optional)

Get a single translation

Get details on a single translation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
 **optional** | ***TranslationShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **branch** | **optional.String**| specify the branch to use | 

### Return type

[**TranslationDetails**](translation_details.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationUnverify

> TranslationDetails TranslationUnverify(ctx, projectId, id, translationUnverifyParameters, optional)

Mark a translation as unverified

Mark an existing translation as unverified.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
**translationUnverifyParameters** | [**TranslationUnverifyParameters**](TranslationUnverifyParameters.md)|  | 
 **optional** | ***TranslationUnverifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationUnverifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**TranslationDetails**](translation_details.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationUpdate

> TranslationDetails TranslationUpdate(ctx, projectId, id, translationUpdateParameters, optional)

Update a translation

Update an existing translation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
**translationUpdateParameters** | [**TranslationUpdateParameters**](TranslationUpdateParameters.md)|  | 
 **optional** | ***TranslationUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**TranslationDetails**](translation_details.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationVerify

> TranslationDetails TranslationVerify(ctx, projectId, id, translationVerifyParameters, optional)

Verify a translation

Verify an existing translation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
**translationVerifyParameters** | [**TranslationVerifyParameters**](TranslationVerifyParameters.md)|  | 
 **optional** | ***TranslationVerifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationVerifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**TranslationDetails**](translation_details.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationsByKey

> []Translation TranslationsByKey(ctx, projectId, keyId, optional)

List translations by key

List translations for a specific key.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keyId** | **string**| Translation Key ID | 
 **optional** | ***TranslationsByKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationsByKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **page** | **optional.Int32**| Page number | 
 **perPage** | **optional.Int32**| allows you to specify a page size up to 100 items, 25 by default | 
 **branch** | **optional.String**| specify the branch to use | 
 **sort** | **optional.String**| Sort criteria. Can be one of: key_name, created_at, updated_at. | 
 **order** | **optional.String**| Order direction. Can be one of: asc, desc. | 
 **q** | **optional.String**| Specify a query to find translations by content (including wildcards).&lt;br&gt;&lt;br&gt; The following qualifiers are supported in the query:&lt;br&gt; &lt;ul&gt;   &lt;li&gt;&lt;code&gt;id:translation_id,...&lt;/code&gt; for queries on a comma-separated list of ids&lt;/li&gt;   &lt;li&gt;&lt;code&gt;unverified:{true|false}&lt;/code&gt; for verification status&lt;/li&gt;   &lt;li&gt;&lt;code&gt;tags:XYZ&lt;/code&gt; for tags on the translation&lt;/li&gt;   &lt;li&gt;&lt;code&gt;excluded:{true|false}&lt;/code&gt; for exclusion status&lt;/li&gt;   &lt;li&gt;&lt;code&gt;updated_at:{&gt;&#x3D;|&lt;&#x3D;}2013-02-21T00:00:00Z&lt;/code&gt; for date range queries&lt;/li&gt; &lt;/ul&gt; Find more examples &lt;a href&#x3D;\&quot;#overview--usage-examples\&quot;&gt;here&lt;/a&gt;.  | 

### Return type

[**[]Translation**](translation.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationsByLocale

> []Translation TranslationsByLocale(ctx, projectId, localeId, optional)

List translations by locale

List translations for a specific locale. If you want to download all translations for one locale we recommend to use the <code>locales#download</code> endpoint.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**localeId** | **string**| Locale ID | 
 **optional** | ***TranslationsByLocaleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationsByLocaleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **page** | **optional.Int32**| Page number | 
 **perPage** | **optional.Int32**| allows you to specify a page size up to 100 items, 25 by default | 
 **branch** | **optional.String**| specify the branch to use | 
 **sort** | **optional.String**| Sort criteria. Can be one of: key_name, created_at, updated_at. | 
 **order** | **optional.String**| Order direction. Can be one of: asc, desc. | 
 **q** | **optional.String**| Specify a query to find translations by content (including wildcards).&lt;br&gt;&lt;br&gt; The following qualifiers are supported in the query:&lt;br&gt; &lt;ul&gt;   &lt;li&gt;&lt;code&gt;id:translation_id,...&lt;/code&gt; for queries on a comma-separated list of ids&lt;/li&gt;   &lt;li&gt;&lt;code&gt;unverified:{true|false}&lt;/code&gt; for verification status&lt;/li&gt;   &lt;li&gt;&lt;code&gt;tags:XYZ&lt;/code&gt; for tags on the translation&lt;/li&gt;   &lt;li&gt;&lt;code&gt;excluded:{true|false}&lt;/code&gt; for exclusion status&lt;/li&gt;   &lt;li&gt;&lt;code&gt;updated_at:{&gt;&#x3D;|&lt;&#x3D;}2013-02-21T00:00:00Z&lt;/code&gt; for date range queries&lt;/li&gt; &lt;/ul&gt; Find more examples &lt;a href&#x3D;\&quot;#overview--usage-examples\&quot;&gt;here&lt;/a&gt;.  | 

### Return type

[**[]Translation**](translation.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationsExcludeCollection

> AffectedCount TranslationsExcludeCollection(ctx, projectId, translationsExcludeParameters, optional)

Set exclude from export flag on translations selected by query

Exclude translations matching query from locale export.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**translationsExcludeParameters** | [**TranslationsExcludeParameters**](TranslationsExcludeParameters.md)|  | 
 **optional** | ***TranslationsExcludeCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationsExcludeCollectionOpts struct


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


## TranslationsIncludeCollection

> AffectedCount TranslationsIncludeCollection(ctx, projectId, translationsIncludeParameters, optional)

Remove exlude from import flag from translations selected by query

Include translations matching query in locale export.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**translationsIncludeParameters** | [**TranslationsIncludeParameters**](TranslationsIncludeParameters.md)|  | 
 **optional** | ***TranslationsIncludeCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationsIncludeCollectionOpts struct


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


## TranslationsList

> []Translation TranslationsList(ctx, projectId, optional)

List all translations

List translations for the given project. If you want to download all translations for one locale we recommend to use the <code>locales#download</code> endpoint.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
 **optional** | ***TranslationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **page** | **optional.Int32**| Page number | 
 **perPage** | **optional.Int32**| allows you to specify a page size up to 100 items, 25 by default | 
 **branch** | **optional.String**| specify the branch to use | 
 **sort** | **optional.String**| Sort criteria. Can be one of: key_name, created_at, updated_at. | 
 **order** | **optional.String**| Order direction. Can be one of: asc, desc. | 
 **q** | **optional.String**| Specify a query to find translations by content (including wildcards).&lt;br&gt;&lt;br&gt; The following qualifiers are supported in the query:&lt;br&gt; &lt;ul&gt;   &lt;li&gt;&lt;code&gt;id:translation_id,...&lt;/code&gt; for queries on a comma-separated list of ids&lt;/li&gt;   &lt;li&gt;&lt;code&gt;tags:XYZ&lt;/code&gt; for tags on the translation&lt;/li&gt;   &lt;li&gt;&lt;code&gt;unverified:{true|false}&lt;/code&gt; for verification status&lt;/li&gt;   &lt;li&gt;&lt;code&gt;excluded:{true|false}&lt;/code&gt; for exclusion status&lt;/li&gt;   &lt;li&gt;&lt;code&gt;updated_at:{&gt;&#x3D;|&lt;&#x3D;}2013-02-21T00:00:00Z&lt;/code&gt; for date range queries&lt;/li&gt; &lt;/ul&gt; Find more examples &lt;a href&#x3D;\&quot;#overview--usage-examples\&quot;&gt;here&lt;/a&gt;.  | 

### Return type

[**[]Translation**](translation.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationsReviewCollection

> AffectedCount TranslationsReviewCollection(ctx, projectId, translationsReviewParameters, optional)

Review translations selected by query

Review translations matching query.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**translationsReviewParameters** | [**TranslationsReviewParameters**](TranslationsReviewParameters.md)|  | 
 **optional** | ***TranslationsReviewCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationsReviewCollectionOpts struct


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


## TranslationsSearch

> []Translation TranslationsSearch(ctx, projectId, translationsSearchParameters, optional)

Search translations

Search translations for the given project. Provides the same search interface as <code>translations#index</code> but allows POST requests to avoid limitations imposed by GET requests. If you want to download all translations for one locale we recommend to use the <code>locales#download</code> endpoint.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**translationsSearchParameters** | [**TranslationsSearchParameters**](TranslationsSearchParameters.md)|  | 
 **optional** | ***TranslationsSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationsSearchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
 **page** | **optional.Int32**| Page number | 
 **perPage** | **optional.Int32**| allows you to specify a page size up to 100 items, 25 by default | 

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


## TranslationsUnverifyCollection

> AffectedCount TranslationsUnverifyCollection(ctx, projectId, translationsUnverifyParameters, optional)

Mark translations selected by query as unverified

Mark translations matching query as unverified.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**translationsUnverifyParameters** | [**TranslationsUnverifyParameters**](TranslationsUnverifyParameters.md)|  | 
 **optional** | ***TranslationsUnverifyCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationsUnverifyCollectionOpts struct


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


## TranslationsVerifyCollection

> AffectedCount TranslationsVerifyCollection(ctx, projectId, translationsVerifyParameters, optional)

Verify translations selected by query

Verify translations matching query.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**translationsVerifyParameters** | [**TranslationsVerifyParameters**](TranslationsVerifyParameters.md)|  | 
 **optional** | ***TranslationsVerifyCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TranslationsVerifyCollectionOpts struct


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

