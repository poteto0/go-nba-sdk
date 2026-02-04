package stats

import (
	"github.com/google/go-querystring/query"
	"github.com/poteto0/go-nba-sdk/constants"
	"github.com/poteto0/go-nba-sdk/gns"
	"github.com/poteto0/go-nba-sdk/internal"
	"github.com/poteto0/go-nba-sdk/parser"
	"github.com/poteto0/go-nba-sdk/types"
)

var DefaultPlayerCareerStatsParams = PlayerCareerStatsParams{
	LeagueID: "00",
	PerMode:  "PerGame",
}

type PlayerCareerStatsParams struct {
	PlayerID string `url:"PlayerID"`
	LeagueID string `url:"LeagueID"`
	PerMode  string `url:"PerMode"`
}

func GetPlayerCareerStats(client gns.IClient, params PlayerCareerStatsParams) types.Response[types.PlayerCareerStatsResponseContent] {
	if params.LeagueID == "" {
		params.LeagueID = DefaultPlayerCareerStatsParams.LeagueID
	}
	if params.PerMode == "" {
		params.PerMode = DefaultPlayerCareerStatsParams.PerMode
	}

	path := constants.StatsBaseUrl + constants.PlayerCareerStatsPath
	v, err := query.Values(params)
	if err != nil {
		return types.Response[types.PlayerCareerStatsResponseContent]{Error: err}
	}

	path = path + "?" + v.Encode()

	resp, err := client.Get(path)
	if err != nil {
		return types.Response[types.PlayerCareerStatsResponseContent]{Error: err}
	}
	defer resp.Body.Close()

	rawResp, err := internal.ParseResponse(resp)
	if err != nil {
		return types.Response[types.PlayerCareerStatsResponseContent]{Error: err}
	}

	contents, err := parser.ParsePlayerCareerStatsResponse(rawResp)
	if err != nil {
		return types.Response[types.PlayerCareerStatsResponseContent]{Error: err}
	}

	return types.Response[types.PlayerCareerStatsResponseContent]{
		Contents:   contents,
		StatusCode: resp.StatusCode,
		Error:      nil,
	}
}
