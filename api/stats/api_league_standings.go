package stats

import (
	"github.com/google/go-querystring/query"
	"github.com/poteto0/go-nba-sdk/api"
	"github.com/poteto0/go-nba-sdk/constants"
	"github.com/poteto0/go-nba-sdk/internal"
	"github.com/poteto0/go-nba-sdk/parser"
	"github.com/poteto0/go-nba-sdk/types"
)

var DefaultLeagueStandingsParams = types.LeagueStandingsParams{
	LeagueID:   "00",
	Season:     "2025-26",
	SeasonType: "Regular Season",
}

func GetLeagueStandings(provider api.IProvider, params *types.LeagueStandingsParams) types.Response[types.LeagueStandingsResponseContent] {
	if params == nil {
		params = &DefaultLeagueStandingsParams
	}

	if params.LeagueID == "" {
		params.LeagueID = DefaultLeagueStandingsParams.LeagueID
	}

	if params.Season == "" {
		params.Season = DefaultLeagueStandingsParams.Season
	}

	if params.SeasonType == "" {
		params.SeasonType = DefaultLeagueStandingsParams.SeasonType
	}

	if err := validateLeagueStandingsParams(*params); err != nil {
		return types.Response[types.LeagueStandingsResponseContent]{Error: err}
	}

	path := constants.StatsBaseUrl + constants.LeagueStandingsPath
	v, err := query.Values(params)
	if err != nil {
		return types.Response[types.LeagueStandingsResponseContent]{Error: err}
	}

	path = path + "?" + v.Encode()

	resp, err := provider.Get(path, &constants.DefaultStatsHeaders)
	if err != nil {
		return types.Response[types.LeagueStandingsResponseContent]{
			Error: err,
		}
	}
	defer resp.Body.Close()

	rawResp, err := internal.ParseResponse(resp)
	if err != nil {
		return types.Response[types.LeagueStandingsResponseContent]{
			StatusCode: resp.StatusCode,
			Error:      err,
		}
	}

	contents, err := parser.ParseLeagueStandingsResponse(rawResp)
	if err != nil {
		return types.Response[types.LeagueStandingsResponseContent]{
			StatusCode: resp.StatusCode,
			Error:      err,
		}
	}

	return types.Response[types.LeagueStandingsResponseContent]{
		Contents:   contents,
		StatusCode: resp.StatusCode,
		Error:      nil,
	}
}

func validateLeagueStandingsParams(params types.LeagueStandingsParams) error {
	if params.SeasonType != "Regular Season" && params.SeasonType != "Pre Season" {
		return types.NewGnsError("SeasonType is allowed, 'Regular Season' or 'Pre Season'")
	}

	return nil
}
