# RepoSync

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**Name** | Pointer to **NullableString** | Optional custom display name for this repo sync. When null or blank, the sync is displayed using the associated project name.  | [optional] 
**Project** | [**ProjectShort**](ProjectShort.md) |  | [optional] 
**Provider** | **string** |  | [optional] 
**Enabled** | **bool** |  | [optional] 
**AutoImport** | **bool** |  | [optional] 
**RepoName** | **string** |  | [optional] 
**PrBranch** | Pointer to **NullableString** | Branch used as the source of exports/PRs. May be &#x60;null&#x60; when the sync is configured to push directly to &#x60;base_branch&#x60;.  | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**LastImportAt** | Pointer to [**NullableTime**](time.Time.md) |  | [optional] 
**LastExportAt** | Pointer to [**NullableTime**](time.Time.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


