# \QualityPerformanceScoreApi

All URIs are relative to *https://api.phrase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectsQualityPerformanceScore**](QualityPerformanceScoreApi.md#ProjectsQualityPerformanceScore) | **Post** /projects/{project_id}/quality_performance_score | Get Translation Quality



## ProjectsQualityPerformanceScore

> ProjectsQualityPerformanceScore200Response ProjectsQualityPerformanceScore(ctx, projectId, projectsQualityPerformanceScoreRequest, optional)

Get Translation Quality

Retrieves the quality scores for your Strings translations. Returns a score, measured by Phrase QPS

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project ID | 
**projectsQualityPerformanceScoreRequest** | [**ProjectsQualityPerformanceScoreRequest**](ProjectsQualityPerformanceScoreRequest.md)|  | 
 **optional** | ***ProjectsQualityPerformanceScoreOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProjectsQualityPerformanceScoreOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPhraseAppOTP** | **optional.String**| Two-Factor-Authentication token (optional) | 

### Return type

[**ProjectsQualityPerformanceScore200Response**](ProjectsQualityPerformanceScore200Response.md)

### Authorization

[Basic](../README.md#Basic), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

