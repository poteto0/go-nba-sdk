package gns

import (
	"github.com/poteto0/go-nba-sdk/api"
	"github.com/poteto0/go-nba-sdk/namespace"
)

type IClient interface{}

type Client struct {
	Stats namespace.IStatsNamespace
}

func NewClient() *Client {
	provider := api.NewProvider()
	return &Client{
		Stats: namespace.NewStatsNamespace(provider),
	}
}
