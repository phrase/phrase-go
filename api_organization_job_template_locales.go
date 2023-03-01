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

// OrganizationJobTemplateLocalesApiService OrganizationJobTemplateLocalesApi service
type OrganizationJobTemplateLocalesApiService service

// OrganizationJobTemplateLocaleDeleteOpts Optional parameters for the method 'OrganizationJobTemplateLocaleDelete'
type OrganizationJobTemplateLocaleDeleteOpts struct {
	XPhraseAppOTP optional.String `json:"X-PhraseApp-OTP,omitempty"`
}

/*
OrganizationJobTemplateLocaleDelete Delete an organization job template locale
Delete an existing organization job template locale.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param accountId Account ID
  - @param jobTemplateId Job Template ID
  - @param jobTemplateLocaleId Job Template Locale ID
  - @param optional nil or *OrganizationJobTemplateLocaleDeleteOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)
*/
func (a *OrganizationJobTemplateLocalesApiService) OrganizationJobTemplateLocaleDelete(ctx _context.Context, accountId string, jobTemplateId string, jobTemplateLocaleId string, localVarOptionals *OrganizationJobTemplateLocaleDeleteOpts) ([]byte, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/accounts/{account_id}/job_templates/{job_template_id}/locales/{job_template_locale_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", _neturl.QueryEscape(parameterToString(accountId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"job_template_id"+"}", _neturl.QueryEscape(parameterToString(jobTemplateId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"job_template_locale_id"+"}", _neturl.QueryEscape(parameterToString(jobTemplateLocaleId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarBody, localVarHTTPResponse, newErr
	}

	return localVarBody, localVarHTTPResponse, nil
}

// OrganizationJobTemplateLocaleShowOpts Optional parameters for the method 'OrganizationJobTemplateLocaleShow'
type OrganizationJobTemplateLocaleShowOpts struct {
	XPhraseAppOTP optional.String `json:"X-PhraseApp-OTP,omitempty"`
}

/*
OrganizationJobTemplateLocaleShow Get a single organization job template locale
Get a single job template locale for a given organization job template.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param accountId Account ID
  - @param jobTemplateId Job Template ID
  - @param jobTemplateLocaleId Job Template Locale ID
  - @param optional nil or *OrganizationJobTemplateLocaleShowOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)

@return JobTemplateLocales
*/
func (a *OrganizationJobTemplateLocalesApiService) OrganizationJobTemplateLocaleShow(ctx _context.Context, accountId string, jobTemplateId string, jobTemplateLocaleId string, localVarOptionals *OrganizationJobTemplateLocaleShowOpts) (JobTemplateLocales, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  JobTemplateLocales
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/accounts/{account_id}/job_templates/{job_template_id}/locales/{job_template_locale_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", _neturl.QueryEscape(parameterToString(accountId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"job_template_id"+"}", _neturl.QueryEscape(parameterToString(jobTemplateId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"job_template_locale_id"+"}", _neturl.QueryEscape(parameterToString(jobTemplateLocaleId, "")), -1)

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

// OrganizationJobTemplateLocaleUpdateOpts Optional parameters for the method 'OrganizationJobTemplateLocaleUpdate'
type OrganizationJobTemplateLocaleUpdateOpts struct {
	XPhraseAppOTP optional.String `json:"X-PhraseApp-OTP,omitempty"`
}

/*
OrganizationJobTemplateLocaleUpdate Update an organization job template locale
Update an existing organization job template locale.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param accountId Account ID
  - @param jobTemplateId Job Template ID
  - @param jobTemplateLocaleId Job Template Locale ID
  - @param organizationJobTemplateLocaleUpdateParameters
  - @param optional nil or *OrganizationJobTemplateLocaleUpdateOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)

@return JobTemplateLocales
*/
func (a *OrganizationJobTemplateLocalesApiService) OrganizationJobTemplateLocaleUpdate(ctx _context.Context, accountId string, jobTemplateId string, jobTemplateLocaleId string, organizationJobTemplateLocaleUpdateParameters OrganizationJobTemplateLocaleUpdateParameters, localVarOptionals *OrganizationJobTemplateLocaleUpdateOpts) (JobTemplateLocales, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  JobTemplateLocales
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/accounts/{account_id}/job_templates/{job_template_id}/locales/{job_template_locale_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", _neturl.QueryEscape(parameterToString(accountId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"job_template_id"+"}", _neturl.QueryEscape(parameterToString(jobTemplateId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"job_template_locale_id"+"}", _neturl.QueryEscape(parameterToString(jobTemplateLocaleId, "")), -1)

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
	// body params
	localVarPostBody = &organizationJobTemplateLocaleUpdateParameters
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

// OrganizationJobTemplateLocalesCreateOpts Optional parameters for the method 'OrganizationJobTemplateLocalesCreate'
type OrganizationJobTemplateLocalesCreateOpts struct {
	XPhraseAppOTP optional.String `json:"X-PhraseApp-OTP,omitempty"`
}

/*
OrganizationJobTemplateLocalesCreate Create an organization job template locale
Create a new organization job template locale.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param accountId Account ID
  - @param jobTemplateId Job Template ID
  - @param organizationJobTemplateLocalesCreateParameters
  - @param optional nil or *OrganizationJobTemplateLocalesCreateOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)

@return JobTemplateLocales
*/
func (a *OrganizationJobTemplateLocalesApiService) OrganizationJobTemplateLocalesCreate(ctx _context.Context, accountId string, jobTemplateId string, organizationJobTemplateLocalesCreateParameters OrganizationJobTemplateLocalesCreateParameters, localVarOptionals *OrganizationJobTemplateLocalesCreateOpts) (JobTemplateLocales, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  JobTemplateLocales
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/accounts/{account_id}/job_templates/{job_template_id}/locales"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", _neturl.QueryEscape(parameterToString(accountId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"job_template_id"+"}", _neturl.QueryEscape(parameterToString(jobTemplateId, "")), -1)

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
	// body params
	localVarPostBody = &organizationJobTemplateLocalesCreateParameters
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

// OrganizationJobTemplateLocalesListOpts Optional parameters for the method 'OrganizationJobTemplateLocalesList'
type OrganizationJobTemplateLocalesListOpts struct {
	XPhraseAppOTP optional.String `json:"X-PhraseApp-OTP,omitempty"`
	Page          optional.Int32  `json:"page,omitempty"`
	PerPage       optional.Int32  `json:"per_page,omitempty"`
}

/*
OrganizationJobTemplateLocalesList List organization job template locales
List all job template locales for a given organization job template.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param accountId Account ID
  - @param jobTemplateId Job Template ID
  - @param optional nil or *OrganizationJobTemplateLocalesListOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)
  - @param "Page" (optional.Int32) -  Page number
  - @param "PerPage" (optional.Int32) -  Limit on the number of objects to be returned, between 1 and 100. 25 by default

@return []JobTemplateLocales
*/
func (a *OrganizationJobTemplateLocalesApiService) OrganizationJobTemplateLocalesList(ctx _context.Context, accountId string, jobTemplateId string, localVarOptionals *OrganizationJobTemplateLocalesListOpts) ([]JobTemplateLocales, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []JobTemplateLocales
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/accounts/{account_id}/job_templates/{job_template_id}/locales"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", _neturl.QueryEscape(parameterToString(accountId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"job_template_id"+"}", _neturl.QueryEscape(parameterToString(jobTemplateId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PerPage.IsSet() {
		localVarQueryParams.Add("per_page", parameterToString(localVarOptionals.PerPage.Value(), ""))
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
