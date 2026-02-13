package live

import (
	"github.com/poteto0/go-nba-sdk/api"
	"github.com/poteto0/go-nba-sdk/constants"
	"github.com/poteto0/go-nba-sdk/internal"
	"github.com/poteto0/go-nba-sdk/types"
)

func GetPlayByPlay(provider api.IProvider, params *types.PlayByPlayParams) types.Response[types.LivePlayByPlayResponse] {
	if err := validatePlayByPlayParams(params); err != nil {
		return types.Response[types.LivePlayByPlayResponse]{Error: err}
	}

	filePath := "playbyplay_" + params.GameID + ".json"
	path := constants.LiveBaseUrl + constants.PlayByPlayPath + "/" + filePath

	resp, err := provider.Get(path, &constants.DefaultLiveStatsHeaders)
	if err != nil {
		return types.Response[types.LivePlayByPlayResponse]{
			Error: err,
		}
	}
	defer resp.Body.Close()

	var playbyplay types.LivePlayByPlayResponse
	err = internal.ParseResponseTo(resp, &playbyplay)
	if err != nil {
		return types.Response[types.LivePlayByPlayResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
		}
	}

	return types.Response[types.LivePlayByPlayResponse]{
		Contents:   playbyplay,
		StatusCode: resp.StatusCode,
		Error:      nil,
	}
}

func validatePlayByPlayParams(params *types.PlayByPlayParams) error {
	if params == nil {
		return types.NewGnsError("playbyplay params is required")
	}

	if params.GameID == "" {
		return types.NewGnsError("game id is required")
	}
	return nil
}
