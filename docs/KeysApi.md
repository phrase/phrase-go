# phrase.KeysApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KeyCreate**](KeysApi.md#KeyCreate) | **Post** /projects/{project_id}/keys | Create a key
[**KeyDelete**](KeysApi.md#KeyDelete) | **Delete** /projects/{project_id}/keys/{id} | Delete a key
[**KeyShow**](KeysApi.md#KeyShow) | **Get** /projects/{project_id}/keys/{id} | Get a single key
[**KeyUpdate**](KeysApi.md#KeyUpdate) | **Patch** /projects/{project_id}/keys/{id} | Update a key
[**KeysDeleteCollection**](KeysApi.md#KeysDeleteCollection) | **Delete** /projects/{project_id}/keys | Delete collection of keys
[**KeysExclude**](KeysApi.md#KeysExclude) | **Patch** /projects/{project_id}/keys/exclude | Exclude a locale on a collection of keys
[**KeysInclude**](KeysApi.md#KeysInclude) | **Patch** /projects/{project_id}/keys/include | Include a locale on a collection of keys
[**KeysList**](KeysApi.md#KeysList) | **Get** /projects/{project_id}/keys | List keys
[**KeysSearch**](KeysApi.md#KeysSearch) | **Post** /projects/{project_id}/keys/search | Search keys
[**KeysTag**](KeysApi.md#KeysTag) | **Patch** /projects/{project_id}/keys/tag | Add tags to collection of keys
[**KeysUntag**](KeysApi.md#KeysUntag) | **Patch** /projects/{project_id}/keys/untag | Remove tags from collection of keys



## KeyCreate

> TranslationKeyDetails KeyCreate(ctx, projectId, keyCreateParameters, optional)

Create a key

Create a new key.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keyCreateParameters** | [**KeyCreateParameters**](KeyCreateParameters.md)|  | 
 **optional** | ***KeyCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KeyCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**TranslationKeyDetails**](TranslationKeyDetails.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyDelete

> KeyDelete(ctx, projectId, id, optional)

Delete a key

Delete an existing key.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
 **optional** | ***KeyDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KeyDeleteOpts struct


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


## KeyShow

> TranslationKeyDetails KeyShow(ctx, projectId, id, optional)

Get a single key

Get details on a single key for a given project.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
 **optional** | ***KeyShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KeyShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**branch** | **optional.String**| specify the branch to use | 

### Return type

[**TranslationKeyDetails**](TranslationKeyDetails.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyUpdate

> TranslationKeyDetails KeyUpdate(ctx, projectId, id, keyUpdateParameters, optional)

Update a key

Update an existing key.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**id** | **string**| ID | 
**keyUpdateParameters** | [**KeyUpdateParameters**](KeyUpdateParameters.md)|  | 
 **optional** | ***KeyUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KeyUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**TranslationKeyDetails**](TranslationKeyDetails.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeysDeleteCollection

> AffectedResources KeysDeleteCollection(ctx, projectId, optional)

Delete collection of keys

Delete all keys matching query. Same constraints as list. Please limit the number of affected keys to about 1,000 as you might experience timeouts otherwise.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
 **optional** | ***KeysDeleteCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KeysDeleteCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**branch** | **optional.String**| specify the branch to use | 
**q** | **optional.String**| Specify a query to do broad search for keys by name (including wildcards).  The following qualifiers are also supported in the search term:  - &#x60;ids:key_id,...&#x60; for queries on a comma-separated list of ids - &#x60;name:key_name&#x60; for text queries on exact key names - spaces, commas, and colons  need to be escaped with double backslashes - &#x60;tags:tag_name&#x60; to filter for keys with certain tags - &#x60;translated:{true|false}&#x60; for translation status (also requires &#x60;locale_id&#x60; to be specified) - &#x60;updated_at:{&gt;&#x3D;|&lt;&#x3D;}2013-02-21T00:00:00Z&#x60; for date range queries - &#x60;unmentioned_in_upload:upload_id,...&#x60; to filter keys unmentioned within upload. When multiple upload IDs provided, matches only keys not mentioned in **all** uploads  **Caution:** Query parameters with empty values will be treated as though they are not included in the request and will be ignored.  Find more examples [here](/en/api/strings/usage-examples).  | 
**localeId** | **optional.String**| Locale used to determine the translation state of a key when filtering for untranslated or translated keys. | 

### Return type

[**AffectedResources**](AffectedResources.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeysExclude

> AffectedResources KeysExclude(ctx, projectId, keysExcludeParameters, optional)

Exclude a locale on a collection of keys

Exclude a locale on keys matching query. Same constraints as list.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keysExcludeParameters** | [**KeysExcludeParameters**](KeysExcludeParameters.md)|  | 
 **optional** | ***KeysExcludeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KeysExcludeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**AffectedResources**](AffectedResources.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeysInclude

> AffectedResources KeysInclude(ctx, projectId, keysIncludeParameters, optional)

Include a locale on a collection of keys

Include a locale on keys matching query. Same constraints as list.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keysIncludeParameters** | [**KeysIncludeParameters**](KeysIncludeParameters.md)|  | 
 **optional** | ***KeysIncludeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KeysIncludeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**AffectedResources**](AffectedResources.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeysList

> []TranslationKey KeysList(ctx, projectId, optional)

List keys

List all keys for the given project. Alternatively you can POST requests to /search.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
 **optional** | ***KeysListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KeysListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**page** | **optional.Int32**| Page number | 
**perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 
**branch** | **optional.String**| specify the branch to use | 
**sort** | **optional.String**| Sort by field. Can be one of: name, created_at, updated_at. | 
**order** | **optional.String**| Order direction. Can be one of: asc, desc. | 
**q** | **optional.String**| Specify a query to do broad search for keys by name (including wildcards).  The following qualifiers are also supported in the search term:  - &#x60;ids:key_id,...&#x60; for queries on a comma-separated list of ids - &#x60;name:key_name&#x60; for text queries on exact key names - spaces, commas, and colons  need to be escaped with double backslashes - &#x60;tags:tag_name&#x60; to filter for keys with certain tags - &#x60;translated:{true|false}&#x60; for translation status (also requires &#x60;locale_id&#x60; to be specified) - &#x60;updated_at:{&gt;&#x3D;|&lt;&#x3D;}2013-02-21T00:00:00Z&#x60; for date range queries - &#x60;unmentioned_in_upload:upload_id,...&#x60; to filter keys unmentioned within upload. When multiple upload IDs provided, matches only keys not mentioned in **all** uploads  **Caution:** Query parameters with empty values will be treated as though they are not included in the request and will be ignored.  Find more examples [here](/en/api/strings/usage-examples).  | 
**localeId** | **optional.String**| Locale used to determine the translation state of a key when filtering for untranslated or translated keys. | 

### Return type

[**[]TranslationKey**](TranslationKey.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeysSearch

> []TranslationKey KeysSearch(ctx, projectId, keysSearchParameters, optional)

Search keys

Search keys for the given project matching query.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keysSearchParameters** | [**KeysSearchParameters**](KeysSearchParameters.md)|  | 
 **optional** | ***KeysSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KeysSearchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**page** | **optional.Int32**| Page number | 
**perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 

### Return type

[**[]TranslationKey**](TranslationKey.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeysTag

> AffectedResources KeysTag(ctx, projectId, keysTagParameters, optional)

Add tags to collection of keys

Tags all keys matching query. Same constraints as list.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keysTagParameters** | [**KeysTagParameters**](KeysTagParameters.md)|  | 
 **optional** | ***KeysTagOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KeysTagOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**AffectedResources**](AffectedResources.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeysUntag

> AffectedResources KeysUntag(ctx, projectId, keysUntagParameters, optional)

Remove tags from collection of keys

Removes specified tags from keys matching query.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**keysUntagParameters** | [**KeysUntagParameters**](KeysUntagParameters.md)|  | 
 **optional** | ***KeysUntagOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KeysUntagOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**AffectedResources**](AffectedResources.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

