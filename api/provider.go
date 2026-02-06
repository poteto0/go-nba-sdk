package api

import (
	"net/http"

	"github.com/poteto0/go-nba-sdk/internal"
)

type IProvider interface {
	Get(path string) (*http.Response, error)
}

type Provider struct {
	client internal.IHttpClient
}

func NewProvider() IProvider {
	return &Provider{
		client: internal.NewHttpClient(),
	}
}

func (p *Provider) Get(path string) (*http.Response, error) {
	return p.client.Get(path, nil)
}
