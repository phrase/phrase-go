package phrase

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// ChecksApiService ChecksApi service
type ChecksApiService service

// CheckIssueDismissOpts Optional parameters for the method 'CheckIssueDismiss'
type CheckIssueDismissOpts struct {
	XPhraseAppOTP optional.String `json:"X-PhraseApp-OTP,omitempty"`
}

/*
CheckIssueDismiss Dismiss a check issue
**Note:** The Checks API is still in development and might change in subsequent releases.  Mark a check issue as dismissed so it no longer appears on the list of active check issues.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param projectId Project ID
  - @param id Check Issue ID
  - @param optional nil or *CheckIssueDismissOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)

@return CheckIssue
*/
func (a *ChecksApiService) CheckIssueDismiss(ctx _context.Context, projectId string, id string, localVarOptionals *CheckIssueDismissOpts) (CheckIssue, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CheckIssue
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/projects/{project_id}/checks/issues/{id}/dismiss"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", _neturl.QueryEscape(parameterToString(projectId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.QueryEscape(parameterToString(id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XPhraseAppOTP.IsSet() {
		localVarHeaderParams["X-PhraseApp-OTP"] = parameterToString(localVarOptionals.XPhraseAppOTP.Value(), "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v DocumentDelete422Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// CheckIssuesListOpts Optional parameters for the method 'CheckIssuesList'
type CheckIssuesListOpts struct {
	XPhraseAppOTP optional.String `json:"X-PhraseApp-OTP,omitempty"`
	Page          optional.Int32  `json:"page,omitempty"`
	PerPage       optional.Int32  `json:"per_page,omitempty"`
	State         optional.String `json:"state,omitempty"`
	LocaleIds     []string        `json:"locale_ids,omitempty"`
	CheckNames    []string        `json:"check_names,omitempty"`
}

/*
CheckIssuesList List check issues
**Note:** The Checks API is still in development and might change in subsequent releases.  List check issues for the given project. Results can be filtered by locale, check name, and state.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param projectId Project ID
  - @param optional nil or *CheckIssuesListOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)
  - @param "Page" (optional.Int32) -  Page number
  - @param "PerPage" (optional.Int32) -  Limit on the number of objects to be returned, between 1 and 100. 25 by default
  - @param "State" (optional.String) -  Filter by state of the check issue. Can be one of: `active`, `solved`, `dismissed`, `all`. Defaults to `active`.
  - @param "LocaleIds" (optional.Interface of []string) -  Filter by one or more locale IDs.
  - @param "CheckNames" (optional.Interface of []string) -  Filter by one or more check names. Valid values are:  - `translation_content_length` — the translation exceeds the maximum character limit configured for the key. - `translation_placeholder_usage` — the translation is missing placeholders present in the source, or contains unexpected ones. - `translation_glossary_usage` — the translation does not follow the glossary term translations.

@return []CheckIssue
*/
func (a *ChecksApiService) CheckIssuesList(ctx _context.Context, projectId string, localVarOptionals *CheckIssuesListOpts) ([]CheckIssue, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []CheckIssue
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/projects/{project_id}/checks/issues"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", _neturl.QueryEscape(parameterToString(projectId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PerPage.IsSet() {
		localVarQueryParams.Add("per_page", parameterToString(localVarOptionals.PerPage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.State.IsSet() {
		localVarQueryParams.Add("state", parameterToString(localVarOptionals.State.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LocaleIds != nil {
		t := localVarOptionals.LocaleIds
		for i := range t {
			localVarQueryParams.Add("locale_ids[]", parameterToString(t[i], "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.CheckNames != nil {
		t := localVarOptionals.CheckNames
		for i := range t {
			localVarQueryParams.Add("check_names[]", parameterToString(t[i], "multi"))
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XPhraseAppOTP.IsSet() {
		localVarHeaderParams["X-PhraseApp-OTP"] = parameterToString(localVarOptionals.XPhraseAppOTP.Value(), "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
