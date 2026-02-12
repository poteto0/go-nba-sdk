package stats

import (
	"github.com/google/go-querystring/query"
	"github.com/poteto0/go-nba-sdk/api"
	"github.com/poteto0/go-nba-sdk/constants"
	"github.com/poteto0/go-nba-sdk/internal"
	"github.com/poteto0/go-nba-sdk/parser"
	"github.com/poteto0/go-nba-sdk/types"
)

var DefaultIstStandingsParams = types.IstStandingsParams{
	LeagueID: "00",
	Season:   "2023-24",
}

func GetIstStandings(provider api.IProvider, params *types.IstStandingsParams) types.Response[types.IstStandingsResponseContent] {
	if params == nil {
		params = &DefaultIstStandingsParams
	}

	if params.LeagueID == "" {
		params.LeagueID = DefaultIstStandingsParams.LeagueID
	}

	if params.Season == "" {
		params.Season = DefaultIstStandingsParams.Season
	}

	path := constants.StatsBaseUrl + constants.IstStandingsPath
	v, err := query.Values(params)
	if err != nil {
		return types.Response[types.IstStandingsResponseContent]{Error: err}
	}

	path = path + "?" + v.Encode()

	resp, err := provider.Get(path, nil)
	if err != nil {
		return types.Response[types.IstStandingsResponseContent]{
			StatusCode: resp.StatusCode,
			Error:      err,
		}
	}
	defer resp.Body.Close()

	rawResp, err := internal.ParseResponse(resp)
	if err != nil {
		return types.Response[types.IstStandingsResponseContent]{
			StatusCode: resp.StatusCode,
			Error:      err,
		}
	}

	contents, err := parser.ParseIstStandingsResponse(rawResp)
	if err != nil {
		return types.Response[types.IstStandingsResponseContent]{
			StatusCode: resp.StatusCode,
			Error:      err,
		}
	}

	return types.Response[types.IstStandingsResponseContent]{
		Contents:   contents,
		StatusCode: resp.StatusCode,
		Error:      nil,
	}
}
