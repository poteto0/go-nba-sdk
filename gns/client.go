package gns

import (
	"net/http"

	"github.com/poteto0/go-nba-sdk/internal"
)

type IClient interface {
	Get(path string) (*http.Response, error)
}

type Client struct {
	client internal.IHttpClient
}

func NewStatsClient() IClient {
	return &Client{
		client: internal.NewHttpClient(),
	}
}

func (c *Client) Get(path string) (*http.Response, error) {
	return c.client.Get(path, nil)
}
