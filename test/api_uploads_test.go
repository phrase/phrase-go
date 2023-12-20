/*
Phrase Strings API Reference

Testing UploadsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package phrase

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/antihax/optional"
	"github.com/phrase/phrase-go/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_phrase_UploadsApiService(t *testing.T) {
	t.Run("Test UploadsApiService UploadCreate", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// assert request body for proper serialization of tested format_options
			r.ParseMultipartForm(0)
			assert.Equal(t, r.FormValue("format_options[translation_columns][en]"), "f")
			assert.Equal(t, r.FormValue("format_options[translation_columns][de]"), "e")
			assert.Equal(t, r.FormValue("format_options[top_level_key]"), "100")

			// Send the mock response
			response := `{"id": "1", "filename": "test.json", "format": "json", "state": "valid" }`
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			w.Write([]byte(response))
		}))

		defer server.Close()

		configuration := phrase.NewConfiguration()
		configuration.BasePath = server.URL
		apiClient := phrase.NewAPIClient(configuration)

		file, _ := os.Create("testfile.json")
		fileFormat := optional.NewString("json")
		fileObject := optional.NewInterface(file)
		localeId := optional.NewString("99")

		// setting format_options
		formatOptions := make(map[string]interface{})
		nestedOptions := make(map[string]interface{})

		nestedOptions["en"] = "f"
		nestedOptions["de"] = "e"

		formatOptions["top_level_key"] = 100
		formatOptions["translation_columns"] = nestedOptions

		formatOptionsMap := optional.NewInterface(formatOptions)

		localVarOptionals := phrase.UploadCreateOpts{FileFormat: fileFormat, File: fileObject, LocaleId: localeId, FormatOptions: formatOptionsMap}
		resp, httpRes, err := apiClient.UploadsApi.UploadCreate(context.Background(), "project_id", &localVarOptionals)
		requestUrl := httpRes.Request.URL

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "1", resp.Id)
		assert.Equal(t, "test.json", resp.Filename)
		assert.Equal(t, "json", resp.Format)
		assert.Equal(t, "/projects/project_id/uploads", requestUrl.Path)
		assert.Equal(t, "POST", httpRes.Request.Method)

		defer os.Remove("testfile.json")
	})

	t.Run("Test UploadsApiService UploadShow", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Send the mock response
			response := `{"id": "1", "filename": "test.json", "format": "json", "state": "valid" }`
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			w.Write([]byte(response))
		}))

		defer server.Close()

		configuration := phrase.NewConfiguration()
		configuration.BasePath = server.URL
		apiClient := phrase.NewAPIClient(configuration)

		localVarOptionals := phrase.UploadShowOpts{}
		resp, httpRes, err := apiClient.UploadsApi.UploadShow(context.Background(), "project_id", "upload_id", &localVarOptionals)
		requestUrl := httpRes.Request.URL

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "1", resp.Id)
		assert.Equal(t, "test.json", resp.Filename)
		assert.Equal(t, "json", resp.Format)
		assert.Equal(t, "/projects/project_id/uploads/upload_id", requestUrl.Path)
		assert.Equal(t, "GET", httpRes.Request.Method)
	})

}
