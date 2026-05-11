package draft

import (
	"github.com/google/go-querystring/query"
	"github.com/poteto0/go-nba-sdk/api"
	"github.com/poteto0/go-nba-sdk/constants"
	"github.com/poteto0/go-nba-sdk/internal"
	"github.com/poteto0/go-nba-sdk/parser"
	"github.com/poteto0/go-nba-sdk/types"
)

func GetCombineStats(provider api.IProvider, params *types.DraftCombineStatsParams) types.Response[types.DraftCombineStatsResponse] {
	if params == nil {
		params = &types.DraftCombineStatsParams{}
	}

	if params.LeagueID == "" {
		params.LeagueID = "00"
	}

	path := constants.StatsBaseUrl + constants.DraftCombineStatsPath
	v, err := query.Values(params)
	if err != nil {
		return types.Response[types.DraftCombineStatsResponse]{Error: err}
	}

	path = path + "?" + v.Encode()

	resp, err := provider.Get(path, &constants.DefaultStatsHeaders)
	if err != nil {
		return types.Response[types.DraftCombineStatsResponse]{Error: err}
	}
	defer resp.Body.Close()

	rawResp, err := internal.ParseResponse(resp)
	if err != nil {
		return types.Response[types.DraftCombineStatsResponse]{StatusCode: resp.StatusCode, Error: err}
	}

	contents, err := parser.ParseDraftCombineStatsResponse(rawResp)
	if err != nil {
		return types.Response[types.DraftCombineStatsResponse]{StatusCode: resp.StatusCode, Error: err}
	}

	return types.Response[types.DraftCombineStatsResponse]{
		Contents:   contents,
		StatusCode: resp.StatusCode,
		Error:      nil,
	}
}
