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

// LinkedKeysApiService LinkedKeysApi service
type LinkedKeysApiService service

// KeyLinksBatchDestroyOpts Optional parameters for the method 'KeyLinksBatchDestroy'
type KeyLinksBatchDestroyOpts struct {
	XPhraseAppOTP                  optional.String    `json:"X-PhraseApp-OTP,omitempty"`
	KeyLinksBatchDestroyParameters optional.Interface `json:"KeyLinksBatchDestroyParameters,omitempty"`
}

/*
KeyLinksBatchDestroy Batch unlink child keys from a parent key
Removes one or more child keys from a parent key&#39;s linked-key group, or dissolves the entire group by setting unlink_parent to true.  Use this when you need to detach specific child keys from a shared translation source, or to fully break apart a linked-key group so each key manages its own translations independently. When child keys are unlinked, their translations are updated with a copy of the parent&#39;s current content (strategy keep_content, the default) or cleared (strategy remove_content).  This operation is only available on main projects. It returns 422 when a child key in &#x60;child_key_ids&#x60; is not currently linked to the parent, or when a translation update fails while unlinking.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param projectId Project ID
  - @param id Parent Translation Key ID
  - @param optional nil or *KeyLinksBatchDestroyOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)
  - @param "KeyLinksBatchDestroyParameters" (optional.Interface of KeyLinksBatchDestroyParameters) -

@return KeyLink
*/
func (a *LinkedKeysApiService) KeyLinksBatchDestroy(ctx _context.Context, projectId string, id string, localVarOptionals *KeyLinksBatchDestroyOpts) (KeyLink, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  KeyLink
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/projects/{project_id}/keys/{id}/key_links"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", _neturl.QueryEscape(parameterToString(projectId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.QueryEscape(parameterToString(id, "")), -1)

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
	if localVarOptionals != nil && localVarOptionals.KeyLinksBatchDestroyParameters.IsSet() {
		localVarOptionalKeyLinksBatchDestroyParameters, localVarOptionalKeyLinksBatchDestroyParametersok := localVarOptionals.KeyLinksBatchDestroyParameters.Value().(KeyLinksBatchDestroyParameters)
		if !localVarOptionalKeyLinksBatchDestroyParametersok {
			return localVarReturnValue, nil, reportError("keyLinksBatchDestroyParameters should be KeyLinksBatchDestroyParameters")
		}
		localVarPostBody = &localVarOptionalKeyLinksBatchDestroyParameters
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

// KeyLinksCreateOpts Optional parameters for the method 'KeyLinksCreate'
type KeyLinksCreateOpts struct {
	XPhraseAppOTP optional.String `json:"X-PhraseApp-OTP,omitempty"`
}

/*
KeyLinksCreate Link child keys to a parent key
Designates a translation key as a parent and links one or more child keys to it. Once linked, child keys receive a special reference marker as their translation content, signalling that their translations are derived from the parent. Use this when you want to group related keys — for example, a short label and its long-form variant — so translators see them in context together.  Pass an empty child_key_ids array to mark the key as a parent without linking any children yet. Both the parent key and every child key must belong to the main project; branch keys cannot participate in key links. A child key can have at most one parent at a time; attempting to link a child that already has a parent returns a 422 error with code CHILD_IS_ALREADY_LINKED. Parent and child key plurality must match — linking a plural child to a non-plural parent (or vice versa) also returns a 422.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param projectId Project ID
  - @param id Parent Translation Key ID
  - @param keyLinksCreateParameters
  - @param optional nil or *KeyLinksCreateOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)

@return KeyLink
*/
func (a *LinkedKeysApiService) KeyLinksCreate(ctx _context.Context, projectId string, id string, keyLinksCreateParameters KeyLinksCreateParameters, localVarOptionals *KeyLinksCreateOpts) (KeyLink, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  KeyLink
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/projects/{project_id}/keys/{id}/key_links"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", _neturl.QueryEscape(parameterToString(projectId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.QueryEscape(parameterToString(id, "")), -1)

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
	localVarPostBody = &keyLinksCreateParameters
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

// KeyLinksDestroyOpts Optional parameters for the method 'KeyLinksDestroy'
type KeyLinksDestroyOpts struct {
	XPhraseAppOTP optional.String `json:"X-PhraseApp-OTP,omitempty"`
}

/*
KeyLinksDestroy Unlink a child key from a parent key
Removes a single child key from a parent key&#39;s link group. A link group is the relationship model that keeps child keys synchronized with a parent: while linked, a child key&#39;s translations are derived from the parent&#39;s content. When you call this endpoint, the child key leaves the group and becomes independent — its existing translations are updated with the parent&#39;s current content and then marked unverified, signalling that reviewers should confirm the content is still appropriate for the child&#39;s context.  Use this endpoint when you need to detach one specific child key while keeping other children linked. To detach multiple children at once, use the batch unlink endpoint. This operation is only available on main projects.  It returns 422 when the child key is not currently linked to the specified parent key, or when a translation update fails during the unlink process.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param projectId Project ID
  - @param id Parent Translation Key ID
  - @param childKeyId The ID of the child translation key to unlink from the parent.
  - @param optional nil or *KeyLinksDestroyOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)
*/
func (a *LinkedKeysApiService) KeyLinksDestroy(ctx _context.Context, projectId string, id string, childKeyId string, localVarOptionals *KeyLinksDestroyOpts) ([]byte, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/projects/{project_id}/keys/{id}/key_links/{child_key_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", _neturl.QueryEscape(parameterToString(projectId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.QueryEscape(parameterToString(id, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"child_key_id"+"}", _neturl.QueryEscape(parameterToString(childKeyId, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 422 {
			var v DocumentDelete422Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return nil, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return nil, localVarHTTPResponse, newErr
		}
		return localVarBody, localVarHTTPResponse, newErr
	}

	return localVarBody, localVarHTTPResponse, nil
}

// KeyLinksIndexOpts Optional parameters for the method 'KeyLinksIndex'
type KeyLinksIndexOpts struct {
	XPhraseAppOTP optional.String `json:"X-PhraseApp-OTP,omitempty"`
}

/*
KeyLinksIndex List child keys of a parent key
Returns the key link record for a parent key, including all child keys associated with it. Key linking lets translation keys share translations — a child key inherits content from its designated parent. Use this endpoint to inspect which keys are linked under a given parent before unlinking them or auditing translation consistency across related keys.  The key identified by &#x60;id&#x60; must be designated as a parent key (it must have at least one child key linked to it). Listing the links of a key that is not a parent returns 400.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param projectId Project ID
  - @param id Parent Translation Key ID
  - @param optional nil or *KeyLinksIndexOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)

@return KeyLink
*/
func (a *LinkedKeysApiService) KeyLinksIndex(ctx _context.Context, projectId string, id string, localVarOptionals *KeyLinksIndexOpts) (KeyLink, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  KeyLink
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/projects/{project_id}/keys/{id}/key_links"
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
