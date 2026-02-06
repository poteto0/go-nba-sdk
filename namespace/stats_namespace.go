package namespace

import (
	"github.com/poteto0/go-nba-sdk/api"
	"github.com/poteto0/go-nba-sdk/api/stats"
	"github.com/poteto0/go-nba-sdk/types"
)

type IStatsNamespace interface {
	GetPlayerCareerStats(params types.PlayerCareerStatsParams) types.Response[types.PlayerCareerStatsResponseContent]
}

type StatsNamespace struct {
	provider api.IProvider
}

func NewStatsNamespace(provider api.IProvider) *StatsNamespace {
	return &StatsNamespace{
		provider: provider,
	}
}

func (s *StatsNamespace) GetPlayerCareerStats(params types.PlayerCareerStatsParams) types.Response[types.PlayerCareerStatsResponseContent] {
	return stats.GetPlayerCareerStats(s.provider, params)
}
