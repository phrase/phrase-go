# phrase.BranchesApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BranchCompare**](BranchesApi.md#BranchCompare) | **Get** /projects/{project_id}/branches/{name}/compare | Compare branches
[**BranchComparisonCreate**](BranchesApi.md#BranchComparisonCreate) | **Post** /projects/{project_id}/branches/{name}/compare | Create comparison (async.)
[**BranchCreate**](BranchesApi.md#BranchCreate) | **Post** /projects/{project_id}/branches | Create a branch
[**BranchDelete**](BranchesApi.md#BranchDelete) | **Delete** /projects/{project_id}/branches/{name} | Delete a branch
[**BranchMerge**](BranchesApi.md#BranchMerge) | **Patch** /projects/{project_id}/branches/{name}/merge | Merge a branch
[**BranchShow**](BranchesApi.md#BranchShow) | **Get** /projects/{project_id}/branches/{name} | Get a single branch
[**BranchSync**](BranchesApi.md#BranchSync) | **Patch** /projects/{project_id}/branches/{name}/sync | Sync a branch
[**BranchUpdate**](BranchesApi.md#BranchUpdate) | **Patch** /projects/{project_id}/branches/{name} | Update a branch
[**BranchesList**](BranchesApi.md#BranchesList) | **Get** /projects/{project_id}/branches | List branches



## BranchCompare

> BranchComparison BranchCompare(ctx, projectId, name, optional)

Compare branches

Compare branch with main branch.  *Note: Comparing a branch may take several minutes depending on the project size. Consider using the `POST /compare` endpoint for creating comparison asynchronously.* 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**name** | **string**| name | 
 **optional** | ***BranchCompareOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BranchCompareOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**BranchComparison**](BranchComparison.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BranchComparisonCreate

> BranchComparisonCreate(ctx, projectId, name, branchCreateComparisonParameters, optional)

Create comparison (async.)

Create a branch comparison asynchronously. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**name** | **string**| name | 
**branchCreateComparisonParameters** | [**BranchCreateComparisonParameters**](BranchCreateComparisonParameters.md)|  | 
 **optional** | ***BranchComparisonCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BranchComparisonCreateOpts struct


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


## BranchCreate

> Branch BranchCreate(ctx, projectId, branchCreateParameters, optional)

Create a branch

Create a new branch.  Branch project provisioning runs asynchronously, so the newly created branch is returned in a transitional state (typically `creating_branch`) and only reaches `success` once the underlying project has been set up. Poll the branch resource until its `state` becomes `success` before performing further operations on it.  Requires the Branching feature to be enabled on the account.  *Note: Creating a new branch may take several minutes depending on the project size.* 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**branchCreateParameters** | [**BranchCreateParameters**](BranchCreateParameters.md)|  | 
 **optional** | ***BranchCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BranchCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**Branch**](Branch.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BranchDelete

> BranchDelete(ctx, projectId, name, optional)

Delete a branch

Delete an existing branch.  A branch cannot be deleted while it still has open jobs or open translation orders attached to its branch project — in that case the request is rejected with `409 Conflict`. A branch whose current `state` does not allow deletion (for example, while a merge or sync is in progress) is rejected with `422 Unprocessable Entity`.  Requires the Branching feature to be enabled on the account. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**name** | **string**| name | 
 **optional** | ***BranchDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BranchDeleteOpts struct


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


## BranchMerge

> BranchMerge(ctx, projectId, name, branchMergeParameters, optional)

Merge a branch

Merge an existing branch back into its base branch.  The merge runs asynchronously. The branch transitions to `merging_branch` and settles in `merged`, `merge_error`, or `merge_conflict` once the background job completes; the response body for this request is empty. Poll the branch resource to observe the final state.  A branch cannot be merged while it still has open jobs or open translation orders attached to its branch project — in that case the request is rejected with `409 Conflict`. A branch whose current `state` does not allow a merge is rejected with `422 Unprocessable Entity`.  Requires the Branching feature to be enabled on the account.  *Note: Merging a branch may take several minutes depending on diff size.* 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**name** | **string**| name | 
**branchMergeParameters** | [**BranchMergeParameters**](BranchMergeParameters.md)|  | 
 **optional** | ***BranchMergeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BranchMergeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BranchShow

> Branch BranchShow(ctx, projectId, name, optional)

Get a single branch

Get details on a single branch for a given project.  Requires the Branching feature to be enabled on the account. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**name** | **string**| name | 
 **optional** | ***BranchShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BranchShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**Branch**](Branch.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BranchSync

> BranchSync(ctx, projectId, name, branchSyncParameters, optional)

Sync a branch

Pull changes from the base branch into this branch, applying the chosen conflict-resolution strategy.  The sync runs asynchronously. The branch transitions to `syncing_branch` and settles back into `success` (or `merge_conflict` / `branch_error`) once the background job completes; the response body for this request is empty. Poll the branch resource to observe the final state.  Only branches created with the newer branching system can be synced. Requests against branches from the older system, or against branches whose current state does not allow a sync, are rejected with `422 Unprocessable Entity` and an empty body.  Requires the Branching feature to be enabled on the account. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**name** | **string**| name | 
**branchSyncParameters** | [**BranchSyncParameters**](BranchSyncParameters.md)|  | 
 **optional** | ***BranchSyncOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BranchSyncOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BranchUpdate

> Branch BranchUpdate(ctx, projectId, name, branchUpdateParameters, optional)

Update a branch

Update an existing branch. Only the branch name can be changed.  Requires the Branching feature to be enabled on the account. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**name** | **string**| name | 
**branchUpdateParameters** | [**BranchUpdateParameters**](BranchUpdateParameters.md)|  | 
 **optional** | ***BranchUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BranchUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**Branch**](Branch.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BranchesList

> []Branch BranchesList(ctx, projectId, optional)

List branches

List all branches of the current project.  Requires the Branching feature to be enabled on the account. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
 **optional** | ***BranchesListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BranchesListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 
**page** | **optional.Int32**| Page number | 
**perPage** | **optional.Int32**| Limit on the number of objects to be returned, between 1 and 100. 25 by default | 

### Return type

[**[]Branch**](Branch.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

