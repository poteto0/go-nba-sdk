package internal

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"

	"github.com/poteto0/go-nba-sdk/types"
)

// 1. check response header for gzip
// 2. if gzip, decompress
// 3. read body
// 4. return body as string in Response struct
//
// Caution: remember to close resp.Body in the caller function
func ParseResponse(resp *http.Response) (types.RawResponse, error) {
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return types.RawResponse{}, err
		}
		reader = gzipReader
		defer reader.Close()
	default:
		reader = resp.Body
	}

	bodyBytes, err := io.ReadAll(reader)
	if err != nil {
		return types.RawResponse{}, err
	}

	var rawResp types.RawResponse
	if err := json.Unmarshal(bodyBytes, &rawResp); err != nil {
		return types.RawResponse{}, err
	}

	return rawResp, nil
}
