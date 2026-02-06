package live

import (
	"github.com/poteto0/go-nba-sdk/api"
	"github.com/poteto0/go-nba-sdk/constants"
	"github.com/poteto0/go-nba-sdk/internal"
	"github.com/poteto0/go-nba-sdk/types"
)

var DefaultScoreBoardParams = types.ScoreBoardParams{
	LeagueID: "00",
}

func GetScoreBoard(provider api.IProvider, params *types.ScoreBoardParams) types.Response[types.LiveScoreBoardResponse] {
	if params == nil {
		params = &DefaultScoreBoardParams
	}

	if params.LeagueID == "" {
		params.LeagueID = DefaultScoreBoardParams.LeagueID
	}

	filePath := "todaysScoreboard_" + params.LeagueID + ".json"
	path := constants.LiveBaseUrl + constants.ScoreBoardPath + "/" + filePath

	resp, err := provider.Get(path, &constants.DefaultLiveStatsHeaders)
	if err != nil {
		return types.Response[types.LiveScoreBoardResponse]{Error: err}
	}
	defer resp.Body.Close()

	var scoreBoard types.LiveScoreBoardResponse
	err = internal.ParseResponseTo(resp, &scoreBoard)
	if err != nil {
		return types.Response[types.LiveScoreBoardResponse]{Error: err}
	}

	return types.Response[types.LiveScoreBoardResponse]{
		Contents:   scoreBoard,
		StatusCode: resp.StatusCode,
		Error:      nil,
	}
}
