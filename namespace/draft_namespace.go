package namespace

import (
	"github.com/poteto0/go-nba-sdk/api"
	"github.com/poteto0/go-nba-sdk/api/draft"
	"github.com/poteto0/go-nba-sdk/types"
)

type IDraftNamespace interface {
	GetCombineStats(params *types.DraftCombineStatsParams) types.Response[types.DraftCombineStatsResponseContent]
}

type DraftNamespace struct {
	provider api.IProvider
}

func NewDraftNamespace(provider api.IProvider) IDraftNamespace {
	return &DraftNamespace{
		provider: provider,
	}
}

func (d *DraftNamespace) GetCombineStats(params *types.DraftCombineStatsParams) types.Response[types.DraftCombineStatsResponseContent] {
	return draft.GetCombineStats(d.provider, params)
}
