package gns

import (
	"github.com/poteto0/go-nba-sdk/api"
	"github.com/poteto0/go-nba-sdk/api/stats"
)

type IClient interface{}

type Client struct {
	Stats stats.IStatsNamespace
}

func NewClient() *Client {
	provider := api.NewProvider()
	return &Client{
		Stats: stats.NewStatsNamespace(provider),
	}
}
