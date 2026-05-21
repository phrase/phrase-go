# phrase.ProjectsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectCreate**](ProjectsApi.md#ProjectCreate) | **Post** /projects | Create a project
[**ProjectDelete**](ProjectsApi.md#ProjectDelete) | **Delete** /projects/{id} | Delete a project
[**ProjectShow**](ProjectsApi.md#ProjectShow) | **Get** /projects/{id} | Get a single project
[**ProjectUpdate**](ProjectsApi.md#ProjectUpdate) | **Patch** /projects/{id} | Update a project
[**ProjectsList**](ProjectsApi.md#ProjectsList) | **Get** /projects | List projects



## ProjectCreate

> ProjectDetails ProjectCreate(ctx, projectCreateParameters, optional)

Create a project

Create a new project in the given account.  When `source_project_id` is supplied, the new project is created as a clone of that project. All locales, keys, and translations are copied asynchronously after the response is returned, so they may not be available immediately. Settings from the source project are inherited unless explicitly overridden in the request; in clone mode, the `shares_translation_memory` field is ignored and inherited from the source.  `shares_translation_memory` defaults to `true` when omitted on a non-clone create. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectCreateParameters** | [**ProjectCreateParameters**](ProjectCreateParameters.md)|  | 
 **optional** | ***ProjectCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProjectCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**ProjectDetails**](ProjectDetails.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectDelete

> ProjectDelete(ctx, id, optional)

Delete a project

Delete an existing project. Associated repository syncs and OTA distributions are removed. A `project:delete` event is dispatched. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| ID | 
 **optional** | ***ProjectDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProjectDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectShow

> ProjectDetails ProjectShow(ctx, id, optional)

Get a single project

Get details on a single project.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| ID | 
 **optional** | ***ProjectShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProjectShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**ProjectDetails**](ProjectDetails.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectUpdate

> ProjectDetails ProjectUpdate(ctx, id, projectUpdateParameters, optional)

Update a project

Update an existing project.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| ID | 
**projectUpdateParameters** | [**ProjectUpdateParameters**](ProjectUpdateParameters.md)|  | 
 **optional** | ***ProjectUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProjectUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**ProjectDetails**](ProjectDetails.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsList

> []Project ProjectsList(ctx, optional)

List projects

List all projects the current user has access to.  When the `account_id` query parameter is omitted, the response includes projects across every account the user is a member of. Pass `account_id` to scope the results to a single account. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProjectsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**page** | **optional.Int32**| Page number | 
**perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 
**accountId** | **optional.String**| Filter by Account ID | 
**sortBy** | **optional.String**| Sort projects. Valid values are &#x60;name_asc&#x60;, &#x60;name_desc&#x60;, &#x60;updated_at_asc&#x60;, &#x60;updated_at_desc&#x60;, &#x60;space_asc&#x60;, and &#x60;space_desc&#x60;. The trailing direction segment is optional; if omitted or invalid, projects are sorted ascending. Any other value is ignored and the default ordering is returned. | 
**filters** | [**optional.Interface of []string**](string.md)| Filter projects. The only supported value is &#x60;favorites&#x60;, which restricts the results to projects the current user has starred. | 
**q** | **optional.String**| Search query. The only supported syntax is &#x60;name:&lt;text&gt;&#x60; — for example &#x60;name:android&#x60; returns projects whose name matches &#x60;android&#x60; (case-insensitive substring). Any value that does not match the &#x60;name:&#x60; prefix is ignored. | 

### Return type

[**[]Project**](Project.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

