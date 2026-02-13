package stats

import (
	"github.com/google/go-querystring/query"
	"github.com/poteto0/go-nba-sdk/api"
	"github.com/poteto0/go-nba-sdk/constants"
	"github.com/poteto0/go-nba-sdk/internal"
	"github.com/poteto0/go-nba-sdk/types"
)

var DefaultScheduleLeagueV2Params = types.ScheduleLeagueV2Params{
	LeagueID: "00",
	Season:   "2025-26",
}

func GetScheduleLeagueV2(provider api.IProvider, params *types.ScheduleLeagueV2Params) types.Response[types.ScheduleLeagueV2Response] {
	if params == nil {
		params = &DefaultScheduleLeagueV2Params
	}

	if params.LeagueID == "" {
		params.LeagueID = DefaultScheduleLeagueV2Params.LeagueID
	}

	if params.Season == "" {
		params.Season = DefaultScheduleLeagueV2Params.Season
	}

	path := constants.StatsBaseUrl + constants.ScheduleLeagueV2Path
	v, err := query.Values(params)
	if err != nil {
		return types.Response[types.ScheduleLeagueV2Response]{Error: err}
	}

	path = path + "?" + v.Encode()

	resp, err := provider.Get(path, nil)
	if err != nil {
		return types.Response[types.ScheduleLeagueV2Response]{
			StatusCode: resp.StatusCode,
			Error:      err,
		}
	}
	defer resp.Body.Close()

	var data types.ScheduleLeagueV2Response
	if err := internal.ParseResponseTo(resp, &data); err != nil {
		return types.Response[types.ScheduleLeagueV2Response]{
			StatusCode: resp.StatusCode,
			Error:      err,
		}
	}

	return types.Response[types.ScheduleLeagueV2Response]{
		Contents:   data,
		StatusCode: resp.StatusCode,
		Error:      nil,
	}
}
