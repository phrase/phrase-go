# \ICUApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IcuSkeleton**](ICUApi.md#IcuSkeleton) | **Post** /icu/skeleton | Build icu skeletons



## IcuSkeleton

> Icu IcuSkeleton(ctx, icuSkeletonParameters, optional)

Build icu skeletons

Returns icu skeletons for multiple locale codes based on a source content.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**icuSkeletonParameters** | [**IcuSkeletonParameters**](IcuSkeletonParameters.md)|  | 
 **optional** | ***IcuSkeletonOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IcuSkeletonOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**Icu**](Icu.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

