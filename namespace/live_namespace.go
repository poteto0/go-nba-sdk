package namespace

import (
	"github.com/poteto0/go-nba-sdk/api"
	"github.com/poteto0/go-nba-sdk/api/live"
	"github.com/poteto0/go-nba-sdk/types"
)

type ILiveNamespace interface {
	GetScoreBoard(params *types.ScoreBoardParams) types.Response[types.LiveScoreBoardResponse]
}

type LiveNamespace struct {
	provider api.IProvider
}

func NewLiveNamespace(provider api.IProvider) ILiveNamespace {
	return &LiveNamespace{
		provider: provider,
	}
}

func (l *LiveNamespace) GetScoreBoard(params *types.ScoreBoardParams) types.Response[types.LiveScoreBoardResponse] {
	return live.GetScoreBoard(l.provider, params)
}
