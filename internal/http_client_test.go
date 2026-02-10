package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateHttpClient(t *testing.T) {
	// Act
	client := NewHttpClient(nil)

	// Assert
	assert.NotNil(t, client)
}

func Test_Get(t *testing.T) {
	// Arrange & Mock
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/test", r.URL.Path)

		if r.Header.Get("Custom-Header") == "Value" {
			// Custom header case
			assert.Equal(t, "Value", r.Header.Get("Custom-Header"))
			// If headers are provided, default headers should NOT be used (based on "headersが指定されている場合には...指定されていない場合にはメンバのheaderを使って")
			assert.Empty(t, r.Header.Get("Accept-Language"))
		} else if r.Header.Get("Custom-Header") == "Error" {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			// Default header case
			assert.Equal(t, "en-US,en;q=0.5", r.Header.Get("Accept-Language"))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Arrange
	client := NewHttpClient(nil)

	t.Run("Get with default headers", func(t *testing.T) {
		// Act
		resp, err := client.Get(server.URL+"/test", nil)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Get with custom headers", func(t *testing.T) {
		// Arrange
		customHeader := http.Header{
			"Custom-Header": {"Value"},
		}

		// Act
		resp, err := client.Get(server.URL+"/test", &customHeader)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
