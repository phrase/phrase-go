/*
Phrase Strings API Reference

Testing LocalesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package phrase

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/antihax/optional"
	"github.com/phrase/phrase-go/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_phrase_LocalesApiService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Send the mock response
		response := `{"id":"1","name":"English","code":"DE","default":true,"main":true,"rtl":true,"plural_forms":["plural_forms"]}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(response))
	}))

	defer server.Close()

	configuration := phrase.NewConfiguration()
	configuration.BasePath = server.URL
	apiClient := phrase.NewAPIClient(configuration)

	t.Run("Test LocalesApiService LocaleDownload", func(t *testing.T) {
		formatOptions := optional.NewInterface(map[string]interface{}{
			"foo": "bar",
			"baz": "bazz",
		})
		localeDownloadOpts := phrase.LocaleDownloadOpts{FormatOptions: formatOptions}

		resp, httpRes, err := apiClient.LocalesApi.LocaleDownload(context.Background(), "project_id_example", "locale_id", &localeDownloadOpts)
		requestUrl := httpRes.Request.URL

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "format_options%5Bbaz%5D=bazz&format_options%5Bfoo%5D=bar", requestUrl.RawQuery)
		assert.Equal(t, "/projects/project_id_example/locales/locale_id/download", requestUrl.Path)
		assert.Equal(t, "GET", httpRes.Request.Method)
	})

	t.Run("Test LocalesApiService LocaleShow", func(t *testing.T) {
		localeShowOpts := phrase.LocaleShowOpts{}
		resp, httpRes, err := apiClient.LocalesApi.LocaleShow(context.Background(), "project_id_example", "locale_id", &localeShowOpts)
		requestUrl := httpRes.Request.URL

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "1", resp.Id)
		assert.Equal(t, "DE", resp.Code)
		assert.Equal(t, "English", resp.Name)
		assert.Equal(t, "/projects/project_id_example/locales/locale_id", requestUrl.Path)
		assert.Equal(t, "GET", httpRes.Request.Method)
	})

	t.Run("Test LocalesApiService LocaleUpdate", func(t *testing.T) {
		localeUpdateOpts := phrase.LocaleUpdateOpts{}
		localeUpdateParameters := phrase.LocaleUpdateParameters{}
		resp, httpRes, err := apiClient.LocalesApi.LocaleUpdate(context.Background(), "project_id_example", "locale_id", localeUpdateParameters, &localeUpdateOpts)
		requestUrl := httpRes.Request.URL

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "1", resp.Id)
		assert.Equal(t, "DE", resp.Code)
		assert.Equal(t, "English", resp.Name)
		assert.Equal(t, "/projects/project_id_example/locales/locale_id", requestUrl.Path)
		assert.Equal(t, "PATCH", httpRes.Request.Method)
	})

	t.Run("Test LocalesApiService LocalesList", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Send the mock response
			response := `[{"id":"1","name":"English","code":"DE","default":true,"main":true,"rtl":true,"plural_forms":["plural_forms"]}]`
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			w.Write([]byte(response))
		}))

		defer server.Close()

		configuration := phrase.NewConfiguration()
		configuration.BasePath = server.URL
		apiClient := phrase.NewAPIClient(configuration)

		localVarOptionals := phrase.LocalesListOpts{}
		resp, httpRes, err := apiClient.LocalesApi.LocalesList(context.Background(), "project_id_example", &localVarOptionals)
		requestUrl := httpRes.Request.URL
		locale := resp[0]

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, 1, len(resp))
		assert.Equal(t, "1", locale.Id)
		assert.Equal(t, "DE", locale.Code)
		assert.Equal(t, "English", locale.Name)
		assert.Equal(t, "/projects/project_id_example/locales", requestUrl.Path)
		assert.Equal(t, "GET", httpRes.Request.Method)
	})

}
