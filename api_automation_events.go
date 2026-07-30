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

// AutomationEventsApiService AutomationEventsApi service
type AutomationEventsApiService service

// AccountAutomationEventsListOpts Optional parameters for the method 'AccountAutomationEventsList'
type AccountAutomationEventsListOpts struct {
	XPhraseAppOTP optional.String `json:"X-PhraseApp-OTP,omitempty"`
	Page          optional.Int32  `json:"page,omitempty"`
	PerPage       optional.Int32  `json:"per_page,omitempty"`
	AutomationId  optional.String `json:"automation_id,omitempty"`
	State         optional.String `json:"state,omitempty"`
	TriggeredBy   optional.String `json:"triggered_by,omitempty"`
	ProjectId     optional.String `json:"project_id,omitempty"`
	ProjectIds    []string        `json:"project_ids,omitempty"`
	CreatedAfter  optional.String `json:"created_after,omitempty"`
	CreatedBefore optional.String `json:"created_before,omitempty"`
}

/*
AccountAutomationEventsList List automation events for an account
Returns the run history across all automations in the account, newest-first.  Use &#x60;automation_id&#x60; to narrow results to a single automation. Use &#x60;project_id&#x60; or &#x60;project_ids&#x60; to narrow by project.  For feature availability, see [Jobs (Strings)](https://support.phrase.com/hc/en-us/articles/5784100517788-Jobs-Strings).
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param accountId Account ID
  - @param optional nil or *AccountAutomationEventsListOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)
  - @param "Page" (optional.Int32) -  Page number
  - @param "PerPage" (optional.Int32) -  Limit on the number of objects to be returned, between 1 and 100. 25 by default
  - @param "AutomationId" (optional.String) -  Filter events to a single automation by its ID.
  - @param "State" (optional.String) -  Filter events by outcome state. Unrecognized values are ignored.
  - @param "TriggeredBy" (optional.String) -  Filter events by what triggered the automation run. Unrecognized values are ignored.
  - @param "ProjectId" (optional.String) -  Filter events by project ID. Accepts a single ID or a comma-separated list of IDs.
  - @param "ProjectIds" (optional.Interface of []string) -  Filter events by one or more project IDs.
  - @param "CreatedAfter" (optional.String) -  Return only events created after this ISO 8601 timestamp. Returns 400 if the value is not a valid date-time.
  - @param "CreatedBefore" (optional.String) -  Return only events created before this ISO 8601 timestamp. Returns 400 if the value is not a valid date-time.

@return []AutomationEvent
*/
func (a *AutomationEventsApiService) AccountAutomationEventsList(ctx _context.Context, accountId string, localVarOptionals *AccountAutomationEventsListOpts) ([]AutomationEvent, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []AutomationEvent
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/accounts/{account_id}/automation_events"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", _neturl.QueryEscape(parameterToString(accountId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PerPage.IsSet() {
		localVarQueryParams.Add("per_page", parameterToString(localVarOptionals.PerPage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AutomationId.IsSet() {
		localVarQueryParams.Add("automation_id", parameterToString(localVarOptionals.AutomationId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.State.IsSet() {
		localVarQueryParams.Add("state", parameterToString(localVarOptionals.State.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TriggeredBy.IsSet() {
		localVarQueryParams.Add("triggered_by", parameterToString(localVarOptionals.TriggeredBy.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProjectId.IsSet() {
		localVarQueryParams.Add("project_id", parameterToString(localVarOptionals.ProjectId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProjectIds != nil {
		t := localVarOptionals.ProjectIds
		for i := range t {
			localVarQueryParams.Add("project_ids[]", parameterToString(t[i], "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.CreatedAfter.IsSet() {
		localVarQueryParams.Add("created_after", parameterToString(localVarOptionals.CreatedAfter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CreatedBefore.IsSet() {
		localVarQueryParams.Add("created_before", parameterToString(localVarOptionals.CreatedBefore.Value(), ""))
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

// AutomationEventsListOpts Optional parameters for the method 'AutomationEventsList'
type AutomationEventsListOpts struct {
	XPhraseAppOTP optional.String `json:"X-PhraseApp-OTP,omitempty"`
	Page          optional.Int32  `json:"page,omitempty"`
	PerPage       optional.Int32  `json:"per_page,omitempty"`
	State         optional.String `json:"state,omitempty"`
	TriggeredBy   optional.String `json:"triggered_by,omitempty"`
	ProjectId     optional.String `json:"project_id,omitempty"`
	ProjectIds    []string        `json:"project_ids,omitempty"`
	CreatedAfter  optional.String `json:"created_after,omitempty"`
	CreatedBefore optional.String `json:"created_before,omitempty"`
}

/*
AutomationEventsList List events for an automation
Returns the run history for a specific automation, newest-first.  For feature availability, see [Jobs (Strings)](https://support.phrase.com/hc/en-us/articles/5784100517788-Jobs-Strings).
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param accountId Account ID
  - @param id ID
  - @param optional nil or *AutomationEventsListOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)
  - @param "Page" (optional.Int32) -  Page number
  - @param "PerPage" (optional.Int32) -  Limit on the number of objects to be returned, between 1 and 100. 25 by default
  - @param "State" (optional.String) -  Filter events by outcome state. Unrecognized values are ignored.
  - @param "TriggeredBy" (optional.String) -  Filter events by what triggered the automation run. Unrecognized values are ignored.
  - @param "ProjectId" (optional.String) -  Filter events by project ID. Accepts a single ID or a comma-separated list of IDs.
  - @param "ProjectIds" (optional.Interface of []string) -  Filter events by one or more project IDs.
  - @param "CreatedAfter" (optional.String) -  Return only events created after this ISO 8601 timestamp. Returns 400 if the value is not a valid date-time.
  - @param "CreatedBefore" (optional.String) -  Return only events created before this ISO 8601 timestamp. Returns 400 if the value is not a valid date-time.

@return []AutomationEvent
*/
func (a *AutomationEventsApiService) AutomationEventsList(ctx _context.Context, accountId string, id string, localVarOptionals *AutomationEventsListOpts) ([]AutomationEvent, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []AutomationEvent
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/accounts/{account_id}/automations/{automation_id}/events"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", _neturl.QueryEscape(parameterToString(accountId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.QueryEscape(parameterToString(id, "")), -1)

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
	if localVarOptionals != nil && localVarOptionals.TriggeredBy.IsSet() {
		localVarQueryParams.Add("triggered_by", parameterToString(localVarOptionals.TriggeredBy.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProjectId.IsSet() {
		localVarQueryParams.Add("project_id", parameterToString(localVarOptionals.ProjectId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProjectIds != nil {
		t := localVarOptionals.ProjectIds
		for i := range t {
			localVarQueryParams.Add("project_ids[]", parameterToString(t[i], "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.CreatedAfter.IsSet() {
		localVarQueryParams.Add("created_after", parameterToString(localVarOptionals.CreatedAfter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CreatedBefore.IsSet() {
		localVarQueryParams.Add("created_before", parameterToString(localVarOptionals.CreatedBefore.Value(), ""))
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
