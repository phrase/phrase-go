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

// ReportsApiService ReportsApi service
type ReportsApiService service

// ReportLocalesListOpts Optional parameters for the method 'ReportLocalesList'
type ReportLocalesListOpts struct {
	XPhraseAppOTP optional.String `json:"X-PhraseApp-OTP,omitempty"`
	Page          optional.Int32  `json:"page,omitempty"`
	PerPage       optional.Int32  `json:"per_page,omitempty"`
	LocaleCodes   optional.String `json:"locale_codes,omitempty"`
	Tag           optional.String `json:"tag,omitempty"`
	Branch        optional.String `json:"branch,omitempty"`
}

/*
ReportLocalesList List Locale Reports
List all locale reports for the given project
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param projectId Project ID
  - @param optional nil or *ReportLocalesListOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)
  - @param "Page" (optional.Int32) -  Page number
  - @param "PerPage" (optional.Int32) -  Limit on the number of objects to be returned, between 1 and 100. 25 by default
  - @param "LocaleCodes" (optional.String) -  Locale Code
  - @param "Tag" (optional.String) -  tag
  - @param "Branch" (optional.String) -  specify the branch to use

@return []LocaleReport
*/
func (a *ReportsApiService) ReportLocalesList(ctx _context.Context, projectId string, localVarOptionals *ReportLocalesListOpts) ([]LocaleReport, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []LocaleReport
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/projects/{project_id}/report/locales"
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
	if localVarOptionals != nil && localVarOptionals.LocaleCodes.IsSet() {
		localVarQueryParams.Add("locale_codes", parameterToString(localVarOptionals.LocaleCodes.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Tag.IsSet() {
		localVarQueryParams.Add("tag", parameterToString(localVarOptionals.Tag.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Branch.IsSet() {
		localVarQueryParams.Add("branch", parameterToString(localVarOptionals.Branch.Value(), ""))
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

// ReportShowOpts Optional parameters for the method 'ReportShow'
type ReportShowOpts struct {
	XPhraseAppOTP optional.String `json:"X-PhraseApp-OTP,omitempty"`
	Branch        optional.String `json:"branch,omitempty"`
}

/*
ReportShow Get Project Report
Get report of a single project.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param projectId Project ID
  - @param optional nil or *ReportShowOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)
  - @param "Branch" (optional.String) -  specify the branch to use

@return ProjectReport
*/
func (a *ReportsApiService) ReportShow(ctx _context.Context, projectId string, localVarOptionals *ReportShowOpts) (ProjectReport, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ProjectReport
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/projects/{project_id}/report"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", _neturl.QueryEscape(parameterToString(projectId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Branch.IsSet() {
		localVarQueryParams.Add("branch", parameterToString(localVarOptionals.Branch.Value(), ""))
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
