package api

import (
	"net/http"

	"github.com/poteto0/go-nba-sdk/internal"
	"github.com/poteto0/go-nba-sdk/types"
)

type IProvider interface {
	Get(path string, headers *http.Header) (*http.Response, error)
}

type Provider struct {
	client internal.IHttpClient
}

func NewProvider(config *types.GnsConfig) IProvider {
	return &Provider{
		client: internal.NewHttpClient(
			config,
		),
	}
}

func (p *Provider) Get(path string, headers *http.Header) (*http.Response, error) {
	return p.client.Get(path, headers)
}
