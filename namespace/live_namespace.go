package namespace

import (
	"github.com/poteto0/go-nba-sdk/api"
	"github.com/poteto0/go-nba-sdk/api/live"
	"github.com/poteto0/go-nba-sdk/types"
)

type ILiveNamespace interface {
	GetScoreBoard(params *types.ScoreBoardParams) types.Response[types.LiveScoreBoardResponse]
	GetBoxScore(params *types.BoxScoreParams) types.Response[types.LiveBoxScoreResponse]
	GetPlayByPlay(params *types.PlayByPlayParams) types.Response[types.LivePlayByPlayResponse]
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

func (l *LiveNamespace) GetBoxScore(params *types.BoxScoreParams) types.Response[types.LiveBoxScoreResponse] {
	return live.GetBoxScore(l.provider, params)
}

func (l *LiveNamespace) GetPlayByPlay(params *types.PlayByPlayParams) types.Response[types.LivePlayByPlayResponse] {
	return live.GetPlayByPlay(l.provider, params)
}
