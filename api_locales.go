package phrase

import (
	"fmt"

	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"os"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// LocalesApiService LocalesApi service
type LocalesApiService service

// AccountLocalesOpts Optional parameters for the method 'AccountLocales'
type AccountLocalesOpts struct {
	XPhraseAppOTP optional.String `json:"X-PhraseApp-OTP,omitempty"`
	Page          optional.Int32  `json:"page,omitempty"`
	PerPage       optional.Int32  `json:"per_page,omitempty"`
}

/*
AccountLocales List locales used in account
List all locales unique by locale code used across all projects within an account.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id ID
  - @param optional nil or *AccountLocalesOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)
  - @param "Page" (optional.Int32) -  Page number
  - @param "PerPage" (optional.Int32) -  Limit on the number of objects to be returned, between 1 and 100. 25 by default

@return []LocalePreview1
*/
func (a *LocalesApiService) AccountLocales(ctx _context.Context, id string, localVarOptionals *AccountLocalesOpts) ([]LocalePreview1, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []LocalePreview1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/accounts/{id}/locales"
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

// LocaleCreateOpts Optional parameters for the method 'LocaleCreate'
type LocaleCreateOpts struct {
	XPhraseAppOTP optional.String `json:"X-PhraseApp-OTP,omitempty"`
}

/*
LocaleCreate Create a locale
Create a new locale.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param projectId Project ID
  - @param localeCreateParameters
  - @param optional nil or *LocaleCreateOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)

@return LocaleDetails
*/
func (a *LocalesApiService) LocaleCreate(ctx _context.Context, projectId string, localeCreateParameters LocaleCreateParameters, localVarOptionals *LocaleCreateOpts) (LocaleDetails, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  LocaleDetails
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/projects/{project_id}/locales"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", _neturl.QueryEscape(parameterToString(projectId, "")), -1)

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
	localVarPostBody = &localeCreateParameters
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

// LocaleDeleteOpts Optional parameters for the method 'LocaleDelete'
type LocaleDeleteOpts struct {
	XPhraseAppOTP optional.String `json:"X-PhraseApp-OTP,omitempty"`
	Branch        optional.String `json:"branch,omitempty"`
}

/*
LocaleDelete Delete a locale
Delete an existing locale.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param projectId Project ID
  - @param id Locale ID or locale name
  - @param optional nil or *LocaleDeleteOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)
  - @param "Branch" (optional.String) -  specify the branch to use
*/
func (a *LocalesApiService) LocaleDelete(ctx _context.Context, projectId string, id string, localVarOptionals *LocaleDeleteOpts) ([]byte, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/projects/{project_id}/locales/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", _neturl.QueryEscape(parameterToString(projectId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.QueryEscape(parameterToString(id, "")), -1)

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

// LocaleDownloadOpts Optional parameters for the method 'LocaleDownload'
type LocaleDownloadOpts struct {
	XPhraseAppOTP                 optional.String    `json:"X-PhraseApp-OTP,omitempty"`
	IfModifiedSince               optional.String    `json:"If-Modified-Since,omitempty"`
	IfNoneMatch                   optional.String    `json:"If-None-Match,omitempty"`
	Branch                        optional.String    `json:"branch,omitempty"`
	FileFormat                    optional.String    `json:"file_format,omitempty"`
	Tags                          optional.String    `json:"tags,omitempty"`
	Tag                           optional.String    `json:"tag,omitempty"`
	IncludeEmptyTranslations      optional.Bool      `json:"include_empty_translations,omitempty"`
	ExcludeEmptyZeroForms         optional.Bool      `json:"exclude_empty_zero_forms,omitempty"`
	IncludeTranslatedKeys         optional.Bool      `json:"include_translated_keys,omitempty"`
	KeepNotranslateTags           optional.Bool      `json:"keep_notranslate_tags,omitempty"`
	ConvertEmoji                  optional.Bool      `json:"convert_emoji,omitempty"`
	FormatOptions                 optional.Interface `json:"format_options,omitempty"`
	Encoding                      optional.String    `json:"encoding,omitempty"`
	SkipUnverifiedTranslations    optional.Bool      `json:"skip_unverified_translations,omitempty"`
	IncludeUnverifiedTranslations optional.Bool      `json:"include_unverified_translations,omitempty"`
	UseLastReviewedVersion        optional.Bool      `json:"use_last_reviewed_version,omitempty"`
	FallbackLocaleId              optional.String    `json:"fallback_locale_id,omitempty"`
	SourceLocaleId                optional.String    `json:"source_locale_id,omitempty"`
	TranslationKeyPrefix          optional.String    `json:"translation_key_prefix,omitempty"`
	FilterByPrefix                optional.Bool      `json:"filter_by_prefix,omitempty"`
	CustomMetadataFilters         optional.Interface `json:"custom_metadata_filters,omitempty"`
	LocaleIds                     []string           `json:"locale_ids,omitempty"`
	UpdatedSince                  optional.String    `json:"updated_since,omitempty"`
}

/*
LocaleDownload Download a locale
Download a locale in a specific file format.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param projectId Project ID
  - @param id Locale ID or locale name
  - @param optional nil or *LocaleDownloadOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)
  - @param "IfModifiedSince" (optional.String) -  Last modified condition, see [Conditional GET requests / HTTP Caching](/en/api/strings/pagination#conditional-get-requests-%2F-http-caching) (optional)
  - @param "IfNoneMatch" (optional.String) -  ETag condition, see [Conditional GET requests / HTTP Caching](/en/api/strings/pagination#conditional-get-requests-%2F-http-caching) (optional)
  - @param "Branch" (optional.String) -  specify the branch to use
  - @param "FileFormat" (optional.String) -  File format name. See the [format guide](https://support.phrase.com/hc/en-us/sections/6111343326364) for all supported file formats.
  - @param "Tags" (optional.String) -  Limit results to keys tagged with a list of comma separated tag names.
  - @param "Tag" (optional.String) -  Limit download to tagged keys. This parameter is deprecated. Please use the \"tags\" parameter instead
  - @param "IncludeEmptyTranslations" (optional.Bool) -  Indicates whether keys without translations should be included in the output as well.
  - @param "ExcludeEmptyZeroForms" (optional.Bool) -  Indicates whether zero forms should be included when empty in pluralized keys.
  - @param "IncludeTranslatedKeys" (optional.Bool) -  Include translated keys in the locale file. Use in combination with include_empty_translations to obtain only untranslated keys.
  - @param "KeepNotranslateTags" (optional.Bool) -  Indicates whether [NOTRANSLATE] tags should be kept.
  - @param "ConvertEmoji" (optional.Bool) -  This option is obsolete. Projects that were created on or after Nov 29th 2019 or that did not contain emoji by then will not require this flag any longer since emoji are now supported natively.
  - @param "FormatOptions" (optional.Interface of map[string]interface{}) -  Additional formatting and render options. See the [format guide](https://support.phrase.com/hc/en-us/sections/6111343326364) for a list of options available for each format. Specify format options like this: `...&format_options[foo]=bar`
  - @param "Encoding" (optional.String) -  Enforces a specific encoding on the file contents. Valid options are \"UTF-8\", \"UTF-16\" and \"ISO-8859-1\".
  - @param "SkipUnverifiedTranslations" (optional.Bool) -  Indicates whether the locale file should skip all unverified translations. This parameter is deprecated and should be replaced with `include_unverified_translations`.
  - @param "IncludeUnverifiedTranslations" (optional.Bool) -  if set to false unverified translations are excluded
  - @param "UseLastReviewedVersion" (optional.Bool) -  If set to true the last reviewed version of a translation is used. This is only available if the review workflow is enabled for the project.
  - @param "FallbackLocaleId" (optional.String) -  If a key has no translation in the locale being downloaded the translation in the fallback locale will be used. Provide the ID of the locale that should be used as the fallback. Requires include_empty_translations to be set to `true`.
  - @param "SourceLocaleId" (optional.String) -  Provides the source language of a corresponding job as the source language of the generated locale file. This parameter will be ignored unless used in combination with a `tag` parameter indicating a specific job.
  - @param "TranslationKeyPrefix" (optional.String) -  Download all translation keys, and remove the specified prefix where possible. Warning: this may create duplicate key names if other keys share the same name after the prefix is removed.
  - @param "FilterByPrefix" (optional.Bool) -  Only download translation keys containing the specified prefix, and remove the prefix from the generated file.
  - @param "CustomMetadataFilters" (optional.Interface of map[string]interface{}) -  Custom metadata filters. Provide the name of the metadata field and the value to filter by. Only keys with matching metadata will be included in the download.
  - @param "LocaleIds" (optional.Interface of []string) -  Locale IDs or locale names
  - @param "UpdatedSince" (optional.String) -  Only include keys that have been updated since the given date. The date must be in ISO 8601 format (e.g., `2023-01-01T00:00:00Z`).

@return *os.File
*/
func (a *LocalesApiService) LocaleDownload(ctx _context.Context, projectId string, id string, localVarOptionals *LocaleDownloadOpts) (*os.File, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/projects/{project_id}/locales/{id}/download"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", _neturl.QueryEscape(parameterToString(projectId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.QueryEscape(parameterToString(id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Branch.IsSet() {
		localVarQueryParams.Add("branch", parameterToString(localVarOptionals.Branch.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileFormat.IsSet() {
		localVarQueryParams.Add("file_format", parameterToString(localVarOptionals.FileFormat.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Tags.IsSet() {
		localVarQueryParams.Add("tags", parameterToString(localVarOptionals.Tags.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Tag.IsSet() {
		localVarQueryParams.Add("tag", parameterToString(localVarOptionals.Tag.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeEmptyTranslations.IsSet() {
		localVarQueryParams.Add("include_empty_translations", parameterToString(localVarOptionals.IncludeEmptyTranslations.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ExcludeEmptyZeroForms.IsSet() {
		localVarQueryParams.Add("exclude_empty_zero_forms", parameterToString(localVarOptionals.ExcludeEmptyZeroForms.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeTranslatedKeys.IsSet() {
		localVarQueryParams.Add("include_translated_keys", parameterToString(localVarOptionals.IncludeTranslatedKeys.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.KeepNotranslateTags.IsSet() {
		localVarQueryParams.Add("keep_notranslate_tags", parameterToString(localVarOptionals.KeepNotranslateTags.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ConvertEmoji.IsSet() {
		localVarQueryParams.Add("convert_emoji", parameterToString(localVarOptionals.ConvertEmoji.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FormatOptions.IsSet() {
		for key, value := range localVarOptionals.FormatOptions.Value().(map[string]interface{}) {
			localVarQueryParams = serializeMapParams(fmt.Sprintf("format_options[%s]", key), value, localVarQueryParams)
		}
	}
	if localVarOptionals != nil && localVarOptionals.Encoding.IsSet() {
		localVarQueryParams.Add("encoding", parameterToString(localVarOptionals.Encoding.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SkipUnverifiedTranslations.IsSet() {
		localVarQueryParams.Add("skip_unverified_translations", parameterToString(localVarOptionals.SkipUnverifiedTranslations.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeUnverifiedTranslations.IsSet() {
		localVarQueryParams.Add("include_unverified_translations", parameterToString(localVarOptionals.IncludeUnverifiedTranslations.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UseLastReviewedVersion.IsSet() {
		localVarQueryParams.Add("use_last_reviewed_version", parameterToString(localVarOptionals.UseLastReviewedVersion.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FallbackLocaleId.IsSet() {
		localVarQueryParams.Add("fallback_locale_id", parameterToString(localVarOptionals.FallbackLocaleId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SourceLocaleId.IsSet() {
		localVarQueryParams.Add("source_locale_id", parameterToString(localVarOptionals.SourceLocaleId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TranslationKeyPrefix.IsSet() {
		localVarQueryParams.Add("translation_key_prefix", parameterToString(localVarOptionals.TranslationKeyPrefix.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterByPrefix.IsSet() {
		localVarQueryParams.Add("filter_by_prefix", parameterToString(localVarOptionals.FilterByPrefix.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CustomMetadataFilters.IsSet() {
		for key, value := range localVarOptionals.CustomMetadataFilters.Value().(map[string]interface{}) {
			localVarQueryParams = serializeMapParams(fmt.Sprintf("custom_metadata_filters[%s]", key), value, localVarQueryParams)
		}
	}
	if localVarOptionals != nil && localVarOptionals.LocaleIds != nil {
		t := localVarOptionals.LocaleIds
		for i := range t {
			localVarQueryParams.Add("locale_ids[]", parameterToString(t[i], "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.UpdatedSince.IsSet() {
		localVarQueryParams.Add("updated_since", parameterToString(localVarOptionals.UpdatedSince.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*"}

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

// LocaleShowOpts Optional parameters for the method 'LocaleShow'
type LocaleShowOpts struct {
	XPhraseAppOTP optional.String `json:"X-PhraseApp-OTP,omitempty"`
	Branch        optional.String `json:"branch,omitempty"`
}

/*
LocaleShow Get a single locale
Get details on a single locale for a given project.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param projectId Project ID
  - @param id Locale ID or locale name
  - @param optional nil or *LocaleShowOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)
  - @param "Branch" (optional.String) -  specify the branch to use

@return LocaleDetails
*/
func (a *LocalesApiService) LocaleShow(ctx _context.Context, projectId string, id string, localVarOptionals *LocaleShowOpts) (LocaleDetails, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  LocaleDetails
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/projects/{project_id}/locales/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", _neturl.QueryEscape(parameterToString(projectId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.QueryEscape(parameterToString(id, "")), -1)

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

// LocaleUpdateOpts Optional parameters for the method 'LocaleUpdate'
type LocaleUpdateOpts struct {
	XPhraseAppOTP optional.String `json:"X-PhraseApp-OTP,omitempty"`
}

/*
LocaleUpdate Update a locale
Update an existing locale.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param projectId Project ID
  - @param id Locale ID or locale name
  - @param localeUpdateParameters
  - @param optional nil or *LocaleUpdateOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)

@return LocaleDetails
*/
func (a *LocalesApiService) LocaleUpdate(ctx _context.Context, projectId string, id string, localeUpdateParameters LocaleUpdateParameters, localVarOptionals *LocaleUpdateOpts) (LocaleDetails, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  LocaleDetails
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/projects/{project_id}/locales/{id}"
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
	localVarPostBody = &localeUpdateParameters
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

// LocalesListOpts Optional parameters for the method 'LocalesList'
type LocalesListOpts struct {
	XPhraseAppOTP optional.String `json:"X-PhraseApp-OTP,omitempty"`
	Page          optional.Int32  `json:"page,omitempty"`
	PerPage       optional.Int32  `json:"per_page,omitempty"`
	SortBy        optional.String `json:"sort_by,omitempty"`
	Branch        optional.String `json:"branch,omitempty"`
}

/*
LocalesList List locales
List all locales for the given project.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param projectId Project ID
  - @param optional nil or *LocalesListOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)
  - @param "Page" (optional.Int32) -  Page number
  - @param "PerPage" (optional.Int32) -  Limit on the number of objects to be returned, between 1 and 100. 25 by default
  - @param "SortBy" (optional.String) -  Sort locales. Valid options are \"name_asc\", \"name_desc\", \"default_asc\", \"default_desc\".
  - @param "Branch" (optional.String) -  specify the branch to use

@return []Locale
*/
func (a *LocalesApiService) LocalesList(ctx _context.Context, projectId string, localVarOptionals *LocalesListOpts) ([]Locale, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Locale
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/projects/{project_id}/locales"
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
	if localVarOptionals != nil && localVarOptionals.SortBy.IsSet() {
		localVarQueryParams.Add("sort_by", parameterToString(localVarOptionals.SortBy.Value(), ""))
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
