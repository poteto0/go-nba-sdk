package stats

import (
	"github.com/google/go-querystring/query"
	"github.com/poteto0/go-nba-sdk/api"
	"github.com/poteto0/go-nba-sdk/constants"
	"github.com/poteto0/go-nba-sdk/internal"
	"github.com/poteto0/go-nba-sdk/types"
)

// GetPlayByPlay fetches play‑by‑play data using the Stats API (playbyplayv3 endpoint).
func GetPlayByPlay(provider api.IProvider, params *types.PlayByPlayParams) types.Response[types.LivePlayByPlayResponse] {
	if params == nil {
		return types.Response[types.LivePlayByPlayResponse]{Error: types.NewGnsError("playbyplay params is required")}
	}
	if params.GameID == "" {
		return types.Response[types.LivePlayByPlayResponse]{Error: types.NewGnsError("game id is required")}
	}
	if params.StartPeriod == 0 {
		params.StartPeriod = 1
	}
	if params.EndPeriod == 0 {
		params.EndPeriod = 4
	}

	path := constants.StatsBaseUrl + constants.PlayByPlayV3Path
	v, err := query.Values(params)
	if err != nil {
		return types.Response[types.LivePlayByPlayResponse]{Error: err}
	}
	path = path + "?" + v.Encode()
	resp, err := provider.Get(path, &constants.DefaultStatsHeaders)
	if err != nil {
		return types.Response[types.LivePlayByPlayResponse]{Error: err}
	}
	defer resp.Body.Close()

	var data types.LivePlayByPlayResponse
	if err := internal.ParseResponseTo(resp, &data); err != nil {
		return types.Response[types.LivePlayByPlayResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
		}
	}

	return types.Response[types.LivePlayByPlayResponse]{
		Contents:   data,
		StatusCode: resp.StatusCode,
		Error:      nil,
	}
}
