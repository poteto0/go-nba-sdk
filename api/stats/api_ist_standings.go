package stats

import (
	"github.com/google/go-querystring/query"
	"github.com/poteto0/go-nba-sdk/api"
	"github.com/poteto0/go-nba-sdk/constants"
	"github.com/poteto0/go-nba-sdk/internal"
	"github.com/poteto0/go-nba-sdk/types"
)

var DefaultIstStandingsParams = types.IstStandingsParams{
	LeagueID: "00",
	Season:   "2025-26",
	Section:  "group",
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

	if params.Section == "" {
		params.Section = DefaultIstStandingsParams.Section
	}

	if err := validateIstStandingsParams(*params); err != nil {
		return types.Response[types.IstStandingsResponseContent]{Error: err}
	}

	path := constants.StatsBaseUrl + constants.IstStandingsPath
	v, err := query.Values(params)
	if err != nil {
		return types.Response[types.IstStandingsResponseContent]{
			Error: err,
		}
	}

	path = path + "?" + v.Encode()

	resp, err := provider.Get(path, nil)
	if err != nil {
		return types.Response[types.IstStandingsResponseContent]{
			Error: err,
		}
	}
	defer resp.Body.Close()

	var contents types.IstStandingsResponseContent
	if err := internal.ParseResponseTo(resp, &contents); err != nil {
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

func validateIstStandingsParams(params types.IstStandingsParams) error {
	if params.Section != "group" && params.Section != "wildcard" {
		return types.NewGnsError("section is allowed, group or wildcard")
	}

	return nil
}
