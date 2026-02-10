package internal

import (
	"net/http"
	"time"

	"github.com/poteto0/go-nba-sdk/constants"
	"github.com/poteto0/go-nba-sdk/types"
)

type HttpClient struct {
	client *http.Client
}

var DefaultConfig = &types.GnsConfig{
	Timeout: 10 * time.Second,
}

type IHttpClient interface {
	Get(path string, headers *http.Header) (*http.Response, error)
}

func NewHttpClient(config *types.GnsConfig) IHttpClient {
	if config == nil {
		config = DefaultConfig
	}

	if config.Timeout == 0 {
		config.Timeout = DefaultConfig.Timeout
	}

	return &HttpClient{
		client: &http.Client{
			Timeout: config.Timeout,
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
