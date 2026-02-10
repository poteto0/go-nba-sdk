package live

import (
	"github.com/poteto0/go-nba-sdk/api"
	"github.com/poteto0/go-nba-sdk/constants"
	"github.com/poteto0/go-nba-sdk/internal"
	"github.com/poteto0/go-nba-sdk/types"
)

func GetBoxScore(provider api.IProvider, params *types.BoxScoreParams) types.Response[types.LiveBoxScoreResponse] {
	if err := validateBoxScoreParams(params); err != nil {
		return types.Response[types.LiveBoxScoreResponse]{Error: err}
	}

	filePath := "boxscore_" + params.GameID + ".json"
	path := constants.LiveBaseUrl + constants.BoxScorePath + "/" + filePath

	resp, err := provider.Get(path, &constants.DefaultLiveStatsHeaders)
	if err != nil {
		return types.Response[types.LiveBoxScoreResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
		}
	}
	defer resp.Body.Close()

	var boxscore types.LiveBoxScoreResponse
	err = internal.ParseResponseTo(resp, &boxscore)
	if err != nil {
		return types.Response[types.LiveBoxScoreResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
		}
	}

	return types.Response[types.LiveBoxScoreResponse]{
		Contents:   boxscore,
		StatusCode: resp.StatusCode,
		Error:      nil,
	}
}

func validateBoxScoreParams(params *types.BoxScoreParams) error {
	if params == nil {
		return types.NewGnsError("boxscore params is required")
	}

	if params.GameID == "" {
		return types.NewGnsError("game id is required")
	}
	return nil
}
