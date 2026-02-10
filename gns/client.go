package gns

import (
	"time"

	"github.com/poteto0/go-nba-sdk/api"
	"github.com/poteto0/go-nba-sdk/namespace"
	"github.com/poteto0/go-nba-sdk/types"
)

type Client struct {
	Stats namespace.IStatsNamespace
	Live  namespace.ILiveNamespace
}

var DefaultConfig = &types.GnsConfig{
	Timeout: 10 * time.Second,
}

func NewClient(config *types.GnsConfig) *Client {
	if config == nil {
		config = DefaultConfig
	}

	if config.Timeout == 0 {
		config.Timeout = DefaultConfig.Timeout
	}

	provider := api.NewProvider(
		config,
	)

	return &Client{
		Stats: namespace.NewStatsNamespace(provider),
		Live:  namespace.NewLiveNamespace(provider),
	}
}
