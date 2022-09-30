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

// UploadsApiService UploadsApi service
type UploadsApiService service

// UploadCreateOpts Optional parameters for the method 'UploadCreate'
type UploadCreateOpts struct {
	XPhraseAppOTP      optional.String    `json:"X-PhraseApp-OTP,omitempty"`
	Branch             optional.String    `json:"branch,omitempty"`
	File               optional.Interface `json:"file,omitempty"`
	FileFormat         optional.String    `json:"file_format,omitempty"`
	LocaleId           optional.String    `json:"locale_id,omitempty"`
	Tags               optional.String    `json:"tags,omitempty"`
	UpdateTranslations optional.Bool      `json:"update_translations,omitempty"`
	UpdateDescriptions optional.Bool      `json:"update_descriptions,omitempty"`
	ConvertEmoji       optional.Bool      `json:"convert_emoji,omitempty"`
	SkipUploadTags     optional.Bool      `json:"skip_upload_tags,omitempty"`
	SkipUnverification optional.Bool      `json:"skip_unverification,omitempty"`
	FileEncoding       optional.String    `json:"file_encoding,omitempty"`
	LocaleMapping      optional.Interface `json:"locale_mapping,omitempty"`
	FormatOptions      optional.Interface `json:"format_options,omitempty"`
	Autotranslate      optional.Bool      `json:"autotranslate,omitempty"`
	MarkReviewed       optional.Bool      `json:"mark_reviewed,omitempty"`
}

