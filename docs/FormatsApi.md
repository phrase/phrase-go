# phrase.FormatsApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FormatsList**](FormatsApi.md#FormatsList) | **Get** /formats | List formats



## FormatsList

> []Format FormatsList(ctx, )

List formats

Returns all file formats that Phrase Strings supports. Use the api_name value from each format as the file_format parameter when uploading or downloading locale files. Not every format supports both directions: check the importable and exportable fields before using a format in a workflow. This endpoint does not require authentication and is not subject to rate limiting. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Format**](Format.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

