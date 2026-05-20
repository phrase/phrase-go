# Branch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseProjectId** | **string** |  | [optional] 
**BranchProjectId** | **string** |  | [optional] 
**Name** | **string** |  | [optional] 
**Base** | Pointer to **NullableString** | Name of the base branch this branch was created from. Only present for branches created with the newer branching system. | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**MergedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**MergedBy** | [**UserPreview**](UserPreview.md) |  | [optional] 
**CreatedBy** | [**UserPreview**](UserPreview.md) |  | [optional] 
**State** | **string** |  | [optional] 
**ChildBranches** | **[]string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


