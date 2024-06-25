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

// LocaleDownloadsApiService LocaleDownloadsApi service
type LocaleDownloadsApiService service

// LocaleDownloadCreateOpts Optional parameters for the method 'LocaleDownloadCreate'
type LocaleDownloadCreateOpts struct {
	XPhraseAppOTP   optional.String `json:"X-PhraseApp-OTP,omitempty"`
	IfModifiedSince optional.String `json:"If-Modified-Since,omitempty"`
	IfNoneMatch     optional.String `json:"If-None-Match,omitempty"`
}

/*
LocaleDownloadCreate Initiate async download of a locale
Prepare a locale for download in a specific file format.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param projectId Project ID
  - @param localeId Locale ID
  - @param localeDownloadCreateParameters
  - @param optional nil or *LocaleDownloadCreateOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)
  - @param "IfModifiedSince" (optional.String) -  Last modified condition, see <a href=\"#overview--conditional-get-requests--http-caching\">Conditional GET requests / HTTP Caching</a> (optional)
  - @param "IfNoneMatch" (optional.String) -  ETag condition, see <a href=\"#overview--conditional-get-requests--http-caching\">Conditional GET requests / HTTP Caching</a> (optional)

@return LocaleDownload
*/
func (a *LocaleDownloadsApiService) LocaleDownloadCreate(ctx _context.Context, projectId string, localeId string, localeDownloadCreateParameters LocaleDownloadCreateParameters, localVarOptionals *LocaleDownloadCreateOpts) (LocaleDownload, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  LocaleDownload
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/projects/{project_id}/locales/{locale_id}/downloads"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", _neturl.QueryEscape(parameterToString(projectId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"locale_id"+"}", _neturl.QueryEscape(parameterToString(localeId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	if localVarOptionals != nil && localVarOptionals.IfModifiedSince.IsSet() {
		localVarHeaderParams["If-Modified-Since"] = parameterToString(localVarOptionals.IfModifiedSince.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.IfNoneMatch.IsSet() {
		localVarHeaderParams["If-None-Match"] = parameterToString(localVarOptionals.IfNoneMatch.Value(), "")
	}
	// body params
	localVarPostBody = &localeDownloadCreateParameters
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

// LocaleDownloadShowOpts Optional parameters for the method 'LocaleDownloadShow'
type LocaleDownloadShowOpts struct {
	XPhraseAppOTP   optional.String `json:"X-PhraseApp-OTP,omitempty"`
	IfModifiedSince optional.String `json:"If-Modified-Since,omitempty"`
	IfNoneMatch     optional.String `json:"If-None-Match,omitempty"`
}

/*
LocaleDownloadShow Show status of an async locale download
Show status of already started async locale download. If the download is finished, the download link will be returned.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param projectId Project ID
  - @param localeId Locale ID
  - @param id ID
  - @param optional nil or *LocaleDownloadShowOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)
  - @param "IfModifiedSince" (optional.String) -  Last modified condition, see <a href=\"#overview--conditional-get-requests--http-caching\">Conditional GET requests / HTTP Caching</a> (optional)
  - @param "IfNoneMatch" (optional.String) -  ETag condition, see <a href=\"#overview--conditional-get-requests--http-caching\">Conditional GET requests / HTTP Caching</a> (optional)

@return LocaleDownload
*/
func (a *LocaleDownloadsApiService) LocaleDownloadShow(ctx _context.Context, projectId string, localeId string, id string, localVarOptionals *LocaleDownloadShowOpts) (LocaleDownload, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  LocaleDownload
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/projects/{project_id}/locales/{locale_id}/downloads/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", _neturl.QueryEscape(parameterToString(projectId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"locale_id"+"}", _neturl.QueryEscape(parameterToString(localeId, "")), -1)

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
	if localVarOptionals != nil && localVarOptionals.IfModifiedSince.IsSet() {
		localVarHeaderParams["If-Modified-Since"] = parameterToString(localVarOptionals.IfModifiedSince.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.IfNoneMatch.IsSet() {
		localVarHeaderParams["If-None-Match"] = parameterToString(localVarOptionals.IfNoneMatch.Value(), "")
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