/*
UploadCreate Upload a new file
Upload a new language file. Creates necessary resources in your project.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param projectId Project ID
  - @param optional nil or *UploadCreateOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)
  - @param "Branch" (optional.String) -  specify the branch to use
  - @param "File" (optional.Interface of *os.File) -  File to be imported
  - @param "FileFormat" (optional.String) -  File format. Auto-detected when possible and not specified.
  - @param "LocaleId" (optional.String) -  Locale of the file's content. Can be the name or public id of the locale. Preferred is the public id.
  - @param "Tags" (optional.String) -  List of tags separated by comma to be associated with the new keys contained in the upload.
  - @param "UpdateTranslations" (optional.Bool) -  Indicates whether existing translations should be updated with the file content.
  - @param "UpdateDescriptions" (optional.Bool) -  Existing key descriptions will be updated with the file content. Empty descriptions overwrite existing descriptions.
  - @param "ConvertEmoji" (optional.Bool) -  This option is obsolete. Providing the option will cause a bad request error.
  - @param "SkipUploadTags" (optional.Bool) -  Indicates whether the upload should not create upload tags.
  - @param "SkipUnverification" (optional.Bool) -  Indicates whether the upload should unverify updated translations.
  - @param "FileEncoding" (optional.String) -  Enforces a specific encoding on the file contents. Valid options are \\\"UTF-8\\\", \\\"UTF-16\\\" and \\\"ISO-8859-1\\\".
  - @param "LocaleMapping" (optional.Interface of map[string]interface{}) -  Optional, format specific mapping between locale names and the columns the translations to those locales are contained in.
  - @param "FormatOptions" (optional.Interface of map[string]interface{}) -  Additional options available for specific formats. See our format guide for complete list.
  - @param "Autotranslate" (optional.Bool) -  If set, translations for the uploaded language will be fetched automatically.
  - @param "MarkReviewed" (optional.Bool) -  Indicated whether the imported translations should be marked as reviewed. This setting is available if the review workflow (currently beta) is enabled for the project.

@return Upload
*/
func (a *UploadsApiService) UploadCreate(ctx _context.Context, projectId string, localVarOptionals *UploadCreateOpts) (Upload, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Upload
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/projects/{project_id}/uploads"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", _neturl.QueryEscape(parameterToString(projectId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	if localVarOptionals != nil && localVarOptionals.Branch.IsSet() {
		localVarFormParams.Add("branch", parameterToString(localVarOptionals.Branch.Value(), ""))
	}
	localVarFormFileName = "file"
	var localVarFile *os.File
	if localVarOptionals != nil && localVarOptionals.File.IsSet() {
		localVarFileOk := false
		localVarFile, localVarFileOk = localVarOptionals.File.Value().(*os.File)
		if !localVarFileOk {
			return localVarReturnValue, nil, reportError("file should be *os.File")
		}
	}
	if localVarFile != nil {
		fbs, _ := _ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	if localVarOptionals != nil && localVarOptionals.FileFormat.IsSet() {
		localVarFormParams.Add("file_format", parameterToString(localVarOptionals.FileFormat.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LocaleId.IsSet() {
		localVarFormParams.Add("locale_id", parameterToString(localVarOptionals.LocaleId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Tags.IsSet() {
		localVarFormParams.Add("tags", parameterToString(localVarOptionals.Tags.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UpdateTranslations.IsSet() {
		localVarFormParams.Add("update_translations", parameterToString(localVarOptionals.UpdateTranslations.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UpdateDescriptions.IsSet() {
		localVarFormParams.Add("update_descriptions", parameterToString(localVarOptionals.UpdateDescriptions.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ConvertEmoji.IsSet() {
		localVarFormParams.Add("convert_emoji", parameterToString(localVarOptionals.ConvertEmoji.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SkipUploadTags.IsSet() {
		localVarFormParams.Add("skip_upload_tags", parameterToString(localVarOptionals.SkipUploadTags.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SkipUnverification.IsSet() {
		localVarFormParams.Add("skip_unverification", parameterToString(localVarOptionals.SkipUnverification.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileEncoding.IsSet() {
		localVarFormParams.Add("file_encoding", parameterToString(localVarOptionals.FileEncoding.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LocaleMapping.IsSet() {
		for key, value := range localVarOptionals.LocaleMapping.Value().(map[string]interface{}) {
			localVarFormParams.Add(fmt.Sprintf("locale_mapping[%s]", key), parameterToString(value, ""))
		}
	}
	if localVarOptionals != nil && localVarOptionals.FormatOptions.IsSet() {
		for key, value := range localVarOptionals.FormatOptions.Value().(map[string]interface{}) {
			localVarFormParams.Add(fmt.Sprintf("format_options[%s]", key), parameterToString(value, ""))
		}
	}
	if localVarOptionals != nil && localVarOptionals.Autotranslate.IsSet() {
		localVarFormParams.Add("autotranslate", parameterToString(localVarOptionals.Autotranslate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MarkReviewed.IsSet() {
		localVarFormParams.Add("mark_reviewed", parameterToString(localVarOptionals.MarkReviewed.Value(), ""))
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

// UploadShowOpts Optional parameters for the method 'UploadShow'
type UploadShowOpts struct {
	XPhraseAppOTP optional.String `json:"X-PhraseApp-OTP,omitempty"`
	Branch        optional.String `json:"branch,omitempty"`
}

/*
UploadShow View upload details
View details and summary for a single upload.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param projectId Project ID
  - @param id ID
  - @param optional nil or *UploadShowOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)
  - @param "Branch" (optional.String) -  specify the branch to use

@return Upload
*/
func (a *UploadsApiService) UploadShow(ctx _context.Context, projectId string, id string, localVarOptionals *UploadShowOpts) (Upload, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Upload
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/projects/{project_id}/uploads/{id}"
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

// UploadsListOpts Optional parameters for the method 'UploadsList'
type UploadsListOpts struct {
	XPhraseAppOTP optional.String `json:"X-PhraseApp-OTP,omitempty"`
	Page          optional.Int32  `json:"page,omitempty"`
	PerPage       optional.Int32  `json:"per_page,omitempty"`
	Branch        optional.String `json:"branch,omitempty"`
}

/*
UploadsList List uploads
List all uploads for the given project.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param projectId Project ID
  - @param optional nil or *UploadsListOpts - Optional Parameters:
  - @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)
  - @param "Page" (optional.Int32) -  Page number
  - @param "PerPage" (optional.Int32) -  allows you to specify a page size up to 100 items, 25 by default
  - @param "Branch" (optional.String) -  specify the branch to use

@return []Upload
*/
func (a *UploadsApiService) UploadsList(ctx _context.Context, projectId string, localVarOptionals *UploadsListOpts) ([]Upload, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Upload
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/projects/{project_id}/uploads"
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
