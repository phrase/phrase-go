# phrase.ICUApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IcuSkeleton**](ICUApi.md#IcuSkeleton) | **Post** /icu/skeleton | Build ICU skeletons



## IcuSkeleton

> Icu IcuSkeleton(ctx, icuSkeletonParameters, optional)

Build ICU skeletons

Generates ICU (International Components for Unicode) message format skeletons for a given source string across one or more locales. An ICU skeleton strips the literal text from a pluralized or select message while preserving its structural rules — argument names, plural categories, select cases, and ordinal forms — adjusted to the pluralization rules of each requested locale.  Use this endpoint to normalize translation templates before importing them into locale files, or to validate that a source string carries the plural forms required by a target language.  Either `content` or `id` must be provided — supplying both or neither returns 400. When `id` is used and the referenced translation does not exist, the endpoint returns 404. When the source string is not valid ICU message format syntax, the endpoint returns 422 with an `error` field describing the parse failure. 

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

