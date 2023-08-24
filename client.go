package phrase

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/oauth2"
)

var (
	jsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	xmlCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
)

// APIClient manages communication with the Phrase Strings API Reference API v2.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	AccountsApi *AccountsApiService

	AuthorizationsApi *AuthorizationsApiService

	BitbucketSyncApi *BitbucketSyncApiService

	BlacklistedKeysApi *BlacklistedKeysApiService

	BranchesApi *BranchesApiService

	CommentReactionsApi *CommentReactionsApiService

	CommentRepliesApi *CommentRepliesApiService

	CommentsApi *CommentsApiService

	DistributionsApi *DistributionsApiService

	DocumentsApi *DocumentsApiService

	FormatsApi *FormatsApiService

	GitHubSyncApi *GitHubSyncApiService

	GitLabSyncApi *GitLabSyncApiService

	GlossariesApi *GlossariesApiService

	GlossaryTermTranslationsApi *GlossaryTermTranslationsApiService

	GlossaryTermsApi *GlossaryTermsApiService

	ICUApi *ICUApiService

	InvitationsApi *InvitationsApiService

	JobCommentsApi *JobCommentsApiService

	JobLocalesApi *JobLocalesApiService

	JobTemplateLocalesApi *JobTemplateLocalesApiService

	JobTemplatesApi *JobTemplatesApiService

	JobsApi *JobsApiService

	KeysApi *KeysApiService

	LocalesApi *LocalesApiService

	MembersApi *MembersApiService

	NotificationGroupsApi *NotificationGroupsApiService

	NotificationsApi *NotificationsApiService

	OrdersApi *OrdersApiService

	OrganizationJobTemplateLocalesApi *OrganizationJobTemplateLocalesApiService

	OrganizationJobTemplatesApi *OrganizationJobTemplatesApiService

	ProjectsApi *ProjectsApiService

	ReleasesApi *ReleasesApiService

	ScreenshotMarkersApi *ScreenshotMarkersApiService

	ScreenshotsApi *ScreenshotsApiService

	SearchApi *SearchApiService

	SpacesApi *SpacesApiService

	StyleGuidesApi *StyleGuidesApiService

	TagsApi *TagsApiService

	TeamsApi *TeamsApiService

	TranslationsApi *TranslationsApiService

	UploadsApi *UploadsApiService

	UsersApi *UsersApiService

	VariablesApi *VariablesApiService

	VersionsHistoryApi *VersionsHistoryApiService

	WebhookDeliveriesApi *WebhookDeliveriesApiService

	WebhooksApi *WebhooksApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.AccountsApi = (*AccountsApiService)(&c.common)
	c.AuthorizationsApi = (*AuthorizationsApiService)(&c.common)
	c.BitbucketSyncApi = (*BitbucketSyncApiService)(&c.common)
	c.BlacklistedKeysApi = (*BlacklistedKeysApiService)(&c.common)
	c.BranchesApi = (*BranchesApiService)(&c.common)
	c.CommentReactionsApi = (*CommentReactionsApiService)(&c.common)
	c.CommentRepliesApi = (*CommentRepliesApiService)(&c.common)
	c.CommentsApi = (*CommentsApiService)(&c.common)
	c.DistributionsApi = (*DistributionsApiService)(&c.common)
	c.DocumentsApi = (*DocumentsApiService)(&c.common)
	c.FormatsApi = (*FormatsApiService)(&c.common)
	c.GitHubSyncApi = (*GitHubSyncApiService)(&c.common)
	c.GitLabSyncApi = (*GitLabSyncApiService)(&c.common)
	c.GlossariesApi = (*GlossariesApiService)(&c.common)
	c.GlossaryTermTranslationsApi = (*GlossaryTermTranslationsApiService)(&c.common)
	c.GlossaryTermsApi = (*GlossaryTermsApiService)(&c.common)
	c.ICUApi = (*ICUApiService)(&c.common)
	c.InvitationsApi = (*InvitationsApiService)(&c.common)
	c.JobCommentsApi = (*JobCommentsApiService)(&c.common)
	c.JobLocalesApi = (*JobLocalesApiService)(&c.common)
	c.JobTemplateLocalesApi = (*JobTemplateLocalesApiService)(&c.common)
	c.JobTemplatesApi = (*JobTemplatesApiService)(&c.common)
	c.JobsApi = (*JobsApiService)(&c.common)
	c.KeysApi = (*KeysApiService)(&c.common)
	c.LocalesApi = (*LocalesApiService)(&c.common)
	c.MembersApi = (*MembersApiService)(&c.common)
	c.NotificationGroupsApi = (*NotificationGroupsApiService)(&c.common)
	c.NotificationsApi = (*NotificationsApiService)(&c.common)
	c.OrdersApi = (*OrdersApiService)(&c.common)
	c.OrganizationJobTemplateLocalesApi = (*OrganizationJobTemplateLocalesApiService)(&c.common)
	c.OrganizationJobTemplatesApi = (*OrganizationJobTemplatesApiService)(&c.common)
	c.ProjectsApi = (*ProjectsApiService)(&c.common)
	c.ReleasesApi = (*ReleasesApiService)(&c.common)
	c.ScreenshotMarkersApi = (*ScreenshotMarkersApiService)(&c.common)
	c.ScreenshotsApi = (*ScreenshotsApiService)(&c.common)
	c.SearchApi = (*SearchApiService)(&c.common)
	c.SpacesApi = (*SpacesApiService)(&c.common)
	c.StyleGuidesApi = (*StyleGuidesApiService)(&c.common)
	c.TagsApi = (*TagsApiService)(&c.common)
	c.TeamsApi = (*TeamsApiService)(&c.common)
	c.TranslationsApi = (*TranslationsApiService)(&c.common)
	c.UploadsApi = (*UploadsApiService)(&c.common)
	c.UsersApi = (*UsersApiService)(&c.common)
	c.VariablesApi = (*VariablesApiService)(&c.common)
	c.VersionsHistoryApi = (*VersionsHistoryApiService)(&c.common)
	c.WebhookDeliveriesApi = (*WebhookDeliveriesApiService)(&c.common)
	c.WebhooksApi = (*WebhooksApiService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	} else if t, ok := obj.(time.Time); ok {
		return t.Format(time.RFC3339)
	}

	return fmt.Sprintf("%v", obj)
}

