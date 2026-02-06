package internal

import (
	"net/http"
	"time"

	"github.com/poteto0/go-nba-sdk/constants"
)

type HttpClient struct {
	client        *http.Client
	defaultHeader http.Header
}

type IHttpClient interface {
	Get(path string, headers *http.Header) (*http.Response, error)
}

func NewHttpClient() IHttpClient {
	return &HttpClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *HttpClient) Get(path string, headers *http.Header) (*http.Response, error) {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		req.Header = *headers
	} else {
		req.Header = constants.DefaultStatsHeaders
	}

	return c.client.Do(req)
}
