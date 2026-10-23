# RepoSyncEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**Type** | **string** |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**Status** | **string** |  | [optional] 
**PullRequestUrl** | **string** | URL of the pull request created on export | [optional] 
**AutoImport** | **bool** | Whether the import was triggered by the repo push event | [optional] 
**Errors** | **[]string** | List of error messages, in case of failure | [optional] 
**Uploads** | Pointer to [**[]RepoSyncEventUploadsInner**](RepoSyncEventUploadsInner.md) | The uploads associated with this import event, or &#x60;null&#x60; if no upload batch has been created yet. Each entry identifies a single file upload by its public code (&#x60;id&#x60;) and its current public state (&#x60;state&#x60;). Present on show responses only; absent from list responses. Possible &#x60;state&#x60; values: &#x60;initialized&#x60;, &#x60;waiting_for_preview&#x60;, &#x60;processing&#x60;, &#x60;success&#x60;, &#x60;error&#x60;.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


