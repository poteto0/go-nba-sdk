package internal

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"

	"github.com/poteto0/go-nba-sdk/types"
)

// ParseResponse parses the HTTP response into types.RawResponse.
// This is primarily used for the Stats API.
// Caution: remember to close resp.Body in the caller function.
func ParseResponse(resp *http.Response) (types.RawResponse, error) {
	var rawResp types.RawResponse
	err := ParseResponseTo(resp, &rawResp)
	return rawResp, err
}

// ParseResponseTo parses the HTTP response into the provided target.
// It handles gzip decompression if necessary.
// Caution: remember to close resp.Body in the caller function.
func ParseResponseTo(resp *http.Response, target any) error {
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return err
		}
		reader = gzipReader
		defer reader.Close()
	default:
		reader = resp.Body
	}

	bodyBytes, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bodyBytes, target); err != nil {
		return err
	}

	return nil
}
