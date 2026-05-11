package namespace

import (
	"github.com/poteto0/go-nba-sdk/api"
	"github.com/poteto0/go-nba-sdk/api/draft"
	"github.com/poteto0/go-nba-sdk/types"
)

type IDraftNamespace interface {
	GetDraftBoard(params *types.DraftBoardParams) types.Response[types.DraftBoardResponse]
}

type DraftNamespace struct {
	provider api.IProvider
}

func NewDraftNamespace(provider api.IProvider) IDraftNamespace {
	return &DraftNamespace{
		provider: provider,
	}
}

func (d *DraftNamespace) GetDraftBoard(params *types.DraftBoardParams) types.Response[types.DraftBoardResponse] {
	return draft.GetDraftBoard(d.provider, params)
}
