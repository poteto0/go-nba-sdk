package gns

import (
	"github.com/poteto0/go-nba-sdk/api"
	"github.com/poteto0/go-nba-sdk/namespace"
)

type Client struct {
	Stats namespace.IStatsNamespace
	Live  namespace.ILiveNamespace
}

func NewClient() *Client {
	provider := api.NewProvider()
	return &Client{
		Stats: namespace.NewStatsNamespace(provider),
		Live:  namespace.NewLiveNamespace(provider),
	}
}
