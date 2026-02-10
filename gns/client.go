package gns

import (
	"github.com/poteto0/go-nba-sdk/api"
	"github.com/poteto0/go-nba-sdk/namespace"
	"github.com/poteto0/go-nba-sdk/types"
)

type Client struct {
	Stats namespace.IStatsNamespace
	Live  namespace.ILiveNamespace
}

func NewClient(config *types.GnsConfig) *Client {
	provider := api.NewProvider(
		config,
	)

	return &Client{
		Stats: namespace.NewStatsNamespace(provider),
		Live:  namespace.NewLiveNamespace(provider),
	}
}
