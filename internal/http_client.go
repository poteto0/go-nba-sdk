package internal

import (
	"net/http"
	"net/url"
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

	httpClient := &http.Client{
		Timeout: config.Timeout,
	}

	if config.ProxyUrl != "" {
		proxyUrl, err := url.Parse(config.ProxyUrl)
		if err == nil {
			httpClient.Transport = &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			}
		}
	}

	return &HttpClient{
		client: httpClient,
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
