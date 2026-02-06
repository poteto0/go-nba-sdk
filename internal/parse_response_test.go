package internal

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/poteto0/go-nba-sdk/types"
	"github.com/stretchr/testify/assert"
)

func Test_ParseResponse(t *testing.T) {
	t.Run("successfully parse raw response", func(t *testing.T) {
		// Arrange
		expected := types.RawResponse{
			Resource: "test",
			Parameters: map[string]any{
				"param1": "value1",
			},
			ResultSets: []types.ResultSet{
				{
					Name:    "set1",
					Headers: []string{"h1", "h2"},
					RowSet:  [][]any{{"r1v1", "r1v2"}},
				},
			},
		}
		bodyBytes, _ := json.Marshal(expected)
		resp := &http.Response{
			Body:   io.NopCloser(bytes.NewReader(bodyBytes)),
			Header: make(http.Header),
		}

		// Act
		result, err := ParseResponse(resp)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expected.Resource, result.Resource)
		assert.Equal(t, expected.Parameters["param1"], result.Parameters["param1"])
		assert.Equal(t, expected.ResultSets[0].Name, result.ResultSets[0].Name)
	})
}

func Test_ParseResponseTo(t *testing.T) {
	type testTarget struct {
		Foo string `json:"foo"`
		Bar int    `json:"bar"`
	}

	t.Run("successfully parse to target struct", func(t *testing.T) {
		// Arrange
		expected := testTarget{Foo: "hello", Bar: 123}
		bodyBytes, _ := json.Marshal(expected)
		resp := &http.Response{
			Body:   io.NopCloser(bytes.NewReader(bodyBytes)),
			Header: make(http.Header),
		}

		// Act
		var result testTarget
		err := ParseResponseTo(resp, &result)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("successfully parse gzipped response", func(t *testing.T) {
		// Arrange
		expected := testTarget{Foo: "zipped", Bar: 456}
		bodyBytes, _ := json.Marshal(expected)

		var buf bytes.Buffer
		gz := gzip.NewWriter(&buf)
		_, _ = gz.Write(bodyBytes)
		_ = gz.Close()

		resp := &http.Response{
			Body: io.NopCloser(&buf),
			Header: http.Header{
				"Content-Encoding": []string{"gzip"},
			},
		}

		// Act
		var result testTarget
		err := ParseResponseTo(resp, &result)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("error when invalid json", func(t *testing.T) {
		// Arrange
		resp := &http.Response{
			Body:   io.NopCloser(bytes.NewReader([]byte("invalid json"))),
			Header: make(http.Header),
		}

		// Act
		var result testTarget
		err := ParseResponseTo(resp, &result)

		// Assert
		assert.Error(t, err)
	})

	t.Run("error when invalid gzip", func(t *testing.T) {
		// Arrange
		resp := &http.Response{
			Body: io.NopCloser(bytes.NewReader([]byte("not a gzip"))),
			Header: http.Header{
				"Content-Encoding": []string{"gzip"},
			},
		}

		// Act
		var result testTarget
		err := ParseResponseTo(resp, &result)

		// Assert
		assert.Error(t, err)
	})
}