// helper for converting interface{} parameters to json strings
func parameterToJson(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*APIResponse, error) {
	if c.cfg.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	resp, err := c.cfg.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	response := NewAPIResponse(resp)

	if c.cfg.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return response, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	return response, err
}

// ChangeBasePath changes base path to allow switching to mocks
func (c *APIClient) ChangeBasePath(path string) {
	c.cfg.BasePath = path
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *Configuration {
	return c.cfg
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFileName string,
	fileName string,
	fileBytes []byte) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		if len(fileBytes) > 0 && fileName != "" {
			w.Boundary()
			//_, fileNm := filepath.Split(fileName)
			part, err := w.CreateFormFile(formFileName, filepath.Base(fileName))
			if err != nil {
				return nil, err
			}
			_, err = part.Write(fileBytes)
			if err != nil {
				return nil, err
			}
		}

		// Set the Boundary in the Content-Type
		headerParams["Content-Type"] = w.FormDataContentType()

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		url.Host = c.cfg.Host
	}

	// Override request scheme, if applicable
	if c.cfg.Scheme != "" {
		url.Scheme = c.cfg.Scheme
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// OAuth2 authentication
		if tok, ok := ctx.Value(ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			latestToken.SetAuthHeader(localVarRequest)
		}

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(ContextBasicAuth).(BasicAuth); ok {
			localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
		}

		// AccessToken Authentication
		if auth, ok := ctx.Value(ContextAccessToken).(string); ok {
			localVarRequest.Header.Add("Authorization", "Bearer "+auth)
		}

	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}

	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}

	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}

	if f, ok := v.(**os.File); ok {
		file, err := ioutil.TempFile(os.TempDir(), "phrase-locale-download-*")
		if err != nil {
			return err
		}
		file.Write(b)
		file.Seek(0, io.SeekStart)
		*f = file
		return nil
	}

	if xmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}

	if jsonCheck.MatchString(contentType) {
		if err = json.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}

	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		err = xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.model
}
